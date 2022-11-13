package reserve

import (
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	ReserveBalance(context.Context, models.Reserve) (uuid.UUID, error)
}
