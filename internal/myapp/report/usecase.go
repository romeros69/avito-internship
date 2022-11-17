package report

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	CreateReport(ctx context.Context, report models.Report) (uuid.UUID, error)
	GetReport(context.Context, int, int) error
}
