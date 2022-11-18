package mock

import (
	"avito-internship/internal/myapp/service"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
)

type MockServiceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockServiceRepositoryMockRecorder
}

type MockServiceRepositoryMockRecorder struct {
	mock *MockServiceRepository
}

var _ service.Repository = (*MockServiceRepository)(nil)

func NewMockServiceRepository(ctrl *gomock.Controller) *MockServiceRepository {
	mock := &MockServiceRepository{ctrl: ctrl}
	mock.recorder = &MockServiceRepositoryMockRecorder{mock}
	return mock
}

func (m *MockServiceRepository) EXPECT() *MockServiceRepositoryMockRecorder {
	return m.recorder
}

func (m *MockServiceRepository) ServiceExistsByID(ctx context.Context, serviceID uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceExistsByID", ctx, serviceID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceRepositoryMockRecorder) ServiceExistsByID(ctx, serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ServiceExistsByID",
		reflect.TypeOf((*MockServiceRepository)(nil).ServiceExistsByID),
		ctx,
		serviceID,
	)
}

func (m *MockServiceRepository) GetServiceNameByID(ctx context.Context, serviceID uuid.UUID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceNameByID", ctx, serviceID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceRepositoryMockRecorder) GetServiceNameByID(ctx, serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"GetServiceNameByID",
		reflect.TypeOf((*MockServiceRepository)(nil).GetServiceNameByID),
		ctx,
		serviceID,
	)
}
