package usecase

import (
	"avito-internship/internal/myapp/balance/usecase/repo"
	"context"
	"github.com/google/uuid"
)

type BalanceContract interface {
	GetBalanceByUserID(context.Context, uuid.UUID) (int64, error)
}

type BalanceUseCase struct {
	repo repo.BalanceRp
}

func NewBalanceUseCase(repo repo.BalanceRp) *BalanceUseCase {
	return &BalanceUseCase{
		repo: repo,
	}
}

var _ BalanceContract = (*BalanceUseCase)(nil)

func (b *BalanceUseCase) GetBalanceByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	return b.repo.GetBalanceByUserID(ctx, userID)
}
