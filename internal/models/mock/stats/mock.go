// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=../mock/stats/mock.go -package=stats -typed
//
// Package stats is a generated GoMock package.
package stats

import (
	context "context"
	reflect "reflect"

	stats "github.com/baking-bad/bcdhub/internal/models/stats"
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
func (m *MockRepository) Get(ctx context.Context) (stats.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(stats.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx any) *RepositoryGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx)
	return &RepositoryGetCall{Call: call}
}

// RepositoryGetCall wrap *gomock.Call
type RepositoryGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryGetCall) Return(arg0 stats.Stats, arg1 error) *RepositoryGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryGetCall) Do(f func(context.Context) (stats.Stats, error)) *RepositoryGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryGetCall) DoAndReturn(f func(context.Context) (stats.Stats, error)) *RepositoryGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
