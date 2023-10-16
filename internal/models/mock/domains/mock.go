// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=../mock/domains/mock.go -package=domains -typed
//
// Package domains is a generated GoMock package.
package domains

import (
	context "context"
	reflect "reflect"

	contract "github.com/baking-bad/bcdhub/internal/models/contract"
	domains "github.com/baking-bad/bcdhub/internal/models/domains"
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

// Same mocks base method.
func (m *MockRepository) Same(ctx context.Context, network string, c contract.Contract, limit, offset int, availiableNetworks ...string) ([]domains.Same, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, network, c, limit, offset}
	for _, a := range availiableNetworks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Same", varargs...)
	ret0, _ := ret[0].([]domains.Same)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Same indicates an expected call of Same.
func (mr *MockRepositoryMockRecorder) Same(ctx, network, c, limit, offset any, availiableNetworks ...any) *RepositorySameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, network, c, limit, offset}, availiableNetworks...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Same", reflect.TypeOf((*MockRepository)(nil).Same), varargs...)
	return &RepositorySameCall{Call: call}
}

// RepositorySameCall wrap *gomock.Call
type RepositorySameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c_2 *RepositorySameCall) Return(arg0 []domains.Same, arg1 error) *RepositorySameCall {
	c_2.Call = c_2.Call.Return(arg0, arg1)
	return c_2
}

// Do rewrite *gomock.Call.Do
func (c_2 *RepositorySameCall) Do(f func(context.Context, string, contract.Contract, int, int, ...string) ([]domains.Same, error)) *RepositorySameCall {
	c_2.Call = c_2.Call.Do(f)
	return c_2
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c_2 *RepositorySameCall) DoAndReturn(f func(context.Context, string, contract.Contract, int, int, ...string) ([]domains.Same, error)) *RepositorySameCall {
	c_2.Call = c_2.Call.DoAndReturn(f)
	return c_2
}

// SameCount mocks base method.
func (m *MockRepository) SameCount(ctx context.Context, c contract.Contract, availiableNetworks ...string) (int, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, c}
	for _, a := range availiableNetworks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SameCount", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SameCount indicates an expected call of SameCount.
func (mr *MockRepositoryMockRecorder) SameCount(ctx, c any, availiableNetworks ...any) *RepositorySameCountCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, c}, availiableNetworks...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SameCount", reflect.TypeOf((*MockRepository)(nil).SameCount), varargs...)
	return &RepositorySameCountCall{Call: call}
}

// RepositorySameCountCall wrap *gomock.Call
type RepositorySameCountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c_2 *RepositorySameCountCall) Return(arg0 int, arg1 error) *RepositorySameCountCall {
	c_2.Call = c_2.Call.Return(arg0, arg1)
	return c_2
}

// Do rewrite *gomock.Call.Do
func (c_2 *RepositorySameCountCall) Do(f func(context.Context, contract.Contract, ...string) (int, error)) *RepositorySameCountCall {
	c_2.Call = c_2.Call.Do(f)
	return c_2
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c_2 *RepositorySameCountCall) DoAndReturn(f func(context.Context, contract.Contract, ...string) (int, error)) *RepositorySameCountCall {
	c_2.Call = c_2.Call.DoAndReturn(f)
	return c_2
}
