package repository

import (
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/reserve"
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type ReserveRepo struct {
	pg *postgres.Postgres
}

func NewReserveRepo(pg *postgres.Postgres) *ReserveRepo {
	return &ReserveRepo{
		pg: pg,
	}
}

var _ reserve.Repository = (*ReserveRepo)(nil)

func (r *ReserveRepo) ReserveBalance(ctx context.Context, reserve models.Reserve) (uuid.UUID, error) {
	query := `update reserve set value = value + $1 where balance_id = $2 returning id`

	rows, err := r.pg.Pool.Query(
		ctx,
		query,
		reserve.Value,
		reserve.BalanceID,
	)
	if err != nil {
		log.Printf("cannot execute query: %w", err)
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		log.Printf("error in add money to reserve for balance (balance_id=%s)", reserve.BalanceID.String())
		return uuid.Nil, fmt.Errorf("error in add money to reserve for balance (balance_id=%s)", reserve.BalanceID.String())
	}
	var reserveID uuid.UUID
	err = rows.Scan(&reserveID)
	if err != nil {
		log.Printf("error in parsing id reserve of reserve reserve balance: %w", err)
		return uuid.Nil, fmt.Errorf("error in parsing id reserve of reserve reserve balance: %w", err)
	}
	return uuid.Nil, nil
}

func (r *ReserveRepo) CreateEmptyReserve(ctx context.Context, reserve models.Reserve) error {
	query := `insert into reserve (id, balance_id, value) values ($1, $2, $3) returning id`

	rows, err := r.pg.Pool.Query(ctx, query, reserve.ID, reserve.BalanceID, reserve.Value)
	if err != nil {
		log.Printf("cannot execute query: %w", err)
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		log.Printf("error in creating empty reserve for balance (balance_id=%s)", reserve.BalanceID.String())
		return fmt.Errorf("error in creating empty reserve for balance (balance_id=%s)", reserve.BalanceID.String())
	}
	return nil
}

func (r *ReserveRepo) SubtractionReserve(ctx context.Context, balanceID uuid.UUID, value int64) error {
	query := `update reserve set value = value - $1 where balance_id = $2 returning id`

	rows, err := r.pg.Pool.Query(ctx, query, value, balanceID)
	if err != nil {
		log.Printf("cannot execute query: %w", err)
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		log.Printf("not enough money in reserve")
		return fmt.Errorf("not enough money in reserve")
	}
	return nil
}
