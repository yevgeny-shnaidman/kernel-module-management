// Code generated by MockGen. DO NOT EDIT.
// Source: build_sign_reconciler.go
//
// Generated by this command:
//
//	mockgen -source=build_sign_reconciler.go -package=controllers -destination=mock_build_sign_reconciler.go buildSignReconcilerHelperAPI
//
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

// MockbuildSignReconcilerHelperAPI is a mock of buildSignReconcilerHelperAPI interface.
type MockbuildSignReconcilerHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockbuildSignReconcilerHelperAPIMockRecorder
}

// MockbuildSignReconcilerHelperAPIMockRecorder is the mock recorder for MockbuildSignReconcilerHelperAPI.
type MockbuildSignReconcilerHelperAPIMockRecorder struct {
	mock *MockbuildSignReconcilerHelperAPI
}

// NewMockbuildSignReconcilerHelperAPI creates a new mock instance.
func NewMockbuildSignReconcilerHelperAPI(ctrl *gomock.Controller) *MockbuildSignReconcilerHelperAPI {
	mock := &MockbuildSignReconcilerHelperAPI{ctrl: ctrl}
	mock.recorder = &MockbuildSignReconcilerHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockbuildSignReconcilerHelperAPI) EXPECT() *MockbuildSignReconcilerHelperAPIMockRecorder {
	return m.recorder
}

// garbageCollect mocks base method.
func (m *MockbuildSignReconcilerHelperAPI) garbageCollect(ctx context.Context, mod *v1beta1.Module, mldMappings map[string]*api.ModuleLoaderData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "garbageCollect", ctx, mod, mldMappings)
	ret0, _ := ret[0].(error)
	return ret0
}

// garbageCollect indicates an expected call of garbageCollect.
func (mr *MockbuildSignReconcilerHelperAPIMockRecorder) garbageCollect(ctx, mod, mldMappings any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "garbageCollect", reflect.TypeOf((*MockbuildSignReconcilerHelperAPI)(nil).garbageCollect), ctx, mod, mldMappings)
}

// getNodesListBySelector mocks base method.
func (m *MockbuildSignReconcilerHelperAPI) getNodesListBySelector(ctx context.Context, mod *v1beta1.Module) ([]v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNodesListBySelector", ctx, mod)
	ret0, _ := ret[0].([]v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getNodesListBySelector indicates an expected call of getNodesListBySelector.
func (mr *MockbuildSignReconcilerHelperAPIMockRecorder) getNodesListBySelector(ctx, mod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNodesListBySelector", reflect.TypeOf((*MockbuildSignReconcilerHelperAPI)(nil).getNodesListBySelector), ctx, mod)
}

// getRelevantKernelMappings mocks base method.
func (m *MockbuildSignReconcilerHelperAPI) getRelevantKernelMappings(ctx context.Context, mod *v1beta1.Module, targetedNodes []v1.Node) (map[string]*api.ModuleLoaderData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRelevantKernelMappings", ctx, mod, targetedNodes)
	ret0, _ := ret[0].(map[string]*api.ModuleLoaderData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getRelevantKernelMappings indicates an expected call of getRelevantKernelMappings.
func (mr *MockbuildSignReconcilerHelperAPIMockRecorder) getRelevantKernelMappings(ctx, mod, targetedNodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRelevantKernelMappings", reflect.TypeOf((*MockbuildSignReconcilerHelperAPI)(nil).getRelevantKernelMappings), ctx, mod, targetedNodes)
}

// getRequestedModule mocks base method.
func (m *MockbuildSignReconcilerHelperAPI) getRequestedModule(ctx context.Context, namespacedName types.NamespacedName) (*v1beta1.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRequestedModule", ctx, namespacedName)
	ret0, _ := ret[0].(*v1beta1.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getRequestedModule indicates an expected call of getRequestedModule.
func (mr *MockbuildSignReconcilerHelperAPIMockRecorder) getRequestedModule(ctx, namespacedName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRequestedModule", reflect.TypeOf((*MockbuildSignReconcilerHelperAPI)(nil).getRequestedModule), ctx, namespacedName)
}

// handleBuild mocks base method.
func (m *MockbuildSignReconcilerHelperAPI) handleBuild(ctx context.Context, mld *api.ModuleLoaderData) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleBuild", ctx, mld)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// handleBuild indicates an expected call of handleBuild.
func (mr *MockbuildSignReconcilerHelperAPIMockRecorder) handleBuild(ctx, mld any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleBuild", reflect.TypeOf((*MockbuildSignReconcilerHelperAPI)(nil).handleBuild), ctx, mld)
}

// handleSigning mocks base method.
func (m *MockbuildSignReconcilerHelperAPI) handleSigning(ctx context.Context, mld *api.ModuleLoaderData) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleSigning", ctx, mld)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// handleSigning indicates an expected call of handleSigning.
func (mr *MockbuildSignReconcilerHelperAPIMockRecorder) handleSigning(ctx, mld any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleSigning", reflect.TypeOf((*MockbuildSignReconcilerHelperAPI)(nil).handleSigning), ctx, mld)
}
