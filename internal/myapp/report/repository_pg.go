package report

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	CreateReport(ctx context.Context, report models.Report) (uuid.UUID, error)
}
