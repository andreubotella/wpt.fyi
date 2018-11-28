// Code generated by MockGen. DO NOT EDIT.
// Source: api/query/cache/index/index.go

// Package index is a generated GoMock package.
package index

import (
	gomock "github.com/golang/mock/gomock"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
	reflect "reflect"
)

// MockIndex is a mock of Index interface
type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *MockIndexMockRecorder
}

// MockIndexMockRecorder is the mock recorder for MockIndex
type MockIndexMockRecorder struct {
	mock *MockIndex
}

// NewMockIndex creates a new mock instance
func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &MockIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndex) EXPECT() *MockIndexMockRecorder {
	return m.recorder
}

// IngestRun mocks base method
func (m *MockIndex) IngestRun(arg0 shared.TestRun) error {
	ret := m.ctrl.Call(m, "IngestRun", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// IngestRun indicates an expected call of IngestRun
func (mr *MockIndexMockRecorder) IngestRun(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestRun", reflect.TypeOf((*MockIndex)(nil).IngestRun), arg0)
}

// EvictAnyRun mocks base method
func (m *MockIndex) EvictAnyRun() error {
	ret := m.ctrl.Call(m, "EvictAnyRun")
	ret0, _ := ret[0].(error)
	return ret0
}

// EvictAnyRun indicates an expected call of EvictAnyRun
func (mr *MockIndexMockRecorder) EvictAnyRun() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictAnyRun", reflect.TypeOf((*MockIndex)(nil).EvictAnyRun))
}
