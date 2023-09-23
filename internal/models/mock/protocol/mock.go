// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=../mock/protocol/mock.go -package=protocol -typed
//
// Package protocol is a generated GoMock package.
package protocol

import (
	context "context"
	reflect "reflect"

	protocol "github.com/baking-bad/bcdhub/internal/models/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRepository) Get(ctx context.Context, hash string, level int64) (protocol.Protocol, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, hash, level)
	ret0, _ := ret[0].(protocol.Protocol)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, hash, level any) *RepositoryGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, hash, level)
	return &RepositoryGetCall{Call: call}
}

// RepositoryGetCall wrap *gomock.Call
type RepositoryGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryGetCall) Return(arg0 protocol.Protocol, arg1 error) *RepositoryGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryGetCall) Do(f func(context.Context, string, int64) (protocol.Protocol, error)) *RepositoryGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryGetCall) DoAndReturn(f func(context.Context, string, int64) (protocol.Protocol, error)) *RepositoryGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAll mocks base method.
func (m *MockRepository) GetAll(ctx context.Context) ([]protocol.Protocol, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]protocol.Protocol)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll(ctx any) *RepositoryGetAllCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll), ctx)
	return &RepositoryGetAllCall{Call: call}
}

// RepositoryGetAllCall wrap *gomock.Call
type RepositoryGetAllCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryGetAllCall) Return(response []protocol.Protocol, err error) *RepositoryGetAllCall {
	c.Call = c.Call.Return(response, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryGetAllCall) Do(f func(context.Context) ([]protocol.Protocol, error)) *RepositoryGetAllCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryGetAllCall) DoAndReturn(f func(context.Context) ([]protocol.Protocol, error)) *RepositoryGetAllCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockRepository) GetByID(ctx context.Context, id int64) (protocol.Protocol, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(protocol.Protocol)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(ctx, id any) *RepositoryGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), ctx, id)
	return &RepositoryGetByIDCall{Call: call}
}

// RepositoryGetByIDCall wrap *gomock.Call
type RepositoryGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryGetByIDCall) Return(response protocol.Protocol, err error) *RepositoryGetByIDCall {
	c.Call = c.Call.Return(response, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryGetByIDCall) Do(f func(context.Context, int64) (protocol.Protocol, error)) *RepositoryGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryGetByIDCall) DoAndReturn(f func(context.Context, int64) (protocol.Protocol, error)) *RepositoryGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByNetworkWithSort mocks base method.
func (m *MockRepository) GetByNetworkWithSort(ctx context.Context, sortField, order string) ([]protocol.Protocol, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNetworkWithSort", ctx, sortField, order)
	ret0, _ := ret[0].([]protocol.Protocol)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNetworkWithSort indicates an expected call of GetByNetworkWithSort.
func (mr *MockRepositoryMockRecorder) GetByNetworkWithSort(ctx, sortField, order any) *RepositoryGetByNetworkWithSortCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNetworkWithSort", reflect.TypeOf((*MockRepository)(nil).GetByNetworkWithSort), ctx, sortField, order)
	return &RepositoryGetByNetworkWithSortCall{Call: call}
}

// RepositoryGetByNetworkWithSortCall wrap *gomock.Call
type RepositoryGetByNetworkWithSortCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryGetByNetworkWithSortCall) Return(response []protocol.Protocol, err error) *RepositoryGetByNetworkWithSortCall {
	c.Call = c.Call.Return(response, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryGetByNetworkWithSortCall) Do(f func(context.Context, string, string) ([]protocol.Protocol, error)) *RepositoryGetByNetworkWithSortCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryGetByNetworkWithSortCall) DoAndReturn(f func(context.Context, string, string) ([]protocol.Protocol, error)) *RepositoryGetByNetworkWithSortCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
