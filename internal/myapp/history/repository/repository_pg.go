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

func (h *HistoryRepo) CreateHistory(ctx context.Context, history models.History) error {
	query := `insert into history (id, balance_id, type_history, reserve_id, report_id, source_replenishment, date) values ($1, $2, $3, $4, $5, $6, $7) returning id`

	rows, err := h.pg.Pool.Query(ctx,
		query,
		history.ID,
		history.BalanceID,
		history.TypeHistory,
		checkID(history.ReserveID),
		checkID(history.ReportID),
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

var _ history.Repository = (*HistoryRepo)(nil)
