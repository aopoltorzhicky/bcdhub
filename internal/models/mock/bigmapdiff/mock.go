// Code generated by MockGen. DO NOT EDIT.
// Source: bigmapdiff/repository.go

// Package bigmapdiff is a generated GoMock package.
package bigmapdiff

import (
	bmd "github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(ctx bmd.GetContext) ([]bmd.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].([]bmd.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx)
}

// GetByAddress mocks base method
func (m *MockRepository) GetByAddress(arg0, arg1 string) ([]bmd.BigMapDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAddress", arg0, arg1)
	ret0, _ := ret[0].([]bmd.BigMapDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAddress indicates an expected call of GetByAddress
func (mr *MockRepositoryMockRecorder) GetByAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAddress", reflect.TypeOf((*MockRepository)(nil).GetByAddress), arg0, arg1)
}

// GetForOperation mocks base method
func (m *MockRepository) GetForOperation(hash string, counter int64, nonce *int64) ([]*bmd.BigMapDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForOperation", hash, counter, nonce)
	ret0, _ := ret[0].([]*bmd.BigMapDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForOperation indicates an expected call of GetForOperation
func (mr *MockRepositoryMockRecorder) GetForOperation(hash, counter, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForOperation", reflect.TypeOf((*MockRepository)(nil).GetForOperation), hash, counter, nonce)
}

// GetUniqueForOperations mocks base method
func (m *MockRepository) GetUniqueForOperations(opg []bmd.OPG) ([]bmd.BigMapDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUniqueForOperations", opg)
	ret0, _ := ret[0].([]bmd.BigMapDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUniqueForOperations indicates an expected call of GetUniqueForOperations
func (mr *MockRepositoryMockRecorder) GetUniqueForOperations(opg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUniqueForOperations", reflect.TypeOf((*MockRepository)(nil).GetUniqueForOperations), opg)
}

// GetByPtr mocks base method
func (m *MockRepository) GetByPtr(network, contract string, ptr int64) ([]bmd.BigMapState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPtr", network, contract, ptr)
	ret0, _ := ret[0].([]bmd.BigMapState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPtr indicates an expected call of GetByPtr
func (mr *MockRepositoryMockRecorder) GetByPtr(network, contract, ptr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPtr", reflect.TypeOf((*MockRepository)(nil).GetByPtr), network, contract, ptr)
}

// GetByPtrAndKeyHash mocks base method
func (m *MockRepository) GetByPtrAndKeyHash(arg0 int64, arg1, arg2 string, arg3, arg4 int64) ([]bmd.BigMapDiff, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPtrAndKeyHash", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]bmd.BigMapDiff)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByPtrAndKeyHash indicates an expected call of GetByPtrAndKeyHash
func (mr *MockRepositoryMockRecorder) GetByPtrAndKeyHash(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPtrAndKeyHash", reflect.TypeOf((*MockRepository)(nil).GetByPtrAndKeyHash), arg0, arg1, arg2, arg3, arg4)
}

// GetForAddress mocks base method
func (m *MockRepository) GetForAddress(arg0 string) ([]bmd.BigMapDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForAddress", arg0)
	ret0, _ := ret[0].([]bmd.BigMapDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetForAddress indicates an expected call of GetForAddress
func (mr *MockRepositoryMockRecorder) GetForAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForAddress", reflect.TypeOf((*MockRepository)(nil).GetForAddress), arg0)
}

// GetByIDs mocks base method
func (m *MockRepository) GetByIDs(ids ...int64) ([]bmd.BigMapDiff, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByIDs", varargs...)
	ret0, _ := ret[0].([]bmd.BigMapDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs
func (mr *MockRepositoryMockRecorder) GetByIDs(ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockRepository)(nil).GetByIDs), ids...)
}

// GetValuesByKey mocks base method
func (m *MockRepository) GetValuesByKey(arg0 string) ([]bmd.BigMapDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValuesByKey", arg0)
	ret0, _ := ret[0].([]bmd.BigMapDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValuesByKey indicates an expected call of GetValuesByKey
func (mr *MockRepositoryMockRecorder) GetValuesByKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValuesByKey", reflect.TypeOf((*MockRepository)(nil).GetValuesByKey), arg0)
}

// Count mocks base method
func (m *MockRepository) Count(network string, ptr int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", network, ptr)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockRepositoryMockRecorder) Count(network, ptr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRepository)(nil).Count), network, ptr)
}

// Current mocks base method
func (m *MockRepository) Current(network, keyHash string, ptr int64) (bmd.BigMapState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current", network, keyHash, ptr)
	ret0, _ := ret[0].(bmd.BigMapState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Current indicates an expected call of Current
func (mr *MockRepositoryMockRecorder) Current(network, keyHash, ptr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockRepository)(nil).Current), network, keyHash, ptr)
}

// CurrentByContract mocks base method
func (m *MockRepository) CurrentByContract(network, contract string) ([]bmd.BigMapState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentByContract", network, contract)
	ret0, _ := ret[0].([]bmd.BigMapState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentByContract indicates an expected call of CurrentByContract
func (mr *MockRepositoryMockRecorder) CurrentByContract(network, contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentByContract", reflect.TypeOf((*MockRepository)(nil).CurrentByContract), network, contract)
}

// Previous mocks base method
func (m *MockRepository) Previous(arg0 []bmd.BigMapDiff) ([]bmd.BigMapDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Previous", arg0)
	ret0, _ := ret[0].([]bmd.BigMapDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Previous indicates an expected call of Previous
func (mr *MockRepositoryMockRecorder) Previous(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Previous", reflect.TypeOf((*MockRepository)(nil).Previous), arg0)
}

// GetStats mocks base method
func (m *MockRepository) GetStats(network string, ptr int64) (bmd.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", network, ptr)
	ret0, _ := ret[0].(bmd.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats
func (mr *MockRepositoryMockRecorder) GetStats(network, ptr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockRepository)(nil).GetStats), network, ptr)
}

// StatesChangedAfter mocks base method
func (m *MockRepository) StatesChangedAfter(network string, level int64) ([]bmd.BigMapState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatesChangedAfter", network, level)
	ret0, _ := ret[0].([]bmd.BigMapState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatesChangedAfter indicates an expected call of StatesChangedAfter
func (mr *MockRepositoryMockRecorder) StatesChangedAfter(network, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatesChangedAfter", reflect.TypeOf((*MockRepository)(nil).StatesChangedAfter), network, level)
}

// LastDiff mocks base method
func (m *MockRepository) LastDiff(network string, ptr int64, keyHash string, skipRemoved bool) (bmd.BigMapDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastDiff", network, ptr, keyHash, skipRemoved)
	ret0, _ := ret[0].(bmd.BigMapDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastDiff indicates an expected call of LastDiff
func (mr *MockRepositoryMockRecorder) LastDiff(network, ptr, keyHash, skipRemoved interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastDiff", reflect.TypeOf((*MockRepository)(nil).LastDiff), network, ptr, keyHash, skipRemoved)
}

// Keys mocks base method
func (m *MockRepository) Keys(ctx bmd.GetContext) ([]bmd.BigMapState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys", ctx)
	ret0, _ := ret[0].([]bmd.BigMapState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Keys indicates an expected call of Keys
func (mr *MockRepositoryMockRecorder) Keys(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRepository)(nil).Keys), ctx)
}
