package balance

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	GetBalanceByUserID(context.Context, uuid.UUID) (int64, error)
	ReplenishmentBalance(context.Context, models.Replenishment) (uuid.UUID, error)
	GetBalanceIDByUserID(context.Context, uuid.UUID) (uuid.UUID, error)
	BalanceExistsByUserID(context.Context, uuid.UUID) (bool, error)
	CreateEmptyBalance(context.Context, models.Balance) (uuid.UUID, error)
	TransferBalance(context.Context, uuid.UUID, int64) error
	ReturnMoneyFromReserve(context.Context, uuid.UUID, int64) error
}
