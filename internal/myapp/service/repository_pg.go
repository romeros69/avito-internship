package service

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	ServiceExistsByID(ctx context.Context, uuid uuid.UUID) (bool, error)
}
