package usecase

import (
	"avito-internship/internal/myapp/balance"
	"avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/report"
	"avito-internship/internal/myapp/reserve"
	"avito-internship/internal/myapp/service"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ReserveUseCase struct {
	repo      reserve.Repository
	balanceUC balance.UseCase
	historyUC history.UseCase
	reportUC  report.UseCase
	serviceUC service.UseCase
}

func NewReserveUseCase(repo reserve.Repository, balanceUC balance.UseCase, historyUC history.UseCase, reportUC report.UseCase, serviceUC service.UseCase) *ReserveUseCase {
	return &ReserveUseCase{repo: repo, balanceUC: balanceUC, historyUC: historyUC, reportUC: reportUC, serviceUC: serviceUC}
}

var _ reserve.UseCase = (*ReserveUseCase)(nil)

func (r *ReserveUseCase) ReserveBalance(ctx context.Context, reserveInfo models.ReserveInfo) error {

	// проверка того, что сервис с таким id есть
	serviceExists, err := r.serviceUC.ServiceExistsByID(ctx, reserveInfo.ServiceID)
	switch {
	case err != nil:
		return err
	case !serviceExists:
		return fmt.Errorf("service with this id (service_id = %s) does not exist", reserveInfo.ServiceID.String())
	}

	// получение id баланса по user_id
	balanceID, err := r.balanceUC.GetBalanceIDByUserID(ctx, reserveInfo.UserID)
	if err != nil {
		return err
	}
	// создание резерва
	reserveID, err := r.repo.ReserveBalance(ctx, models.Reserve{
		ID:        uuid.New(),
		BalanceID: balanceID,
		Value:     reserveInfo.Value,
		Status:    "active",
	})
	if err != nil {
		return err
	}

	// снятие денег с основного счета
	err = r.balanceUC.TransferBalance(ctx, balanceID, reserveInfo.Value)
	if err != nil {
		return err
	}

	// создание отчета
	reportID, err := r.reportUC.CreateReport(
		ctx,
		models.Report{
			ID:        uuid.New(),
			ServiceID: reserveInfo.ServiceID,
			OrderID:   reserveInfo.OrderID,
			Value:     reserveInfo.Value,
			Date:      time.Now(),
		})
	if err != nil {
		return err
	}
	// создание истории
	err = r.historyUC.CreateHistory(ctx, models.History{
		ID:          uuid.New(),
		BalanceID:   balanceID,
		TypeHistory: "reserve",
		ReserveID:   reserveID,
		ReportID:    reportID,
		Date:        time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
