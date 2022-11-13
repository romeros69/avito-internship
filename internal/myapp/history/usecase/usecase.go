package usecase

import (
	"avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/models"
	"context"
)

type HistoryUseCase struct {
	repo history.Repository
}

func NewHistoryUseCase(repo history.Repository) *HistoryUseCase {
	return &HistoryUseCase{
		repo: repo,
	}
}

func (h *HistoryUseCase) CreateHistory(ctx context.Context, history models.History) error {
	return h.repo.CreateHistory(ctx, history)
}

var _ history.UseCase = (*HistoryUseCase)(nil)
