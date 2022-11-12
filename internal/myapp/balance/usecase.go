package balance

import (
	historyEntity "avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

// Balance use case
type UseCase interface {
	GetBalanceByUserID(context.Context, uuid.UUID) (int64, error)
	ReplenishmentBalance(context.Context, historyEntity.Replenishment) error
}
