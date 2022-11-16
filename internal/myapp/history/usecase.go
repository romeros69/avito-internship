package history

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	CreateHistory(context.Context, models.History) error
	CheckHistoryForReserve(context.Context, models.ReserveInfo, uuid.UUID) (bool, error)
	GetHistoryByUserID(context.Context, models.Pagination, uuid.UUID) (models.HistoryTransfer, error)
}
