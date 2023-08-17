// Code generated by MockGen. DO NOT EDIT.
// Source: module_nmc_reconciler.go

// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	v1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	api "github.com/rh-ecosystem-edge/kernel-module-management/internal/api"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
)

// MockmoduleNMCReconcilerHelperAPI is a mock of moduleNMCReconcilerHelperAPI interface.
type MockmoduleNMCReconcilerHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockmoduleNMCReconcilerHelperAPIMockRecorder
}

// MockmoduleNMCReconcilerHelperAPIMockRecorder is the mock recorder for MockmoduleNMCReconcilerHelperAPI.
type MockmoduleNMCReconcilerHelperAPIMockRecorder struct {
	mock *MockmoduleNMCReconcilerHelperAPI
}

// NewMockmoduleNMCReconcilerHelperAPI creates a new mock instance.
func NewMockmoduleNMCReconcilerHelperAPI(ctrl *gomock.Controller) *MockmoduleNMCReconcilerHelperAPI {
	mock := &MockmoduleNMCReconcilerHelperAPI{ctrl: ctrl}
	mock.recorder = &MockmoduleNMCReconcilerHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmoduleNMCReconcilerHelperAPI) EXPECT() *MockmoduleNMCReconcilerHelperAPIMockRecorder {
	return m.recorder
}

// disableModuleOnNode mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) disableModuleOnNode(ctx context.Context, modNamespace, modName, nodeName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "disableModuleOnNode", ctx, modNamespace, modName, nodeName)
	ret0, _ := ret[0].(error)
	return ret0
}

// disableModuleOnNode indicates an expected call of disableModuleOnNode.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) disableModuleOnNode(ctx, modNamespace, modName, nodeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "disableModuleOnNode", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).disableModuleOnNode), ctx, modNamespace, modName, nodeName)
}

// enableModuleOnNode mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) enableModuleOnNode(ctx context.Context, mld *api.ModuleLoaderData, nodeName, kernelVersion string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "enableModuleOnNode", ctx, mld, nodeName, kernelVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// enableModuleOnNode indicates an expected call of enableModuleOnNode.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) enableModuleOnNode(ctx, mld, nodeName, kernelVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "enableModuleOnNode", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).enableModuleOnNode), ctx, mld, nodeName, kernelVersion)
}

// finalizeModule mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) finalizeModule(ctx context.Context, mod *v1beta1.Module) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "finalizeModule", ctx, mod)
	ret0, _ := ret[0].(error)
	return ret0
}

// finalizeModule indicates an expected call of finalizeModule.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) finalizeModule(ctx, mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "finalizeModule", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).finalizeModule), ctx, mod)
}

// getNodesList mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) getNodesList(ctx context.Context) ([]v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNodesList", ctx)
	ret0, _ := ret[0].([]v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getNodesList indicates an expected call of getNodesList.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) getNodesList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNodesList", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).getNodesList), ctx)
}

// getRequestedModule mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) getRequestedModule(ctx context.Context, namespacedName types.NamespacedName) (*v1beta1.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRequestedModule", ctx, namespacedName)
	ret0, _ := ret[0].(*v1beta1.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getRequestedModule indicates an expected call of getRequestedModule.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) getRequestedModule(ctx, namespacedName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRequestedModule", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).getRequestedModule), ctx, namespacedName)
}

// setFinalizer mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) setFinalizer(ctx context.Context, mod *v1beta1.Module) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setFinalizer", ctx, mod)
	ret0, _ := ret[0].(error)
	return ret0
}

// setFinalizer indicates an expected call of setFinalizer.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) setFinalizer(ctx, mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setFinalizer", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).setFinalizer), ctx, mod)
}

// shouldModuleRunOnNode mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) shouldModuleRunOnNode(node v1.Node, mld *api.ModuleLoaderData) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "shouldModuleRunOnNode", node, mld)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// shouldModuleRunOnNode indicates an expected call of shouldModuleRunOnNode.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) shouldModuleRunOnNode(node, mld interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "shouldModuleRunOnNode", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).shouldModuleRunOnNode), node, mld)
}
