// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source repository.go -destination mock/repository_mock.go -package mock
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockcommonRepo is a mock of commonRepo interface.
type MockcommonRepo struct {
	ctrl     *gomock.Controller
	recorder *MockcommonRepoMockRecorder
}

// MockcommonRepoMockRecorder is the mock recorder for MockcommonRepo.
type MockcommonRepoMockRecorder struct {
	mock *MockcommonRepo
}

// NewMockcommonRepo creates a new mock instance.
func NewMockcommonRepo(ctrl *gomock.Controller) *MockcommonRepo {
	mock := &MockcommonRepo{ctrl: ctrl}
	mock.recorder = &MockcommonRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcommonRepo) EXPECT() *MockcommonRepoMockRecorder {
	return m.recorder
}

// GetBlockedDuration mocks base method.
func (m *MockcommonRepo) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockedDuration", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockedDuration indicates an expected call of GetBlockedDuration.
func (mr *MockcommonRepoMockRecorder) GetBlockedDuration(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockedDuration", reflect.TypeOf((*MockcommonRepo)(nil).GetBlockedDuration), ctx, key)
}

// SaveBlockedDuration mocks base method.
func (m *MockcommonRepo) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBlockedDuration", ctx, key, BlockedDuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBlockedDuration indicates an expected call of SaveBlockedDuration.
func (mr *MockcommonRepoMockRecorder) SaveBlockedDuration(ctx, key, BlockedDuration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBlockedDuration", reflect.TypeOf((*MockcommonRepo)(nil).SaveBlockedDuration), ctx, key, BlockedDuration)
}

// UpsertRequest mocks base method.
func (m *MockcommonRepo) UpsertRequest(ctx context.Context, key string, rl entity.RateLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRequest", ctx, key, rl)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRequest indicates an expected call of UpsertRequest.
func (mr *MockcommonRepoMockRecorder) UpsertRequest(ctx, key, rl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRequest", reflect.TypeOf((*MockcommonRepo)(nil).UpsertRequest), ctx, key, rl)
}

// MockAPITokenRepository is a mock of APITokenRepository interface.
type MockAPITokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAPITokenRepositoryMockRecorder
}

// MockAPITokenRepositoryMockRecorder is the mock recorder for MockAPITokenRepository.
type MockAPITokenRepositoryMockRecorder struct {
	mock *MockAPITokenRepository
}

// NewMockAPITokenRepository creates a new mock instance.
func NewMockAPITokenRepository(ctrl *gomock.Controller) *MockAPITokenRepository {
	mock := &MockAPITokenRepository{ctrl: ctrl}
	mock.recorder = &MockAPITokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPITokenRepository) EXPECT() *MockAPITokenRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAPITokenRepository) Get(ctx context.Context, value string) entity.APIToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, value)
	ret0, _ := ret[0].(entity.APIToken)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockAPITokenRepositoryMockRecorder) Get(ctx, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAPITokenRepository)(nil).Get), ctx, value)
}

// GetBlockedDuration mocks base method.
func (m *MockAPITokenRepository) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockedDuration", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockedDuration indicates an expected call of GetBlockedDuration.
func (mr *MockAPITokenRepositoryMockRecorder) GetBlockedDuration(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockedDuration", reflect.TypeOf((*MockAPITokenRepository)(nil).GetBlockedDuration), ctx, key)
}

// Save mocks base method.
func (m *MockAPITokenRepository) Save(ctx context.Context, token *entity.APIToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockAPITokenRepositoryMockRecorder) Save(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAPITokenRepository)(nil).Save), ctx, token)
}

// SaveBlockedDuration mocks base method.
func (m *MockAPITokenRepository) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBlockedDuration", ctx, key, BlockedDuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBlockedDuration indicates an expected call of SaveBlockedDuration.
func (mr *MockAPITokenRepositoryMockRecorder) SaveBlockedDuration(ctx, key, BlockedDuration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBlockedDuration", reflect.TypeOf((*MockAPITokenRepository)(nil).SaveBlockedDuration), ctx, key, BlockedDuration)
}

// UpsertRequest mocks base method.
func (m *MockAPITokenRepository) UpsertRequest(ctx context.Context, key string, rl entity.RateLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRequest", ctx, key, rl)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRequest indicates an expected call of UpsertRequest.
func (mr *MockAPITokenRepositoryMockRecorder) UpsertRequest(ctx, key, rl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRequest", reflect.TypeOf((*MockAPITokenRepository)(nil).UpsertRequest), ctx, key, rl)
}

// MockIPRepository is a mock of IPRepository interface.
type MockIPRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPRepositoryMockRecorder
}

// MockIPRepositoryMockRecorder is the mock recorder for MockIPRepository.
type MockIPRepositoryMockRecorder struct {
	mock *MockIPRepository
}

// NewMockIPRepository creates a new mock instance.
func NewMockIPRepository(ctrl *gomock.Controller) *MockIPRepository {
	mock := &MockIPRepository{ctrl: ctrl}
	mock.recorder = &MockIPRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPRepository) EXPECT() *MockIPRepositoryMockRecorder {
	return m.recorder
}

// GetBlockedDuration mocks base method.
func (m *MockIPRepository) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockedDuration", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockedDuration indicates an expected call of GetBlockedDuration.
func (mr *MockIPRepositoryMockRecorder) GetBlockedDuration(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockedDuration", reflect.TypeOf((*MockIPRepository)(nil).GetBlockedDuration), ctx, key)
}

// SaveBlockedDuration mocks base method.
func (m *MockIPRepository) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBlockedDuration", ctx, key, BlockedDuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBlockedDuration indicates an expected call of SaveBlockedDuration.
func (mr *MockIPRepositoryMockRecorder) SaveBlockedDuration(ctx, key, BlockedDuration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBlockedDuration", reflect.TypeOf((*MockIPRepository)(nil).SaveBlockedDuration), ctx, key, BlockedDuration)
}

// UpsertRequest mocks base method.
func (m *MockIPRepository) UpsertRequest(ctx context.Context, key string, rl entity.RateLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRequest", ctx, key, rl)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRequest indicates an expected call of UpsertRequest.
func (mr *MockIPRepositoryMockRecorder) UpsertRequest(ctx, key, rl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRequest", reflect.TypeOf((*MockIPRepository)(nil).UpsertRequest), ctx, key, rl)
}
