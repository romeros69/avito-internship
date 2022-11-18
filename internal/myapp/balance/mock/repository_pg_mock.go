package mock

import (
	"avito-internship/internal/myapp/balance"
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
)

type MockBalanceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBalanceRepositoryMockRecorder
}

type MockBalanceRepositoryMockRecorder struct {
	mock *MockBalanceRepository
}

var _ balance.Repository = (*MockBalanceRepository)(nil)

func NewMockBalanceRepository(ctrl *gomock.Controller) *MockBalanceRepository {
	mock := &MockBalanceRepository{ctrl: ctrl}
	mock.recorder = &MockBalanceRepositoryMockRecorder{mock}
	return mock
}

func (m *MockBalanceRepository) EXPECT() *MockBalanceRepositoryMockRecorder {
	return m.recorder
}

func (m *MockBalanceRepository) GetBalanceByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceByUserID", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBalanceRepositoryMockRecorder) GetBalanceByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"GetBalanceByUserID",
		reflect.TypeOf((*MockBalanceRepository)(nil).GetBalanceByUserID),
		ctx,
		userID,
	)
}

func (m *MockBalanceRepository) ReplenishmentBalance(ctx context.Context, replenishment models.Replenishment) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplenishmentBalance", ctx, replenishment)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBalanceRepositoryMockRecorder) ReplenishmentBalance(ctx, replenishment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ReplenishmentBalance",
		reflect.TypeOf((*MockBalanceRepository)(nil).ReplenishmentBalance),
		ctx,
		replenishment,
	)
}

func (m *MockBalanceRepository) GetBalanceIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceIDByUserID", ctx, userID)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBalanceRepositoryMockRecorder) GetBalanceIDByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"GetBalanceIDByUserID",
		reflect.TypeOf((*MockBalanceRepository)(nil).GetBalanceIDByUserID),
		ctx,
		userID,
	)
}

func (m *MockBalanceRepository) BalanceExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceExistsByUserID", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBalanceRepositoryMockRecorder) BalanceExistsByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"BalanceExistsByUserID",
		reflect.TypeOf((*MockBalanceRepository)(nil).BalanceExistsByUserID),
		ctx,
		userID,
	)
}

func (m *MockBalanceRepository) CreateEmptyBalance(ctx context.Context, balance models.Balance) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmptyBalance", ctx, balance)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBalanceRepositoryMockRecorder) CreateEmptyBalance(ctx, balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"CreateEmptyBalance",
		reflect.TypeOf((*MockBalanceRepository)(nil).CreateEmptyBalance),
		ctx,
		balance,
	)
}

func (m *MockBalanceRepository) TransferBalance(ctx context.Context, balanceID uuid.UUID, value int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferBalance", ctx, balanceID, value)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockBalanceRepositoryMockRecorder) TransferBalance(ctx, balanceID, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"TransferBalance",
		reflect.TypeOf((*MockBalanceRepository)(nil).TransferBalance),
		ctx,
		balanceID,
		value,
	)
}

func (m *MockBalanceRepository) ReturnMoneyFromReserve(ctx context.Context, balanceID uuid.UUID, value int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnMoneyFromReserve", ctx, balanceID, value)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockBalanceRepositoryMockRecorder) ReturnMoneyFromReserve(ctx, balanceID, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ReturnMoneyFromReserve",
		reflect.TypeOf((*MockBalanceRepository)(nil).ReturnMoneyFromReserve),
		ctx,
		balanceID,
		value,
	)
}
