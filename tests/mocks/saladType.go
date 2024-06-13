// Code generated by MockGen. DO NOT EDIT.
// Source: domain/saladType.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "github.com/Mx1q/ppo_services/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockISaladTypeRepository is a mock of ISaladTypeRepository interface.
type MockISaladTypeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockISaladTypeRepositoryMockRecorder
}

// MockISaladTypeRepositoryMockRecorder is the mock recorder for MockISaladTypeRepository.
type MockISaladTypeRepositoryMockRecorder struct {
	mock *MockISaladTypeRepository
}

// NewMockISaladTypeRepository creates a new mock instance.
func NewMockISaladTypeRepository(ctrl *gomock.Controller) *MockISaladTypeRepository {
	mock := &MockISaladTypeRepository{ctrl: ctrl}
	mock.recorder = &MockISaladTypeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISaladTypeRepository) EXPECT() *MockISaladTypeRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockISaladTypeRepository) Create(ctx context.Context, saladType *domain.SaladType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, saladType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockISaladTypeRepositoryMockRecorder) Create(ctx, saladType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockISaladTypeRepository)(nil).Create), ctx, saladType)
}

// DeleteById mocks base method.
func (m *MockISaladTypeRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockISaladTypeRepositoryMockRecorder) DeleteById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockISaladTypeRepository)(nil).DeleteById), ctx, id)
}

// GetAll mocks base method.
func (m *MockISaladTypeRepository) GetAll(ctx context.Context, page int) ([]*domain.SaladType, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, page)
	ret0, _ := ret[0].([]*domain.SaladType)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAll indicates an expected call of GetAll.
func (mr *MockISaladTypeRepositoryMockRecorder) GetAll(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockISaladTypeRepository)(nil).GetAll), ctx, page)
}

// GetAllBySaladId mocks base method.
func (m *MockISaladTypeRepository) GetAllBySaladId(ctx context.Context, saladId uuid.UUID) ([]*domain.SaladType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBySaladId", ctx, saladId)
	ret0, _ := ret[0].([]*domain.SaladType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBySaladId indicates an expected call of GetAllBySaladId.
func (mr *MockISaladTypeRepositoryMockRecorder) GetAllBySaladId(ctx, saladId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBySaladId", reflect.TypeOf((*MockISaladTypeRepository)(nil).GetAllBySaladId), ctx, saladId)
}

// GetById mocks base method.
func (m *MockISaladTypeRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.SaladType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*domain.SaladType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockISaladTypeRepositoryMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockISaladTypeRepository)(nil).GetById), ctx, id)
}

// Link mocks base method.
func (m *MockISaladTypeRepository) Link(ctx context.Context, saladId, saladTypeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Link", ctx, saladId, saladTypeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Link indicates an expected call of Link.
func (mr *MockISaladTypeRepositoryMockRecorder) Link(ctx, saladId, saladTypeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Link", reflect.TypeOf((*MockISaladTypeRepository)(nil).Link), ctx, saladId, saladTypeId)
}

// Unlink mocks base method.
func (m *MockISaladTypeRepository) Unlink(ctx context.Context, saladId, saladTypeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlink", ctx, saladId, saladTypeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlink indicates an expected call of Unlink.
func (mr *MockISaladTypeRepositoryMockRecorder) Unlink(ctx, saladId, saladTypeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlink", reflect.TypeOf((*MockISaladTypeRepository)(nil).Unlink), ctx, saladId, saladTypeId)
}

// Update mocks base method.
func (m *MockISaladTypeRepository) Update(ctx context.Context, saladType *domain.SaladType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, saladType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockISaladTypeRepositoryMockRecorder) Update(ctx, saladType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockISaladTypeRepository)(nil).Update), ctx, saladType)
}

// MockISaladTypeService is a mock of ISaladTypeService interface.
type MockISaladTypeService struct {
	ctrl     *gomock.Controller
	recorder *MockISaladTypeServiceMockRecorder
}

// MockISaladTypeServiceMockRecorder is the mock recorder for MockISaladTypeService.
type MockISaladTypeServiceMockRecorder struct {
	mock *MockISaladTypeService
}

// NewMockISaladTypeService creates a new mock instance.
func NewMockISaladTypeService(ctrl *gomock.Controller) *MockISaladTypeService {
	mock := &MockISaladTypeService{ctrl: ctrl}
	mock.recorder = &MockISaladTypeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISaladTypeService) EXPECT() *MockISaladTypeServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockISaladTypeService) Create(ctx context.Context, saladType *domain.SaladType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, saladType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockISaladTypeServiceMockRecorder) Create(ctx, saladType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockISaladTypeService)(nil).Create), ctx, saladType)
}

// DeleteById mocks base method.
func (m *MockISaladTypeService) DeleteById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockISaladTypeServiceMockRecorder) DeleteById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockISaladTypeService)(nil).DeleteById), ctx, id)
}

// GetAll mocks base method.
func (m *MockISaladTypeService) GetAll(ctx context.Context, page int) ([]*domain.SaladType, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, page)
	ret0, _ := ret[0].([]*domain.SaladType)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAll indicates an expected call of GetAll.
func (mr *MockISaladTypeServiceMockRecorder) GetAll(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockISaladTypeService)(nil).GetAll), ctx, page)
}

// GetAllBySaladId mocks base method.
func (m *MockISaladTypeService) GetAllBySaladId(ctx context.Context, saladId uuid.UUID) ([]*domain.SaladType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBySaladId", ctx, saladId)
	ret0, _ := ret[0].([]*domain.SaladType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBySaladId indicates an expected call of GetAllBySaladId.
func (mr *MockISaladTypeServiceMockRecorder) GetAllBySaladId(ctx, saladId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBySaladId", reflect.TypeOf((*MockISaladTypeService)(nil).GetAllBySaladId), ctx, saladId)
}

// GetById mocks base method.
func (m *MockISaladTypeService) GetById(ctx context.Context, id uuid.UUID) (*domain.SaladType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*domain.SaladType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockISaladTypeServiceMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockISaladTypeService)(nil).GetById), ctx, id)
}

// Link mocks base method.
func (m *MockISaladTypeService) Link(ctx context.Context, saladId, saladTypeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Link", ctx, saladId, saladTypeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Link indicates an expected call of Link.
func (mr *MockISaladTypeServiceMockRecorder) Link(ctx, saladId, saladTypeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Link", reflect.TypeOf((*MockISaladTypeService)(nil).Link), ctx, saladId, saladTypeId)
}

// Unlink mocks base method.
func (m *MockISaladTypeService) Unlink(ctx context.Context, saladId, saladTypeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlink", ctx, saladId, saladTypeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlink indicates an expected call of Unlink.
func (mr *MockISaladTypeServiceMockRecorder) Unlink(ctx, saladId, saladTypeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlink", reflect.TypeOf((*MockISaladTypeService)(nil).Unlink), ctx, saladId, saladTypeId)
}

// Update mocks base method.
func (m *MockISaladTypeService) Update(ctx context.Context, measurement *domain.SaladType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, measurement)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockISaladTypeServiceMockRecorder) Update(ctx, measurement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockISaladTypeService)(nil).Update), ctx, measurement)
}