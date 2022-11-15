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

func NewSimpleReserveUseCase(repo reserve.Repository, historyUC history.UseCase, reportUC report.UseCase, serviceUC service.UseCase) *ReserveUseCase {
	return &ReserveUseCase{
		repo:      repo,
		historyUC: historyUC,
		reportUC:  reportUC,
		serviceUC: serviceUC,
	}
}

func NewReserveUseCase(repo reserve.Repository, balanceUC balance.UseCase, historyUC history.UseCase, reportUC report.UseCase, serviceUC service.UseCase) *ReserveUseCase {
	return &ReserveUseCase{
		repo:      repo,
		balanceUC: balanceUC,
		historyUC: historyUC,
		reportUC:  reportUC,
		serviceUC: serviceUC,
	}
}

var _ reserve.UseCase = (*ReserveUseCase)(nil)

func (r *ReserveUseCase) CreateEmptyReserve(ctx context.Context, balanceID uuid.UUID) error {
	return r.repo.CreateEmptyReserve(ctx, models.Reserve{
		ID:        uuid.New(),
		BalanceID: balanceID,
		Value:     0,
	})
}

func (r *ReserveUseCase) ReserveBalance(ctx context.Context, reserveInfo models.ReserveInfo) error {
	serviceExists, err := r.serviceUC.ServiceExistsByID(ctx, reserveInfo.ServiceID)
	switch {
	case err != nil:
		return err
	case !serviceExists:
		return fmt.Errorf("service with this id (service_id = %s) does not exist", reserveInfo.ServiceID.String())
	}

	balanceID, err := r.balanceUC.GetBalanceIDByUserID(ctx, reserveInfo.UserID)
	if err != nil {
		return err
	}
	ok, err := r.balanceUC.CheckBeforeReserve(ctx, reserveInfo.UserID, reserveInfo.Value)
	switch {
	case err != nil:
		return err
	case !ok:
		return fmt.Errorf("not enough money on balance")
	}

	_, err = r.repo.ReserveBalance(ctx, models.Reserve{
		BalanceID: balanceID,
		Value:     reserveInfo.Value,
	})
	if err != nil {
		return err
	}

	err = r.balanceUC.TransferBalance(ctx, balanceID, reserveInfo.Value)
	if err != nil {
		return err
	}

	err = r.historyUC.CreateHistory(ctx, models.History{
		ID:          uuid.New(),
		BalanceID:   balanceID,
		TypeHistory: "reserve",
		OrderID:     reserveInfo.OrderID,
		ServiceID:   reserveInfo.ServiceID,
		Value:       reserveInfo.Value,
		Date:        time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *ReserveUseCase) AcceptReserve(ctx context.Context, reserveInfo models.ReserveInfo) error {
	serviceExists, err := r.serviceUC.ServiceExistsByID(ctx, reserveInfo.ServiceID)
	switch {
	case err != nil:
		return err
	case !serviceExists:
		return fmt.Errorf("service with this id (service_id = %s) does not exist", reserveInfo.ServiceID.String())
	}

	balanceID, err := r.balanceUC.GetBalanceIDByUserID(ctx, reserveInfo.UserID)
	if err != nil {
		return err
	}

	canConfirm, err := r.historyUC.CheckHistoryForReserve(ctx, reserveInfo, balanceID)
	switch {
	case err != nil:
		return err
	case !canConfirm:
		return fmt.Errorf("it is impossible to confirm the service due to the lack of a reserve")
	}

	err = r.repo.SubtractionReserve(ctx, balanceID, reserveInfo.Value)
	if err != nil {
		return err
	}

	err = r.historyUC.CreateHistory(ctx, models.History{
		ID:          uuid.New(),
		BalanceID:   balanceID,
		TypeHistory: "confirmation",
		OrderID:     reserveInfo.OrderID,
		ServiceID:   reserveInfo.ServiceID,
		Value:       reserveInfo.Value,
		Date:        time.Now(),
	})
	if err != nil {
		return err
	}

	_, err = r.reportUC.CreateReport(
		ctx,
		models.Report{
			ID:        uuid.New(),
			ServiceID: reserveInfo.ServiceID,
			Value:     reserveInfo.Value,
			Date:      time.Now(),
		})
	if err != nil {
		return err
	}
	return nil
}
