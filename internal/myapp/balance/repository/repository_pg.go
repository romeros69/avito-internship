package repository

import (
	"avito-internship/internal/myapp/balance"
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
)

// add check is ecists by id
// add create new balance when replenship

type BalanceRepo struct {
	pg *postgres.Postgres
}

func NewBalanceRepo(pg *postgres.Postgres) *BalanceRepo {
	return &BalanceRepo{
		pg: pg,
	}
}

var _ balance.Repository = (*BalanceRepo)(nil)

func (b *BalanceRepo) GetBalanceByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	query := `select value from balance where user_id = $1`

	rows, err := b.pg.Pool.Query(ctx, query, userID)
	if err != nil {
		return -1, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var value int64
	if rows.Next() {
		err = rows.Scan(&value)
		if err != nil {
			return -1, fmt.Errorf("error in parsing value balance by user id: %w", err)
		}
	} else {
		return -1, fmt.Errorf("balance user (user_id = %s) not found", userID.String())
	}
	return value, nil
}

func (b *BalanceRepo) ReplenishmentBalance(ctx context.Context, replenishment models.Replenishment) error {
	query := `update balance set value = value + $1 where user_id = $2`

	rows, err := b.pg.Pool.Query(ctx, query, replenishment.Value, replenishment.UserID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
