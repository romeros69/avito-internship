package usecase

import (
	"avito-internship/internal/myapp/service"
	"context"
	"github.com/google/uuid"
)

type ServiceUseCase struct {
	repo service.Repository
}

func NewServiceUseCase(repo service.Repository) *ServiceUseCase {
	return &ServiceUseCase{repo: repo}
}

var _ service.UseCase = (*ServiceUseCase)(nil)

func (s *ServiceUseCase) ServiceExistsByID(ctx context.Context, serviceID uuid.UUID) (bool, error) {
	return s.repo.ServiceExistsByID(ctx, serviceID)
}
