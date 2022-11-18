package mock

import (
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/report"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type MockReportRepository struct {
	ctrl     *gomock.Controller
	recorder *MockReportRepositoryMockRecorder
}

type MockReportRepositoryMockRecorder struct {
	mock *MockReportRepository
}

var _ report.Repository = (*MockReportRepository)(nil)

func NewMockReportRepository(ctrl *gomock.Controller) *MockReportRepository {
	mock := &MockReportRepository{ctrl: ctrl}
	mock.recorder = &MockReportRepositoryMockRecorder{mock}
	return mock
}

func (m *MockReportRepository) EXPECT() *MockReportRepositoryMockRecorder {
	return m.recorder
}

func (m *MockReportRepository) CreateReport(ctx context.Context, report models.Report) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReport", ctx, report)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReportRepositoryMockRecorder) CreateReport(ctx, report interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"CreateReport",
		reflect.TypeOf((*MockReportRepository)(nil).CreateReport),
		ctx,
		report,
	)
}

func (m *MockReportRepository) GetReport(ctx context.Context, start time.Time) ([]models.ReportResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReport", ctx, start)
	ret0, _ := ret[0].([]models.ReportResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReportRepositoryMockRecorder) GetReport(ctx, start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"GetReport",
		reflect.TypeOf((*MockReportRepository)(nil).GetReport),
		ctx,
		start,
	)
}
