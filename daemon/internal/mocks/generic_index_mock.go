// Code generated by MockGen. DO NOT EDIT.
// Source: ./daemon/internal/index/factory/config.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	factory "victord/daemon/internal/index/factory"

	gomock "github.com/golang/mock/gomock"
)

// MockGenericIndex is a mock of GenericIndex interface.
type MockGenericIndex struct {
	ctrl     *gomock.Controller
	recorder *MockGenericIndexMockRecorder
}

// MockGenericIndexMockRecorder is the mock recorder for MockGenericIndex.
type MockGenericIndexMockRecorder struct {
	mock *MockGenericIndex
}

// NewMockGenericIndex creates a new mock instance.
func NewMockGenericIndex(ctrl *gomock.Controller) *MockGenericIndex {
	mock := &MockGenericIndex{ctrl: ctrl}
	mock.recorder = &MockGenericIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenericIndex) EXPECT() *MockGenericIndexMockRecorder {
	return m.recorder
}

// Dimension mocks base method.
func (m *MockGenericIndex) Dimension() uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dimension")
	ret0, _ := ret[0].(uint16)
	return ret0
}

// Dimension indicates an expected call of Dimension.
func (mr *MockGenericIndexMockRecorder) Dimension() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dimension", reflect.TypeOf((*MockGenericIndex)(nil).Dimension))
}

// IndexType mocks base method.
func (m *MockGenericIndex) IndexType() factory.IndexType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexType")
	ret0, _ := ret[0].(factory.IndexType)
	return ret0
}

// IndexType indicates an expected call of IndexType.
func (mr *MockGenericIndexMockRecorder) IndexType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexType", reflect.TypeOf((*MockGenericIndex)(nil).IndexType))
}

// Method mocks base method.
func (m *MockGenericIndex) Method() factory.MethodType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(factory.MethodType)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockGenericIndexMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockGenericIndex)(nil).Method))
}

// Parameters mocks base method.
func (m *MockGenericIndex) Parameters() map[string]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]int)
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *MockGenericIndexMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MockGenericIndex)(nil).Parameters))
}
