package reserve

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	ReserveBalance(context.Context, models.ReserveInfo) error
	AcceptReserve(context.Context, models.ReserveInfo) error
	CancelReserve(context.Context, models.ReserveInfo) error
	CreateEmptyReserve(context.Context, uuid.UUID) error
}
