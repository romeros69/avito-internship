package usecase

import (
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/report"
	"context"
	"github.com/google/uuid"
)

type ReportUseCase struct {
	repo report.Repository
}

func NewReportUseCase(repo report.Repository) *ReportUseCase {
	return &ReportUseCase{repo: repo}
}

var _ report.UseCase = (*ReportUseCase)(nil)

func (r *ReportUseCase) CreateReport(ctx context.Context, report models.Report) (uuid.UUID, error) {
	return r.repo.CreateReport(ctx, report)
}
