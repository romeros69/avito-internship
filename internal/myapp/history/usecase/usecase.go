package usecase

import (
	"avito-internship/internal/myapp/balance"
	"avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/service"
	"context"
	"github.com/google/uuid"
)

type HistoryUseCase struct {
	repo      history.Repository
	balanceUC balance.UseCase
	serviceUC service.UseCase
}

func NewSimpleHistoryUseCase(repo history.Repository) *HistoryUseCase {
	return &HistoryUseCase{repo: repo}
}

func NewHistoryUseCase(repo history.Repository, balanceUC balance.UseCase, serviceUC service.UseCase) *HistoryUseCase {
	return &HistoryUseCase{repo: repo, balanceUC: balanceUC, serviceUC: serviceUC}
}

var _ history.UseCase = (*HistoryUseCase)(nil)

func (h *HistoryUseCase) GetCountHistoryByBalanceID(ctx context.Context, balanceID uuid.UUID) (int64, error) {
	return h.repo.GetCountHistoryByBalanceID(ctx, balanceID)
}

func (h *HistoryUseCase) GetHistoryByUserID(ctx context.Context, pagination models.Pagination, userID uuid.UUID) (models.HistoryTransfer, error) {
	balanceID, err := h.balanceUC.GetBalanceIDByUserID(ctx, userID)
	if err != nil {
		return models.HistoryTransfer{}, err
	}
	historyList, err := h.repo.GetHistoryByBalanceID(ctx, pagination, balanceID)
	if err != nil {
		return models.HistoryTransfer{}, err
	}
	var serviceNameList []string
	var serviceName string
	for _, v := range historyList {
		if v.TypeHistory != "replenishment" {
			serviceName, err = h.serviceUC.GetServiceNameByID(ctx, v.ServiceID)
			if err != nil {
				return models.HistoryTransfer{}, err
			}
		} else {
			serviceName = ""
		}
		serviceNameList = append(serviceNameList, serviceName)
	}
	count, err := h.GetCountHistoryByBalanceID(ctx, balanceID)
	if err != nil {
		return models.HistoryTransfer{}, err
	}
	return models.HistoryTransfer{
		Histories:    historyList,
		ServiceNames: serviceNameList,
		Count:        count,
	}, nil
}

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
	countCancel, err := h.repo.GetCountHistoryForReserveByType(ctx, models.HistoryInfo{
		BalanceID:   balanceID,
		OrderID:     reserveInfo.OrderID,
		ServiceID:   reserveInfo.ServiceID,
		TypeHistory: "cancel_reserve",
	})
	if err != nil {
		return false, err
	}
	if countOpen > (countClose + countCancel) {
		return true, nil
	}
	return false, nil
}
