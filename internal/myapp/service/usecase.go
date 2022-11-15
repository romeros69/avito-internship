package service

import (
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	ServiceExistsByID(context.Context, uuid.UUID) (bool, error)
	GetServiceNameByID(context.Context, uuid.UUID) (string, error)
}
