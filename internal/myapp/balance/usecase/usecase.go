package usecase

import (
	"avito-internship/internal/myapp/balance"
	historyUseCase "avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
	"time"
)

type BalanceUseCase struct {
	repo      balance.Repository
	historyUC historyUseCase.UseCase
}

func NewBalanceUseCase(repo balance.Repository, historyUC historyUseCase.UseCase) *BalanceUseCase {
	return &BalanceUseCase{repo: repo, historyUC: historyUC}
}

var _ balance.UseCase = (*BalanceUseCase)(nil)

func (b *BalanceUseCase) CreateEmptyBalance(ctx context.Context, balance models.Balance) error {
	return b.repo.CreateEmptyBalance(ctx, balance)
}

func (b *BalanceUseCase) BalanceExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	return b.repo.BalanceExistsByUserID(ctx, userID)
}

func (b *BalanceUseCase) GetBalanceByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	return b.repo.GetBalanceByUserID(ctx, userID)
}

func (b *BalanceUseCase) GetBalanceIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	return b.repo.GetBalanceIDByUserID(ctx, userID)
}

func (b *BalanceUseCase) TransferBalance(ctx context.Context, balanceID uuid.UUID, value int64) error {
	return b.repo.TransferBalance(ctx, balanceID, value)
}

func (b *BalanceUseCase) ReplenishmentBalance(ctx context.Context, replenishment models.Replenishment) error {
	isExists, err := b.BalanceExistsByUserID(ctx, replenishment.UserID)
	if err != nil {
		return err
	}
	if !isExists {
		err = b.CreateEmptyBalance(ctx, models.Balance{
			ID:     uuid.New(),
			UserID: replenishment.UserID,
		})
		if err != nil {
			return err
		}
	}
	balanceID, err := b.repo.ReplenishmentBalance(ctx, replenishment)
	if err != nil {
		return err
	}
	err = b.historyUC.CreateHistory(ctx, models.History{
		ID:                  uuid.New(),
		BalanceID:           balanceID,
		TypeHistory:         "replenishment",
		SourceReplenishment: replenishment.Source,
		Date:                time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
