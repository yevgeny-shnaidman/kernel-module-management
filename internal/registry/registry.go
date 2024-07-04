package registry

import (
	"archive/tar"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"

	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
	kmmv1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/auth"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	modulesLocationPath = "lib/modules"
)

type DriverToolkitEntry struct {
	ImageURL            string `json:"imageURL"`
	KernelFullVersion   string `json:"kernelFullVersion"`
	RTKernelFullVersion string `json:"RTKernelFullVersion"`
	OSVersion           string `json:"OSVersion"`
}

type RepoPullConfig struct {
	repo        string
	authOptions []crane.Option
}

//go:generate mockgen -source=registry.go -package=registry -destination=mock_registry_api.go

type Registry interface {
	ImageExists(ctx context.Context, image string, tlsOptions *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) (bool, error)
	VerifyModuleExists(layer v1.Layer, pathPrefix, kernelVersion, moduleFileName string) bool
	GetLayersDigests(ctx context.Context, image string, tlsOptions *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) ([]string, *RepoPullConfig, error)
	GetLayerByDigest(digest string, pullConfig *RepoPullConfig) (v1.Layer, error)
	LastLayer(ctx context.Context, image string, po *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) (v1.Layer, error)
	GetHeaderDataFromLayer(layer v1.Layer, headerName string) ([]byte, error)
	GetDigest(ctx context.Context, image string, tlsOptions *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) (string, error)
}

type registry struct {
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) ImageExists(ctx context.Context, image string, tlsOptions *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) (bool, error) {
	_, _, err := r.getImageManifest(ctx, image, tlsOptions, registryAuthGetter)
	if err != nil {
		te := &transport.Error{}
		if errors.As(err, &te) && te.StatusCode == http.StatusNotFound {
			return false, nil
		}
		return false, fmt.Errorf("could not get image %s: %w", image, err)
	}
	return true, nil
}

func (r *registry) GetLayersDigests(ctx context.Context, image string, tlsOptions *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) ([]string, *RepoPullConfig, error) {
	manifest, pullConfig, err := r.getImageManifest(ctx, image, tlsOptions, registryAuthGetter)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get manifest from image %s: %w", image, err)
	}

	digests, err := r.getLayersDigestsFromManifestStream(manifest)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get layers digests from manifest of the image %s: %w", image, err)
	}

	return digests, pullConfig, nil
}

func (r *registry) GetLayerByDigest(digest string, pullConfig *RepoPullConfig) (v1.Layer, error) {
	return crane.PullLayer(pullConfig.repo+"@"+digest, pullConfig.authOptions...)
}

func (r *registry) LastLayer(ctx context.Context, image string, po *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) (v1.Layer, error) {
	digests, repoConfig, err := r.GetLayersDigests(ctx, image, po, registryAuthGetter)
	if err != nil {
		return nil, fmt.Errorf("failed to get layers digests from image %s: %v", image, err)
	}
	return r.GetLayerByDigest(digests[len(digests)-1], repoConfig)
}

func (r *registry) VerifyModuleExists(layer v1.Layer, pathPrefix, kernelVersion, moduleFileName string) bool {
	// in layers headers, there is no root prefix
	fullPath := filepath.Join(strings.TrimPrefix(pathPrefix, "/"), modulesLocationPath, kernelVersion, moduleFileName)

	// if getHeaderReaderFromLayer does not return an error, it means that the file exists in the layer,
	// and that's all the indication that we need
	_, readerCloser, err := r.getHeaderReaderFromLayer(layer, fullPath)
	if err != nil {
		return false
	}
	readerCloser.Close()
	return true
}

func (r *registry) GetDigest(ctx context.Context, image string, tlsOptions *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) (string, error) {
	pullConfig, err := r.getPullOptions(ctx, image, tlsOptions, registryAuthGetter)
	if err != nil {
		return "", fmt.Errorf("failed to get pull options for image %s: %v", image, err)
	}

	digest, err := crane.Digest(image, pullConfig.authOptions...)
	if err != nil {
		return "", fmt.Errorf("failed to get digest for image %s: %v", image, err)
	}
	return digest, err
}

func (r *registry) getPullOptions(ctx context.Context, image string, tlsOptions *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) (*RepoPullConfig, error) {
	var repo string
	if hash := strings.Split(image, "@"); len(hash) > 1 {
		repo = hash[0]
	} else if tag := strings.Split(image, ":"); len(tag) > 1 {
		repo = tag[0]
	}

	options := []crane.Option{
		crane.WithContext(ctx),
	}

	if tlsOptions != nil {
		if tlsOptions.Insecure {
			options = append(options, crane.Insecure)
		}

		if tlsOptions.InsecureSkipTLSVerify {
			rt := http.DefaultTransport.(*http.Transport).Clone()
			rt.TLSClientConfig.InsecureSkipVerify = true

			options = append(
				options,
				crane.WithTransport(rt),
			)
		}
	}

	if registryAuthGetter != nil {
		keyChain, err := registryAuthGetter.GetKeyChain(ctx)
		if err != nil {
			return nil, fmt.Errorf("cannot get keychain from the registry auth getter: %w", err)
		}
		options = append(
			options,
			crane.WithAuthFromKeychain(keyChain),
		)
	}

	return &RepoPullConfig{repo: repo, authOptions: options}, nil
}

func (r *registry) getImageManifest(ctx context.Context, image string, tlsOptions *kmmv1beta1.TLSOptions, registryAuthGetter auth.RegistryAuthGetter) ([]byte, *RepoPullConfig, error) {
	pullConfig, err := r.getPullOptions(ctx, image, tlsOptions, registryAuthGetter)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get pull options for image %s: %w", image, err)
	}
	manifest, err := r.getManifestStreamFromImage(image, pullConfig.repo, pullConfig.authOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get manifest stream from image %s: %w", image, err)
	}

	return manifest, pullConfig, nil
}

func (r *registry) getManifestStreamFromImage(image, repo string, options []crane.Option) ([]byte, error) {
	manifest, err := crane.Manifest(image, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to get crane manifest from image %s: %w", image, err)
	}

	release := unstructured.Unstructured{}
	if err = json.Unmarshal(manifest, &release.Object); err != nil {
		return nil, fmt.Errorf("failed to unmarshal crane manifest: %w", err)
	}

	imageMediaType, mediaTypeFound, err := unstructured.NestedString(release.Object, "mediaType")
	if err != nil {
		return nil, fmt.Errorf("unmarshalled manifests invalid format: %w", err)
	}
	if !mediaTypeFound {
		return nil, fmt.Errorf("mediaType is missing from the image %s manifest", image)
	}

	if strings.Contains(imageMediaType, "manifest.list") {
		archDigest, err := r.getImageDigestFromMultiImage(manifest)
		if err != nil {
			return nil, fmt.Errorf("failed to get arch digets from multi arch image: %w", err)
		}
		// get the manifest stream for the image of the architecture
		manifest, err = crane.Manifest(repo+"@"+archDigest, options...)
		if err != nil {
			return nil, fmt.Errorf("failed to get crane manifest for the arch image: %w", err)
		}
	}
	return manifest, nil
}

func (r *registry) getLayersDigestsFromManifestStream(manifestStream []byte) ([]string, error) {
	manifest := v1.Manifest{}

	if err := json.Unmarshal(manifestStream, &manifest); err != nil {
		return nil, fmt.Errorf("failed to unmarshal manifest stream: %w", err)
	}

	digests := make([]string, len(manifest.Layers))
	for i, layer := range manifest.Layers {
		digests[i] = layer.Digest.Algorithm + ":" + layer.Digest.Hex
	}
	return digests, nil
}

func (r *registry) GetHeaderDataFromLayer(layer v1.Layer, headerName string) ([]byte, error) {
	reader, readerCloser, err := r.getHeaderReaderFromLayer(layer, headerName)
	if err != nil {
		return nil, fmt.Errorf("failed to get reader for the header %s in the layer: %v", headerName, err)
	}
	defer readerCloser.Close()
	buff, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read header %s data: %v", headerName, err)
	}
	return buff, nil
}

func (r *registry) getHeaderReaderFromLayer(layer v1.Layer, headerName string) (io.Reader, io.ReadCloser, error) {
	layerreader, err := layer.Uncompressed()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get layerreader from layer: %v", err)
	}
	// err ignored because we're only reading
	defer func() {
		if err != nil {
			layerreader.Close()
		}
	}()

	tr := tar.NewReader(layerreader)

	for {
		header, err := tr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, nil, fmt.Errorf("failed to get next entry from targz: %w", err)
		}
		if header.Name == headerName {
			return tr, layerreader, nil
		}
	}

	return nil, nil, fmt.Errorf("header %s not found in the layer", headerName)
}

func (r *registry) getImageDigestFromMultiImage(manifestListStream []byte) (string, error) {
	arch := runtime.GOARCH
	manifestList := v1.IndexManifest{}

	if err := json.Unmarshal(manifestListStream, &manifestList); err != nil {
		return "", fmt.Errorf("failed to unmarshal manifest stream: %w", err)
	}
	for _, manifest := range manifestList.Manifests {
		if manifest.Platform != nil && manifest.Platform.Architecture == arch {
			return manifest.Digest.Algorithm + ":" + manifest.Digest.Hex, nil
		}
	}
	return "", fmt.Errorf("Failed to find manifest for architecture %s", arch)
}
