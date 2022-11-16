package history

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	CreateHistory(context.Context, models.History) error
	GetCountHistoryForReserveByType(context.Context, models.HistoryInfo) (uint, error)
	GetHistoryByBalanceID(context.Context, models.Pagination, uuid.UUID) ([]models.History, error)
	GetCountHistoryByBalanceID(context.Context, uuid.UUID) (int64, error)
}
