package usecase

import (
	"avito-internship/internal/myapp/balance"
	historyUseCase "avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/reserve"
	"context"
	"github.com/google/uuid"
	"time"
)

type BalanceUseCase struct {
	repo      balance.Repository
	historyUC historyUseCase.UseCase
	reserveUC reserve.UseCase
}

func NewBalanceUseCase(repo balance.Repository, historyUC historyUseCase.UseCase, reserveUC reserve.UseCase) *BalanceUseCase {
	return &BalanceUseCase{repo: repo, historyUC: historyUC, reserveUC: reserveUC}
}

var _ balance.UseCase = (*BalanceUseCase)(nil)

func (b *BalanceUseCase) CheckBeforeReserve(ctx context.Context, userID uuid.UUID, value int64) (bool, error) {
	currentValue, err := b.GetBalanceByUserID(ctx, userID)
	switch {
	case err != nil:
		return false, err
	case currentValue >= value:
		return true, nil
	default:
		return false, nil
	}
}

// use ok
func (b *BalanceUseCase) CreateEmptyBalance(ctx context.Context, balance models.Balance) (uuid.UUID, error) {
	return b.repo.CreateEmptyBalance(ctx, balance)
}

// use ok
func (b *BalanceUseCase) BalanceExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	return b.repo.BalanceExistsByUserID(ctx, userID)
}

// use ok
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
	// проверка на существование балана
	isExists, err := b.BalanceExistsByUserID(ctx, replenishment.UserID)
	if err != nil {
		return err
	}
	// создание баланса и резерва если это надо
	if !isExists {
		balanceID, err := b.CreateEmptyBalance(ctx, models.Balance{
			ID:     uuid.New(),
			UserID: replenishment.UserID,
		})
		if err != nil {
			return err
		}
		err = b.reserveUC.CreateEmptyReserve(ctx, balanceID)
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
		Value:               replenishment.Value,
		SourceReplenishment: replenishment.Source,
		Date:                time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}
