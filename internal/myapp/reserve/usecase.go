package reserve

import (
	"avito-internship/internal/myapp/models"
	"context"
)

type UseCase interface {
	ReserveBalance(context.Context, models.ReserveInfo) error
}
