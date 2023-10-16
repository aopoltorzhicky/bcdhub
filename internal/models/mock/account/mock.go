// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=../mock/account/mock.go -package=account -typed
//
// Package account is a generated GoMock package.
package account

import (
	context "context"
	reflect "reflect"

	account "github.com/baking-bad/bcdhub/internal/models/account"
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
func (m *MockRepository) Get(ctx context.Context, address string) (account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, address)
	ret0, _ := ret[0].(account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, address any) *RepositoryGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, address)
	return &RepositoryGetCall{Call: call}
}

// RepositoryGetCall wrap *gomock.Call
type RepositoryGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryGetCall) Return(arg0 account.Account, arg1 error) *RepositoryGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryGetCall) Do(f func(context.Context, string) (account.Account, error)) *RepositoryGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryGetCall) DoAndReturn(f func(context.Context, string) (account.Account, error)) *RepositoryGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RecentlyCalledContracts mocks base method.
func (m *MockRepository) RecentlyCalledContracts(ctx context.Context, offset, size int64) ([]account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecentlyCalledContracts", ctx, offset, size)
	ret0, _ := ret[0].([]account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecentlyCalledContracts indicates an expected call of RecentlyCalledContracts.
func (mr *MockRepositoryMockRecorder) RecentlyCalledContracts(ctx, offset, size any) *RepositoryRecentlyCalledContractsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecentlyCalledContracts", reflect.TypeOf((*MockRepository)(nil).RecentlyCalledContracts), ctx, offset, size)
	return &RepositoryRecentlyCalledContractsCall{Call: call}
}

// RepositoryRecentlyCalledContractsCall wrap *gomock.Call
type RepositoryRecentlyCalledContractsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryRecentlyCalledContractsCall) Return(accounts []account.Account, err error) *RepositoryRecentlyCalledContractsCall {
	c.Call = c.Call.Return(accounts, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryRecentlyCalledContractsCall) Do(f func(context.Context, int64, int64) ([]account.Account, error)) *RepositoryRecentlyCalledContractsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryRecentlyCalledContractsCall) DoAndReturn(f func(context.Context, int64, int64) ([]account.Account, error)) *RepositoryRecentlyCalledContractsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
