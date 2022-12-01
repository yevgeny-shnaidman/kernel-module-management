// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	authn "github.com/google/go-containerregistry/pkg/authn"
	v1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
)

// MockRegistryAuthGetter is a mock of RegistryAuthGetter interface.
type MockRegistryAuthGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryAuthGetterMockRecorder
}

// MockRegistryAuthGetterMockRecorder is the mock recorder for MockRegistryAuthGetter.
type MockRegistryAuthGetterMockRecorder struct {
	mock *MockRegistryAuthGetter
}

// NewMockRegistryAuthGetter creates a new mock instance.
func NewMockRegistryAuthGetter(ctrl *gomock.Controller) *MockRegistryAuthGetter {
	mock := &MockRegistryAuthGetter{ctrl: ctrl}
	mock.recorder = &MockRegistryAuthGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistryAuthGetter) EXPECT() *MockRegistryAuthGetterMockRecorder {
	return m.recorder
}

// GetKeyChain mocks base method.
func (m *MockRegistryAuthGetter) GetKeyChain(ctx context.Context) (authn.Keychain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyChain", ctx)
	ret0, _ := ret[0].(authn.Keychain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyChain indicates an expected call of GetKeyChain.
func (mr *MockRegistryAuthGetterMockRecorder) GetKeyChain(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyChain", reflect.TypeOf((*MockRegistryAuthGetter)(nil).GetKeyChain), ctx)
}

// MockRegistryAuthGetterFactory is a mock of RegistryAuthGetterFactory interface.
type MockRegistryAuthGetterFactory struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryAuthGetterFactoryMockRecorder
}

// MockRegistryAuthGetterFactoryMockRecorder is the mock recorder for MockRegistryAuthGetterFactory.
type MockRegistryAuthGetterFactoryMockRecorder struct {
	mock *MockRegistryAuthGetterFactory
}

// NewMockRegistryAuthGetterFactory creates a new mock instance.
func NewMockRegistryAuthGetterFactory(ctrl *gomock.Controller) *MockRegistryAuthGetterFactory {
	mock := &MockRegistryAuthGetterFactory{ctrl: ctrl}
	mock.recorder = &MockRegistryAuthGetterFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistryAuthGetterFactory) EXPECT() *MockRegistryAuthGetterFactoryMockRecorder {
	return m.recorder
}

// NewClusterAuthGetter mocks base method.
func (m *MockRegistryAuthGetterFactory) NewClusterAuthGetter() RegistryAuthGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClusterAuthGetter")
	ret0, _ := ret[0].(RegistryAuthGetter)
	return ret0
}

// NewClusterAuthGetter indicates an expected call of NewClusterAuthGetter.
func (mr *MockRegistryAuthGetterFactoryMockRecorder) NewClusterAuthGetter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClusterAuthGetter", reflect.TypeOf((*MockRegistryAuthGetterFactory)(nil).NewClusterAuthGetter))
}

// NewRegistryAuthGetterFrom mocks base method.
func (m *MockRegistryAuthGetterFactory) NewRegistryAuthGetterFrom(mod *v1beta1.Module) RegistryAuthGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRegistryAuthGetterFrom", mod)
	ret0, _ := ret[0].(RegistryAuthGetter)
	return ret0
}

// NewRegistryAuthGetterFrom indicates an expected call of NewRegistryAuthGetterFrom.
func (mr *MockRegistryAuthGetterFactoryMockRecorder) NewRegistryAuthGetterFrom(mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRegistryAuthGetterFrom", reflect.TypeOf((*MockRegistryAuthGetterFactory)(nil).NewRegistryAuthGetterFrom), mod)
}
