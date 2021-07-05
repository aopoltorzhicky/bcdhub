// Code generated by MockGen. DO NOT EDIT.
// Source: contract/repository.go

// Package contract is a generated GoMock package.
package contract

import (
	contractModel "github.com/baking-bad/bcdhub/internal/models/contract"
	types "github.com/baking-bad/bcdhub/internal/models/types"
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
func (m *MockRepository) Get(network types.Network, address string) (contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", network, address)
	ret0, _ := ret[0].(contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(network, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), network, address)
}

// GetMany mocks base method
func (m *MockRepository) GetMany(by map[string]interface{}) ([]contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", by)
	ret0, _ := ret[0].([]contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany
func (mr *MockRepositoryMockRecorder) GetMany(by interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockRepository)(nil).GetMany), by)
}

// GetRandom mocks base method
func (m *MockRepository) GetRandom(network types.Network) (contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandom", network)
	ret0, _ := ret[0].(contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandom indicates an expected call of GetRandom
func (mr *MockRepositoryMockRecorder) GetRandom(network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandom", reflect.TypeOf((*MockRepository)(nil).GetRandom), network)
}

// GetTokens mocks base method
func (m *MockRepository) GetTokens(network types.Network, tokenInterface string, offset, size int64) ([]contractModel.Contract, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokens", network, tokenInterface, offset, size)
	ret0, _ := ret[0].([]contractModel.Contract)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTokens indicates an expected call of GetTokens
func (mr *MockRepositoryMockRecorder) GetTokens(network, tokenInterface, offset, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokens", reflect.TypeOf((*MockRepository)(nil).GetTokens), network, tokenInterface, offset, size)
}

// GetProjectsLastContract mocks base method
func (m *MockRepository) GetProjectsLastContract(c contractModel.Contract, size, offset int64) ([]*contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsLastContract", c, size, offset)
	ret0, _ := ret[0].([]*contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsLastContract indicates an expected call of GetProjectsLastContract
func (mr *MockRepositoryMockRecorder) GetProjectsLastContract(c, size, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsLastContract", reflect.TypeOf((*MockRepository)(nil).GetProjectsLastContract), c, size, offset)
}

// GetSameContracts mocks base method
func (m *MockRepository) GetSameContracts(contact contractModel.Contract, manager string, size, offset int64) (contractModel.SameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSameContracts", contact, manager, size, offset)
	ret0, _ := ret[0].(contractModel.SameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSameContracts indicates an expected call of GetSameContracts
func (mr *MockRepositoryMockRecorder) GetSameContracts(contact, manager, size, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSameContracts", reflect.TypeOf((*MockRepository)(nil).GetSameContracts), contact, manager, size, offset)
}

// GetSimilarContracts mocks base method
func (m *MockRepository) GetSimilarContracts(arg0 contractModel.Contract, arg1, arg2 int64) ([]contractModel.Similar, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSimilarContracts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]contractModel.Similar)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSimilarContracts indicates an expected call of GetSimilarContracts
func (mr *MockRepositoryMockRecorder) GetSimilarContracts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimilarContracts", reflect.TypeOf((*MockRepository)(nil).GetSimilarContracts), arg0, arg1, arg2)
}

// GetByIDs mocks base method
func (m *MockRepository) GetByIDs(ids ...int64) ([]contractModel.Contract, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByIDs", varargs...)
	ret0, _ := ret[0].([]contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs
func (mr *MockRepositoryMockRecorder) GetByIDs(ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockRepository)(nil).GetByIDs), ids...)
}

// Stats mocks base method
func (m *MockRepository) Stats(c contractModel.Contract) (contractModel.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats", c)
	ret0, _ := ret[0].(contractModel.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockRepositoryMockRecorder) Stats(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockRepository)(nil).Stats), c)
}

// GetProjectIDByHash mocks base method
func (m *MockRepository) GetProjectIDByHash(hash string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectIDByHash", hash)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectIDByHash indicates an expected call of GetProjectIDByHash
func (mr *MockRepositoryMockRecorder) GetProjectIDByHash(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectIDByHash", reflect.TypeOf((*MockRepository)(nil).GetProjectIDByHash), hash)
}
