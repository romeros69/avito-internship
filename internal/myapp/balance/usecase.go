package balance

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	GetBalanceByUserID(context.Context, uuid.UUID) (int64, error)
	GetBalanceIDByUserID(context.Context, uuid.UUID) (uuid.UUID, error)
	ReplenishmentBalance(context.Context, models.Replenishment) error
	BalanceExistsByUserID(context.Context, uuid.UUID) (bool, error)
	CreateEmptyBalance(context.Context, models.Balance) (uuid.UUID, error)
	TransferBalance(context.Context, uuid.UUID, int64) error
	CheckBeforeReserve(context.Context, uuid.UUID, int64) (bool, error)
	ReturnMoneyFromReserve(context.Context, uuid.UUID, int64) error
}
