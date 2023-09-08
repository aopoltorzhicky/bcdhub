// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mock/general.go -package=mock -typed
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/baking-bad/bcdhub/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockGeneralRepository is a mock of GeneralRepository interface.
type MockGeneralRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGeneralRepositoryMockRecorder
}

// MockGeneralRepositoryMockRecorder is the mock recorder for MockGeneralRepository.
type MockGeneralRepositoryMockRecorder struct {
	mock *MockGeneralRepository
}

// NewMockGeneralRepository creates a new mock instance.
func NewMockGeneralRepository(ctrl *gomock.Controller) *MockGeneralRepository {
	mock := &MockGeneralRepository{ctrl: ctrl}
	mock.recorder = &MockGeneralRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeneralRepository) EXPECT() *MockGeneralRepositoryMockRecorder {
	return m.recorder
}

// BulkDelete mocks base method.
func (m *MockGeneralRepository) BulkDelete(arg0 context.Context, arg1 []models.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDelete indicates an expected call of BulkDelete.
func (mr *MockGeneralRepositoryMockRecorder) BulkDelete(arg0, arg1 any) *GeneralRepositoryBulkDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDelete", reflect.TypeOf((*MockGeneralRepository)(nil).BulkDelete), arg0, arg1)
	return &GeneralRepositoryBulkDeleteCall{Call: call}
}

// GeneralRepositoryBulkDeleteCall wrap *gomock.Call
type GeneralRepositoryBulkDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryBulkDeleteCall) Return(arg0 error) *GeneralRepositoryBulkDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryBulkDeleteCall) Do(f func(context.Context, []models.Model) error) *GeneralRepositoryBulkDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryBulkDeleteCall) DoAndReturn(f func(context.Context, []models.Model) error) *GeneralRepositoryBulkDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateTables mocks base method.
func (m *MockGeneralRepository) CreateTables() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTables")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTables indicates an expected call of CreateTables.
func (mr *MockGeneralRepositoryMockRecorder) CreateTables() *GeneralRepositoryCreateTablesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTables", reflect.TypeOf((*MockGeneralRepository)(nil).CreateTables))
	return &GeneralRepositoryCreateTablesCall{Call: call}
}

// GeneralRepositoryCreateTablesCall wrap *gomock.Call
type GeneralRepositoryCreateTablesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryCreateTablesCall) Return(arg0 error) *GeneralRepositoryCreateTablesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryCreateTablesCall) Do(f func() error) *GeneralRepositoryCreateTablesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryCreateTablesCall) DoAndReturn(f func() error) *GeneralRepositoryCreateTablesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteByContract mocks base method.
func (m *MockGeneralRepository) DeleteByContract(indices []string, address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByContract", indices, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByContract indicates an expected call of DeleteByContract.
func (mr *MockGeneralRepositoryMockRecorder) DeleteByContract(indices, address any) *GeneralRepositoryDeleteByContractCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByContract", reflect.TypeOf((*MockGeneralRepository)(nil).DeleteByContract), indices, address)
	return &GeneralRepositoryDeleteByContractCall{Call: call}
}

// GeneralRepositoryDeleteByContractCall wrap *gomock.Call
type GeneralRepositoryDeleteByContractCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryDeleteByContractCall) Return(arg0 error) *GeneralRepositoryDeleteByContractCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryDeleteByContractCall) Do(f func([]string, string) error) *GeneralRepositoryDeleteByContractCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryDeleteByContractCall) DoAndReturn(f func([]string, string) error) *GeneralRepositoryDeleteByContractCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Drop mocks base method.
func (m *MockGeneralRepository) Drop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockGeneralRepositoryMockRecorder) Drop(ctx any) *GeneralRepositoryDropCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockGeneralRepository)(nil).Drop), ctx)
	return &GeneralRepositoryDropCall{Call: call}
}

// GeneralRepositoryDropCall wrap *gomock.Call
type GeneralRepositoryDropCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryDropCall) Return(arg0 error) *GeneralRepositoryDropCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryDropCall) Do(f func(context.Context) error) *GeneralRepositoryDropCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryDropCall) DoAndReturn(f func(context.Context) error) *GeneralRepositoryDropCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAll mocks base method.
func (m *MockGeneralRepository) GetAll(index string) ([]models.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", index)
	ret0, _ := ret[0].([]models.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockGeneralRepositoryMockRecorder) GetAll(index any) *GeneralRepositoryGetAllCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockGeneralRepository)(nil).GetAll), index)
	return &GeneralRepositoryGetAllCall{Call: call}
}

// GeneralRepositoryGetAllCall wrap *gomock.Call
type GeneralRepositoryGetAllCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryGetAllCall) Return(arg0 []models.Model, arg1 error) *GeneralRepositoryGetAllCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryGetAllCall) Do(f func(string) ([]models.Model, error)) *GeneralRepositoryGetAllCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryGetAllCall) DoAndReturn(f func(string) ([]models.Model, error)) *GeneralRepositoryGetAllCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockGeneralRepository) GetByID(output models.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", output)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByID indicates an expected call of GetByID.
func (mr *MockGeneralRepositoryMockRecorder) GetByID(output any) *GeneralRepositoryGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockGeneralRepository)(nil).GetByID), output)
	return &GeneralRepositoryGetByIDCall{Call: call}
}

// GeneralRepositoryGetByIDCall wrap *gomock.Call
type GeneralRepositoryGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryGetByIDCall) Return(arg0 error) *GeneralRepositoryGetByIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryGetByIDCall) Do(f func(models.Model) error) *GeneralRepositoryGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryGetByIDCall) DoAndReturn(f func(models.Model) error) *GeneralRepositoryGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsRecordNotFound mocks base method.
func (m *MockGeneralRepository) IsRecordNotFound(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRecordNotFound", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRecordNotFound indicates an expected call of IsRecordNotFound.
func (mr *MockGeneralRepositoryMockRecorder) IsRecordNotFound(err any) *GeneralRepositoryIsRecordNotFoundCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRecordNotFound", reflect.TypeOf((*MockGeneralRepository)(nil).IsRecordNotFound), err)
	return &GeneralRepositoryIsRecordNotFoundCall{Call: call}
}

// GeneralRepositoryIsRecordNotFoundCall wrap *gomock.Call
type GeneralRepositoryIsRecordNotFoundCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryIsRecordNotFoundCall) Return(arg0 bool) *GeneralRepositoryIsRecordNotFoundCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryIsRecordNotFoundCall) Do(f func(error) bool) *GeneralRepositoryIsRecordNotFoundCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryIsRecordNotFoundCall) DoAndReturn(f func(error) bool) *GeneralRepositoryIsRecordNotFoundCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m *MockGeneralRepository) Save(ctx context.Context, items []models.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, items)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockGeneralRepositoryMockRecorder) Save(ctx, items any) *GeneralRepositorySaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockGeneralRepository)(nil).Save), ctx, items)
	return &GeneralRepositorySaveCall{Call: call}
}

// GeneralRepositorySaveCall wrap *gomock.Call
type GeneralRepositorySaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositorySaveCall) Return(arg0 error) *GeneralRepositorySaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositorySaveCall) Do(f func(context.Context, []models.Model) error) *GeneralRepositorySaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositorySaveCall) DoAndReturn(f func(context.Context, []models.Model) error) *GeneralRepositorySaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TablesExist mocks base method.
func (m *MockGeneralRepository) TablesExist() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TablesExist")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TablesExist indicates an expected call of TablesExist.
func (mr *MockGeneralRepositoryMockRecorder) TablesExist() *GeneralRepositoryTablesExistCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TablesExist", reflect.TypeOf((*MockGeneralRepository)(nil).TablesExist))
	return &GeneralRepositoryTablesExistCall{Call: call}
}

// GeneralRepositoryTablesExistCall wrap *gomock.Call
type GeneralRepositoryTablesExistCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryTablesExistCall) Return(arg0 bool) *GeneralRepositoryTablesExistCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryTablesExistCall) Do(f func() bool) *GeneralRepositoryTablesExistCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryTablesExistCall) DoAndReturn(f func() bool) *GeneralRepositoryTablesExistCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateDoc mocks base method.
func (m *MockGeneralRepository) UpdateDoc(model models.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDoc", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDoc indicates an expected call of UpdateDoc.
func (mr *MockGeneralRepositoryMockRecorder) UpdateDoc(model any) *GeneralRepositoryUpdateDocCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDoc", reflect.TypeOf((*MockGeneralRepository)(nil).UpdateDoc), model)
	return &GeneralRepositoryUpdateDocCall{Call: call}
}

// GeneralRepositoryUpdateDocCall wrap *gomock.Call
type GeneralRepositoryUpdateDocCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *GeneralRepositoryUpdateDocCall) Return(err error) *GeneralRepositoryUpdateDocCall {
	c.Call = c.Call.Return(err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *GeneralRepositoryUpdateDocCall) Do(f func(models.Model) error) *GeneralRepositoryUpdateDocCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *GeneralRepositoryUpdateDocCall) DoAndReturn(f func(models.Model) error) *GeneralRepositoryUpdateDocCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
