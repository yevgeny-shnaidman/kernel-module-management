// Code generated by MockGen. DO NOT EDIT.
// Source: daemonset.go

// Package daemonset is a generated GoMock package.
package daemonset

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockDaemonSetCreator is a mock of DaemonSetCreator interface.
type MockDaemonSetCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetCreatorMockRecorder
}

// MockDaemonSetCreatorMockRecorder is the mock recorder for MockDaemonSetCreator.
type MockDaemonSetCreatorMockRecorder struct {
	mock *MockDaemonSetCreator
}

// NewMockDaemonSetCreator creates a new mock instance.
func NewMockDaemonSetCreator(ctrl *gomock.Controller) *MockDaemonSetCreator {
	mock := &MockDaemonSetCreator{ctrl: ctrl}
	mock.recorder = &MockDaemonSetCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetCreator) EXPECT() *MockDaemonSetCreatorMockRecorder {
	return m.recorder
}

// GarbageCollect mocks base method.
func (m *MockDaemonSetCreator) GarbageCollect(ctx context.Context, existingDS map[string]*v1.DaemonSet, validKernels sets.String) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GarbageCollect", ctx, existingDS, validKernels)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GarbageCollect indicates an expected call of GarbageCollect.
func (mr *MockDaemonSetCreatorMockRecorder) GarbageCollect(ctx, existingDS, validKernels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GarbageCollect", reflect.TypeOf((*MockDaemonSetCreator)(nil).GarbageCollect), ctx, existingDS, validKernels)
}

// GetNodeLabelFromPod mocks base method.
func (m *MockDaemonSetCreator) GetNodeLabelFromPod(pod *v10.Pod, moduleName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeLabelFromPod", pod, moduleName)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNodeLabelFromPod indicates an expected call of GetNodeLabelFromPod.
func (mr *MockDaemonSetCreatorMockRecorder) GetNodeLabelFromPod(pod, moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeLabelFromPod", reflect.TypeOf((*MockDaemonSetCreator)(nil).GetNodeLabelFromPod), pod, moduleName)
}

// ModuleDaemonSetsByKernelVersion mocks base method.
func (m *MockDaemonSetCreator) ModuleDaemonSetsByKernelVersion(ctx context.Context, name, namespace string) (map[string]*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModuleDaemonSetsByKernelVersion", ctx, name, namespace)
	ret0, _ := ret[0].(map[string]*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModuleDaemonSetsByKernelVersion indicates an expected call of ModuleDaemonSetsByKernelVersion.
func (mr *MockDaemonSetCreatorMockRecorder) ModuleDaemonSetsByKernelVersion(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModuleDaemonSetsByKernelVersion", reflect.TypeOf((*MockDaemonSetCreator)(nil).ModuleDaemonSetsByKernelVersion), ctx, name, namespace)
}

// SetDevicePluginAsDesired mocks base method.
func (m *MockDaemonSetCreator) SetDevicePluginAsDesired(ctx context.Context, ds *v1.DaemonSet, mod *v1alpha1.Module) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDevicePluginAsDesired", ctx, ds, mod)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDevicePluginAsDesired indicates an expected call of SetDevicePluginAsDesired.
func (mr *MockDaemonSetCreatorMockRecorder) SetDevicePluginAsDesired(ctx, ds, mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDevicePluginAsDesired", reflect.TypeOf((*MockDaemonSetCreator)(nil).SetDevicePluginAsDesired), ctx, ds, mod)
}

// SetDriverContainerAsDesired mocks base method.
func (m *MockDaemonSetCreator) SetDriverContainerAsDesired(ctx context.Context, ds *v1.DaemonSet, image string, mod v1alpha1.Module, kernelVersion string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDriverContainerAsDesired", ctx, ds, image, mod, kernelVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDriverContainerAsDesired indicates an expected call of SetDriverContainerAsDesired.
func (mr *MockDaemonSetCreatorMockRecorder) SetDriverContainerAsDesired(ctx, ds, image, mod, kernelVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDriverContainerAsDesired", reflect.TypeOf((*MockDaemonSetCreator)(nil).SetDriverContainerAsDesired), ctx, ds, image, mod, kernelVersion)
}
