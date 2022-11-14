package usecase

import (
	"avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/models"
	"context"
	"github.com/google/uuid"
)

type HistoryUseCase struct {
	repo history.Repository
}

func NewHistoryUseCase(repo history.Repository) *HistoryUseCase {
	return &HistoryUseCase{
		repo: repo,
	}
}

var _ history.UseCase = (*HistoryUseCase)(nil)

func (h *HistoryUseCase) CreateHistory(ctx context.Context, history models.History) error {
	return h.repo.CreateHistory(ctx, history)
}

func (h *HistoryUseCase) CheckHistoryForReserve(ctx context.Context, reserveInfo models.ReserveInfo, balanceID uuid.UUID) (bool, error) {
	countOpen, err := h.repo.GetCountHistoryForReserveByType(ctx, models.HistoryInfo{
		BalanceID:   balanceID,
		OrderID:     reserveInfo.OrderID,
		ServiceID:   reserveInfo.ServiceID,
		TypeHistory: "reserve",
	})
	if err != nil {
		return false, err
	}
	countClose, err := h.repo.GetCountHistoryForReserveByType(ctx, models.HistoryInfo{
		BalanceID:   balanceID,
		OrderID:     reserveInfo.OrderID,
		ServiceID:   reserveInfo.ServiceID,
		TypeHistory: "confirmation",
	})
	if err != nil {
		return false, err
	}
	if countOpen > countClose {
		return true, nil
	}
	return false, nil
}
