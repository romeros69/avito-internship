package mock

import (
	"avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
)

type MockHistoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHistoryRepositoryMockRecorder
}

type MockHistoryRepositoryMockRecorder struct {
	mock *MockHistoryRepository
}

var _ history.Repository = (*MockHistoryRepository)(nil)

func NewMockHistoryRepository(ctrl *gomock.Controller) *MockHistoryRepository {
	mock := &MockHistoryRepository{ctrl: ctrl}
	mock.recorder = &MockHistoryRepositoryMockRecorder{mock}
	return mock
}

func (m *MockHistoryRepository) EXPECT() *MockHistoryRepositoryMockRecorder {
	return m.recorder
}

func (m *MockHistoryRepository) CreateHistory(ctx context.Context, history models.History) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHistory", ctx, history)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockHistoryRepositoryMockRecorder) CreateHistory(ctx, history interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"CreateHistory",
		reflect.TypeOf((*MockHistoryRepository)(nil).CreateHistory),
		ctx,
		history,
	)
}

func (m *MockHistoryRepository) GetCountHistoryForReserveByType(ctx context.Context, info models.HistoryInfo) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountHistoryForReserveByType", ctx, info)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHistoryRepositoryMockRecorder) GetCountHistoryForReserveByType(ctx, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"GetCountHistoryForReserveByType",
		reflect.TypeOf((*MockHistoryRepository)(nil).GetCountHistoryForReserveByType),
		ctx,
		info,
	)

}

func (m *MockHistoryRepository) GetHistoryByBalanceID(ctx context.Context, pagination models.Pagination, balanceID uuid.UUID) ([]models.History, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryByBalanceID", ctx, pagination, balanceID)
	ret0, _ := ret[0].([]models.History)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHistoryRepositoryMockRecorder) GetHistoryByBalanceID(ctx, pagination, balanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"GetHistoryByBalanceID",
		reflect.TypeOf((*MockHistoryRepository)(nil).GetHistoryByBalanceID),
		ctx,
		pagination,
		balanceID,
	)
}

func (m *MockHistoryRepository) GetCountHistoryByBalanceID(ctx context.Context, balanceID uuid.UUID) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountHistoryByBalanceID", ctx, balanceID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHistoryRepositoryMockRecorder) GetCountHistoryByBalanceID(ctx, balanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"GetCountHistoryByBalanceID",
		reflect.TypeOf((*MockHistoryRepository)(nil).GetCountHistoryByBalanceID),
		ctx,
		balanceID,
	)
}
