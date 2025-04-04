// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package build is a generated GoMock package.
package build

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1alpha1"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Sync mocks base method.
func (m_2 *MockManager) Sync(ctx context.Context, mod v1alpha1.Module, m v1alpha1.KernelMapping, targetKernel string) (Result, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Sync", ctx, mod, m, targetKernel)
	ret0, _ := ret[0].(Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockManagerMockRecorder) Sync(ctx, mod, m, targetKernel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockManager)(nil).Sync), ctx, mod, m, targetKernel)
}
