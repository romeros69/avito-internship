package repository

import (
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/reserve"
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type ReserveRepo struct {
	pg *postgres.Postgres
}

func NewReserveRepo(pg *postgres.Postgres) *ReserveRepo {
	return &ReserveRepo{pg: pg}
}

var _ reserve.Repository = (*ReserveRepo)(nil)

func (r *ReserveRepo) ReserveBalance(ctx context.Context, reserve models.Reserve) (uuid.UUID, error) {
	query := `insert into reserve (id, balance_id, value, status) values ($1, $2, $3, $4) returning id`

	rows, err := r.pg.Pool.Query(
		ctx,
		query,
		reserve.ID,
		reserve.BalanceID,
		reserve.Value,
		reserve.Status,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return uuid.Nil, fmt.Errorf("error in creating reserve for balance (balance_id=%s)", reserve.BalanceID.String())
	}
	var reserveID uuid.UUID
	err = rows.Scan(&reserveID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error in parsing id reserve of reserve reserve balance: %w", err)
	}
	return reserveID, nil
}
