// Code generated by MockGen. DO NOT EDIT.
// Source: script_saver.go

// Package mock_contract is a generated GoMock package.
package contract

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockScriptSaver is a mock of ScriptSaver interface
type MockScriptSaver struct {
	ctrl     *gomock.Controller
	recorder *MockScriptSaverMockRecorder
}

// MockScriptSaverMockRecorder is the mock recorder for MockScriptSaver
type MockScriptSaverMockRecorder struct {
	mock *MockScriptSaver
}

// NewMockScriptSaver creates a new mock instance
func NewMockScriptSaver(ctrl *gomock.Controller) *MockScriptSaver {
	mock := &MockScriptSaver{ctrl: ctrl}
	mock.recorder = &MockScriptSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScriptSaver) EXPECT() *MockScriptSaverMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockScriptSaver) Save(code []byte, ctx ScriptSaveContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", code, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockScriptSaverMockRecorder) Save(code, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockScriptSaver)(nil).Save), code, ctx)
}
