// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/env0/terraform-provider-env0/client/http (interfaces: HttpClientInterface)

// Package http is a generated GoMock package.
package http

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpClientInterface is a mock of HttpClientInterface interface.
type MockHttpClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientInterfaceMockRecorder
}

// MockHttpClientInterfaceMockRecorder is the mock recorder for MockHttpClientInterface.
type MockHttpClientInterfaceMockRecorder struct {
	mock *MockHttpClientInterface
}

// NewMockHttpClientInterface creates a new mock instance.
func NewMockHttpClientInterface(ctrl *gomock.Controller) *MockHttpClientInterface {
	mock := &MockHttpClientInterface{ctrl: ctrl}
	mock.recorder = &MockHttpClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpClientInterface) EXPECT() *MockHttpClientInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockHttpClientInterface) Delete(arg0 string, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockHttpClientInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHttpClientInterface)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockHttpClientInterface) Get(arg0 string, arg1 map[string]string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockHttpClientInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClientInterface)(nil).Get), arg0, arg1, arg2)
}

// Patch mocks base method.
func (m *MockHttpClientInterface) Patch(arg0 string, arg1, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockHttpClientInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockHttpClientInterface)(nil).Patch), arg0, arg1, arg2)
}

// Post mocks base method.
func (m *MockHttpClientInterface) Post(arg0 string, arg1, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockHttpClientInterfaceMockRecorder) Post(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHttpClientInterface)(nil).Post), arg0, arg1, arg2)
}

// Put mocks base method.
func (m *MockHttpClientInterface) Put(arg0 string, arg1, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockHttpClientInterfaceMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockHttpClientInterface)(nil).Put), arg0, arg1, arg2)
}
