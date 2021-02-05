// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package noderpc is a generated GoMock package.
package noderpc

import (
	gomock "github.com/golang/mock/gomock"
	gjson "github.com/tidwall/gjson"
	reflect "reflect"
	time "time"
)

// MockINode is a mock of INode interface
type MockINode struct {
	ctrl     *gomock.Controller
	recorder *MockINodeMockRecorder
}

// MockINodeMockRecorder is the mock recorder for MockINode
type MockINodeMockRecorder struct {
	mock *MockINode
}

// NewMockINode creates a new mock instance
func NewMockINode(ctrl *gomock.Controller) *MockINode {
	mock := &MockINode{ctrl: ctrl}
	mock.recorder = &MockINodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockINode) EXPECT() *MockINodeMockRecorder {
	return m.recorder
}

// GetHead mocks base method
func (m *MockINode) GetHead() (Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead")
	ret0, _ := ret[0].(Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead
func (mr *MockINodeMockRecorder) GetHead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockINode)(nil).GetHead))
}

// GetHeader mocks base method
func (m *MockINode) GetHeader(arg0 int64) (Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0)
	ret0, _ := ret[0].(Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockINodeMockRecorder) GetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockINode)(nil).GetHeader), arg0)
}

// GetLevel mocks base method
func (m *MockINode) GetLevel() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLevel")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLevel indicates an expected call of GetLevel
func (mr *MockINodeMockRecorder) GetLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLevel", reflect.TypeOf((*MockINode)(nil).GetLevel))
}

// GetLevelTime mocks base method
func (m *MockINode) GetLevelTime(arg0 int) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLevelTime", arg0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLevelTime indicates an expected call of GetLevelTime
func (mr *MockINodeMockRecorder) GetLevelTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLevelTime", reflect.TypeOf((*MockINode)(nil).GetLevelTime), arg0)
}

// GetScriptJSON mocks base method
func (m *MockINode) GetScriptJSON(arg0 string, arg1 int64) (gjson.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScriptJSON", arg0, arg1)
	ret0, _ := ret[0].(gjson.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScriptJSON indicates an expected call of GetScriptJSON
func (mr *MockINodeMockRecorder) GetScriptJSON(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScriptJSON", reflect.TypeOf((*MockINode)(nil).GetScriptJSON), arg0, arg1)
}

// GetScriptStorageJSON mocks base method
func (m *MockINode) GetScriptStorageJSON(arg0 string, arg1 int64) (gjson.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScriptStorageJSON", arg0, arg1)
	ret0, _ := ret[0].(gjson.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScriptStorageJSON indicates an expected call of GetScriptStorageJSON
func (mr *MockINodeMockRecorder) GetScriptStorageJSON(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScriptStorageJSON", reflect.TypeOf((*MockINode)(nil).GetScriptStorageJSON), arg0, arg1)
}

// GetContractBalance mocks base method
func (m *MockINode) GetContractBalance(arg0 string, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractBalance", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractBalance indicates an expected call of GetContractBalance
func (mr *MockINodeMockRecorder) GetContractBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractBalance", reflect.TypeOf((*MockINode)(nil).GetContractBalance), arg0, arg1)
}

// GetContractData mocks base method
func (m *MockINode) GetContractData(arg0 string, arg1 int64) (ContractData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractData", arg0, arg1)
	ret0, _ := ret[0].(ContractData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractData indicates an expected call of GetContractData
func (mr *MockINodeMockRecorder) GetContractData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractData", reflect.TypeOf((*MockINode)(nil).GetContractData), arg0, arg1)
}

// GetOperations mocks base method
func (m *MockINode) GetOperations(arg0 int64) (gjson.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperations", arg0)
	ret0, _ := ret[0].(gjson.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperations indicates an expected call of GetOperations
func (mr *MockINodeMockRecorder) GetOperations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperations", reflect.TypeOf((*MockINode)(nil).GetOperations), arg0)
}

// GetContractsByBlock mocks base method
func (m *MockINode) GetContractsByBlock(arg0 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractsByBlock", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractsByBlock indicates an expected call of GetContractsByBlock
func (mr *MockINodeMockRecorder) GetContractsByBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractsByBlock", reflect.TypeOf((*MockINode)(nil).GetContractsByBlock), arg0)
}

// GetNetworkConstants mocks base method
func (m *MockINode) GetNetworkConstants(arg0 int64) (Constants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkConstants", arg0)
	ret0, _ := ret[0].(Constants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkConstants indicates an expected call of GetNetworkConstants
func (mr *MockINodeMockRecorder) GetNetworkConstants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkConstants", reflect.TypeOf((*MockINode)(nil).GetNetworkConstants), arg0)
}

// RunCode mocks base method
func (m *MockINode) RunCode(arg0, arg1, arg2 gjson.Result, arg3, arg4, arg5, arg6, arg7 string, arg8, arg9 int64) (gjson.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCode", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	ret0, _ := ret[0].(gjson.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunCode indicates an expected call of RunCode
func (mr *MockINodeMockRecorder) RunCode(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCode", reflect.TypeOf((*MockINode)(nil).RunCode), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RunOperation mocks base method
func (m *MockINode) RunOperation(arg0, arg1, arg2, arg3 string, arg4, arg5, arg6, arg7, arg8 int64, arg9 gjson.Result) (gjson.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunOperation", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	ret0, _ := ret[0].(gjson.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunOperation indicates an expected call of RunOperation
func (mr *MockINodeMockRecorder) RunOperation(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunOperation", reflect.TypeOf((*MockINode)(nil).RunOperation), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// GetCounter mocks base method
func (m *MockINode) GetCounter(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounter", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounter indicates an expected call of GetCounter
func (mr *MockINodeMockRecorder) GetCounter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounter", reflect.TypeOf((*MockINode)(nil).GetCounter), arg0)
}

// GetCode mocks base method
func (m *MockINode) GetCode(address string, level int64) (gjson.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", address, level)
	ret0, _ := ret[0].(gjson.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCode indicates an expected call of GetCode
func (mr *MockINodeMockRecorder) GetCode(address, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockINode)(nil).GetCode), address, level)
}
