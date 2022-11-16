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

func (h *HistoryRepo) GetCountHistoryByBalanceID(ctx context.Context, balanceID uuid.UUID) (int64, error) {
	query := `select count(*) from history where balance_id=$1`

	rows, err := h.pg.Pool.Query(ctx, query, balanceID)
	if err != nil {
		return -1, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var count int64
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, fmt.Errorf("error parsing count of story by user: %w", err)
		}
	}
	return count, nil
}

func (h *HistoryRepo) GetHistoryByBalanceID(ctx context.Context, pagination models.Pagination, balanceID uuid.UUID) ([]models.History, error) {
	query := ``
	if pagination.OrderBy == "date" {
		query = `select * from history where balance_id = $1 order by date limit $2 offset $3`
	} else {
		query = `select * from history where balance_id = $1 order by value limit $2 offset $3`
	}

	offset := pagination.GetOffset()
	rows, err := h.pg.Pool.Query(ctx, query, balanceID, pagination.Size, offset)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var historyList []models.History

	var source interface{}
	for rows.Next() {
		var historyEntity models.History

		err = rows.Scan(
			&historyEntity.ID,
			&historyEntity.BalanceID,
			&historyEntity.TypeHistory,
			&historyEntity.OrderID,
			&historyEntity.ServiceID,
			&historyEntity.Value,
			&source, // вот тут
			&historyEntity.Date,
		)
		if err != nil {
			return nil, fmt.Errorf("error in parsing history: %w", err)
		}
		if source != nil {
			historyEntity.SourceReplenishment = source.(string)
		} else {
			historyEntity.SourceReplenishment = ""
		}
		historyList = append(historyList, historyEntity)
	}

	return historyList, nil
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
	query := `insert into history (id, balance_id, type_history, order_id, service_id, value, source_replenishment, date) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	rows, err := h.pg.Pool.Query(ctx,
		query,
		history.ID,
		history.BalanceID,
		history.TypeHistory,
		checkID(history.OrderID),
		checkID(history.ServiceID),
		history.Value,
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
