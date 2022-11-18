package mock

import (
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/reserve"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
)

type MockReserveRepository struct {
	ctrl     *gomock.Controller
	recorder *MockReserveRepositoryMockRecorder
}

type MockReserveRepositoryMockRecorder struct {
	mock *MockReserveRepository
}

var _ reserve.Repository = (*MockReserveRepository)(nil)

func NewMockReserveRepository(ctrl *gomock.Controller) *MockReserveRepository {
	mock := &MockReserveRepository{ctrl: ctrl}
	mock.recorder = &MockReserveRepositoryMockRecorder{mock}
	return mock
}

func (m *MockReserveRepository) EXPECT() *MockReserveRepositoryMockRecorder {
	return m.recorder
}

func (m *MockReserveRepository) ReserveBalance(ctx context.Context, reserve models.Reserve) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveBalance", ctx, reserve)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReserveRepositoryMockRecorder) ReserveBalance(ctx, reserve interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ReserveBalance",
		reflect.TypeOf((*MockReserveRepository)(nil).ReserveBalance),
		ctx,
		reserve,
	)
}

func (m *MockReserveRepository) SubtractionReserve(ctx context.Context, balanceID uuid.UUID, value int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubtractionReserve", ctx, balanceID, value)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockReserveRepositoryMockRecorder) SubtractionReserve(ctx, balanceID, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"SubtractionReserve",
		reflect.TypeOf((*MockReserveRepository)(nil).SubtractionReserve),
		ctx,
		balanceID,
		value,
	)
}

func (m *MockReserveRepository) CreateEmptyReserve(ctx context.Context, reserve models.Reserve) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmptyReserve", ctx, reserve)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockReserveRepositoryMockRecorder) CreateEmptyReserve(ctx, reserve interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"CreateEmptyReserve",
		reflect.TypeOf((*MockReserveRepository)(nil).CreateEmptyReserve),
		ctx,
		reserve,
	)
}
