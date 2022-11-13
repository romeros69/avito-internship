package service

import (
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	ServiceExistsByID(ctx context.Context, uuid uuid.UUID) (bool, error)
}
