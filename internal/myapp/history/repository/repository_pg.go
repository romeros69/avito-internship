package repository

import (
	"avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type HistoryRepo struct {
	pg *postgres.Postgres
}

func NewHistoryRepo(pg *postgres.Postgres) *HistoryRepo {
	return &HistoryRepo{
		pg: pg,
	}
}

var _ history.Repository = (*HistoryRepo)(nil)

func checkID(id uuid.UUID) any {
	if id == uuid.Nil {
		return nil
	}
	return id
}

func checkSourceReplenishment(source string) any {
	if source == "" {
		return nil
	}
	return source
}

func (h *HistoryRepo) GetCountHistoryForReserveByType(ctx context.Context, historyInfo models.HistoryInfo) (uint, error) {
	query := `select count(*) from history where balance_id = $1 and order_id = $2 and service_id = $3 and type_history = $4`

	rows, err := h.pg.Pool.Query(ctx, query, historyInfo.BalanceID, historyInfo.OrderID, historyInfo.ServiceID, historyInfo.TypeHistory)
	if err != nil {
		return 0, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var count uint
	if !rows.Next() {
		return 0, fmt.Errorf("this order no longer exists")
	}
	err = rows.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error in parsing count history for reserve by type history: %w", err)
	}
	return count, nil
}

func (h *HistoryRepo) CreateHistory(ctx context.Context, history models.History) error {
	query := `insert into history (id, balance_id, type_history, order_id, service_id, source_replenishment, date) values ($1, $2, $3, $4, $5, $6, $7) returning id`

	rows, err := h.pg.Pool.Query(ctx,
		query,
		history.ID,
		history.BalanceID,
		history.TypeHistory,
		checkID(history.OrderID),
		checkID(history.ServiceID),
		checkSourceReplenishment(history.SourceReplenishment),
		history.Date)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("error on creating history")
	}
	return nil
}
