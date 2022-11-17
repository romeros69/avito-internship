package report

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	CreateReport(ctx context.Context, report models.Report) (uuid.UUID, error)
	GetReport(context.Context, time.Time) ([]models.ReportResult, error)
}
