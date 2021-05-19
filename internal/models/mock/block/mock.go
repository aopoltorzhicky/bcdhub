// Code generated by MockGen. DO NOT EDIT.
// Source: block/repository.go

// Package block is a generated GoMock package.
package block

import (
	modelBlock "github.com/baking-bad/bcdhub/internal/models/block"
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
func (m *MockRepository) Get(network types.Network, level int64) (modelBlock.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", network, level)
	ret0, _ := ret[0].(modelBlock.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(network, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), network, level)
}

// Last mocks base method
func (m *MockRepository) Last(network types.Network) (modelBlock.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last", network)
	ret0, _ := ret[0].(modelBlock.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Last indicates an expected call of Last
func (mr *MockRepositoryMockRecorder) Last(network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockRepository)(nil).Last), network)
}

// LastByNetworks mocks base method
func (m *MockRepository) LastByNetworks() ([]modelBlock.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastByNetworks")
	ret0, _ := ret[0].([]modelBlock.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastByNetworks indicates an expected call of LastByNetworks
func (mr *MockRepositoryMockRecorder) LastByNetworks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastByNetworks", reflect.TypeOf((*MockRepository)(nil).LastByNetworks))
}

// GetNetworkAlias mocks base method
func (m *MockRepository) GetNetworkAlias(chainID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkAlias", chainID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkAlias indicates an expected call of GetNetworkAlias
func (mr *MockRepositoryMockRecorder) GetNetworkAlias(chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkAlias", reflect.TypeOf((*MockRepository)(nil).GetNetworkAlias), chainID)
}
