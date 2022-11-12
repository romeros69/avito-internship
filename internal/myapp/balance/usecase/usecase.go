package usecase

import (
	"avito-internship/internal/myapp/balance"
	historyEntity "avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type BalanceUseCase struct {
	repo balance.Repository
}

func NewBalanceUseCase(repo balance.Repository) *BalanceUseCase {
	return &BalanceUseCase{repo: repo}
}

var _ balance.UseCase = (*BalanceUseCase)(nil)

func (b *BalanceUseCase) GetBalanceByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	return b.repo.GetBalanceByUserID(ctx, userID)
}

func (b *BalanceUseCase) ReplenishmentBalance(ctx context.Context, replenishment historyEntity.Replenishment) error {
	return nil
}
