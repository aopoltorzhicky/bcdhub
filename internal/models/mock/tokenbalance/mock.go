// Code generated by MockGen. DO NOT EDIT.
// Source: tokenbalance/repository.go

// Package tokenbalance is a generated GoMock package.
package tokenbalance

import (
	tb "github.com/baking-bad/bcdhub/internal/models/tokenbalance"
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
func (m *MockRepository) Get(network types.Network, contract, address string, tokenID uint64) (tb.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", network, contract, address, tokenID)
	ret0, _ := ret[0].(tb.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(network, contract, address, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), network, contract, address, tokenID)
}

// GetHolders mocks base method
func (m *MockRepository) GetHolders(network types.Network, contract string, tokenID uint64) ([]tb.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHolders", network, contract, tokenID)
	ret0, _ := ret[0].([]tb.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHolders indicates an expected call of GetHolders
func (mr *MockRepositoryMockRecorder) GetHolders(network, contract, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHolders", reflect.TypeOf((*MockRepository)(nil).GetHolders), network, contract, tokenID)
}

// Batch mocks base method
func (m *MockRepository) Batch(network types.Network, addresses []string) (map[string][]tb.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Batch", network, addresses)
	ret0, _ := ret[0].(map[string][]tb.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Batch indicates an expected call of Batch
func (mr *MockRepositoryMockRecorder) Batch(network, addresses interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockRepository)(nil).Batch), network, addresses)
}

// CountByContract mocks base method
func (m *MockRepository) CountByContract(network types.Network, address string) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByContract", network, address)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByContract indicates an expected call of CountByContract
func (mr *MockRepositoryMockRecorder) CountByContract(network, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByContract", reflect.TypeOf((*MockRepository)(nil).CountByContract), network, address)
}

// TokenSupply mocks base method
func (m *MockRepository) TokenSupply(network types.Network, contract string, tokenID uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenSupply", network, contract, tokenID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenSupply indicates an expected call of TokenSupply
func (mr *MockRepositoryMockRecorder) TokenSupply(network, contract, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenSupply", reflect.TypeOf((*MockRepository)(nil).TokenSupply), network, contract, tokenID)
}
