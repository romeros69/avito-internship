package repo

import (
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type BalanceRp interface {
	GetBalanceByUserID(context.Context, uuid.UUID) (int64, error)
}

type BalanceRepo struct {
	pg *postgres.Postgres
}

func NewBalanceRepo(pg *postgres.Postgres) *BalanceRepo {
	return &BalanceRepo{
		pg: pg,
	}
}

var _ BalanceRp = (*BalanceRepo)(nil)

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
