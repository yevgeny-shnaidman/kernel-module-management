/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"os"
	//"strings"

	"github.com/rh-ecosystem-edge/kernel-module-management/internal/registry"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/utilexecute"
	"github.com/google/go-containerregistry/pkg/authn"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/docker/cli/cli/config"
	dockertypes "github.com/docker/cli/cli/config/types"
	"github.com/alessio/shellescape"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

/*
var dockerfile string = `ARG START_IMAGE
FROM quay.io/yshnaidm/fedora-coreos:testing-devel-with-add-layers
RUN ostree container commit
`
*/
var dockerfile string = `ARG START_IMAGE
FROM ${START_IMAGE}
RUN ostree container commit
`

func main() {
	registryAPI := registry.NewRegistry()
	ctx := context.Background()

	imageName := "quay.io/yshnaidm/fedora-coreos:testing-devel"
	layersImageName := "quay.io/yshnaidm/check-layer:latest"

	res, err := registryAPI.ImageExists(ctx, imageName, nil, nil)
	if err != nil {
		fmt.Printf("failed to check if image %s exists: %w\n", imageName, err)
	} else {
		fmt.Printf("check if image %s exists: %t\n", imageName, res)
	}

	res, err = registryAPI.ImageExists(ctx, layersImageName, nil, nil)
        if err != nil {
                fmt.Printf("failed to check if image %s exists: %w\n", layersImageName, err)
        } else {
                fmt.Printf("check if image %s exists: %t\n", layersImageName, res)
        }

	image, err := registryAPI.GetImageByName(imageName, nil)
	if err != nil {
		fmt.Printf("failed to get image: %w\n", err)
		return
	}

	digests, repoConfig, err := registryAPI.GetLayersDigests(ctx, layersImageName, nil, nil)
	if err != nil {
		fmt.Printf("failed to get layers: %w\n", err)
		return
	}



	fmt.Printf("number of layers is %d\n", len(digests))
	layersToAdd := []v1.Layer{}

	for i := len(digests) - 1; i >= 0; i-- {
                layer, err := registryAPI.GetLayerByDigest(digests[i], repoConfig)
                if err != nil {
			fmt.Printf("failed to get layer for digest %s: %v", digests[i], err)
                }

		if registryAPI.FileExists(layer, "etc/FirstLayer.txt") {
			break
		}

		layersToAdd = append(layersToAdd, layer)
		/*
		_, err = registryAPI.PrintAllFilesInLayer(layer)
		if err != nil {
			fmt.Printf("failed in PrintAllFilesInLayer: %s", err)
		}
		*/
        }

	newImage := image

	fmt.Printf("number of layers to add is %d\n", len(layersToAdd))
	for i := len(layersToAdd) -1; i >= 0; i-- {
		newImage, err = registryAPI.AddOneLayerToImage(layersToAdd[i], newImage)
		if err != nil {
			fmt.Printf("Failed to add layer %d to image\n", i)
			return
		}
	}

	ath, err := getAuthFromFile("/home/yevgeny/.docker/config.json", "quay.io")
	if err != nil {
		fmt.Printf("could not read docker auth file: %w\n", err)
		return
	}

	interImage := "quay.io/yshnaidm/fedora-coreos:testing-devel-with-add-layers"

	err = registryAPI.WriteImageByName(interImage, newImage, ath)
	if err != nil {
		fmt.Printf("failed to push image quay.io/yshnaidm/simple-kmod:minikube-with-add-layers: %w\n", err)
		return
	}
	fmt.Printf("successuflly pushed image quay.io/yshnaidm/simple-kmod:minikube-with-add-layers\n")

	finalImage := "quay.io/yshnaidm/fedora-coreos:testing-devel-with-add-layers-final"
	dockerfilePath := "/home/yevgeny/UPSTREAM_MINIKUBE_TESTING_FILES/BUILD_LAYERS/Dockerfile.test"
	err = createDockerfile(dockerfilePath, dockerfile)
	if err != nil {
		fmt.Printf("create dockerfile failed %w\n", err)
		return
	}
	fmt.Printf("create dockerfile sucessfull\n")
	err = buildFinalImage(interImage, finalImage, dockerfilePath)
	if err != nil {
		fmt.Printf("failed to build image %s: %w\n", finalImage, err)
		return
	}

	fmt.Printf("Successfully built image %s\n", finalImage)
	err = pushFinalImage(finalImage)
	if err != nil {
		fmt.Printf("failed to push image %s\n", finalImage)
		return
	}
	fmt.Printf("Successfully pushed image %s\n", finalImage)
}

func createDockerfile(filePath, fileData string) error {
	return os.WriteFile(filePath, []byte(fileData), 0644)
}

func buildFinalImage(interImage, finalImage, dockerfilePath string) error {
	startImage := fmt.Sprintf("START_IMAGE=%s", interImage)
	command := []string{"docker", "build", "--build-arg", startImage, "-f", dockerfilePath, "-t", finalImage, "."}
	//fmt.Sprintf("%s %s %s", shellescape.QuoteCommand(dockerCommand), dockerfilePath, dockerArgs)
	//args := []string{"-c", strings.Join(command, " ")}
	args := []string{"-c", shellescape.QuoteCommand(command)}
	fmt.Printf("commands args <%s>\n", shellescape.QuoteCommand(command))
	stdout, stderr, errCode := utilexecute.Execute("sh", args...)
	fmt.Printf("stdout is <%s>\n", stdout)
	fmt.Printf("stderr is <%s>\n", stderr)
	if errCode != 0 {
		return fmt.Errorf("ExecutePrivileged returned error")
	}
	fmt.Printf("utilexecute.ExecutePrivileged successfull\n")
	return  nil
}

func pushFinalImage(finalImage string) error {
	command := []string{"docker", "push", finalImage}
	args := []string{"-c", shellescape.QuoteCommand(command)}
	_, _, errCode := utilexecute.Execute("sh", args...)
	if errCode != 0 {
                return fmt.Errorf("ExecutePrivileged returned error")
        }
        fmt.Printf("utilexecute.ExecutePrivileged successfull\n")
        return  nil
}

func getAuthFromFile(configfile string, repo string) (authn.Authenticator, error) {

        if configfile == "" {
                fmt.Printf("no pull secret defined, default to Anonymous\n")
                return authn.Anonymous, nil
        }

        f, err := os.Open(configfile)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        cf, err := config.LoadFromReader(f)
        if err != nil {
                return nil, err
        }

        var cfg dockertypes.AuthConfig

        cfg, err = cf.GetAuthConfig(repo)
        if err != nil {
                return nil, err
        }

        return authn.FromConfig(authn.AuthConfig{
                Username:      cfg.Username,
                Password:      cfg.Password,
                Auth:          cfg.Auth,
                IdentityToken: cfg.IdentityToken,
                RegistryToken: cfg.RegistryToken,
        }), nil

}
