package repository

import (
	"avito-internship/internal/myapp/balance"
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type BalanceRepo struct {
	pg *postgres.Postgres
}

func NewBalanceRepo(pg *postgres.Postgres) *BalanceRepo {
	return &BalanceRepo{
		pg: pg,
	}
}

var _ balance.Repository = (*BalanceRepo)(nil)

func (b *BalanceRepo) CreateEmptyBalance(ctx context.Context, balance models.Balance) (uuid.UUID, error) {
	query := `insert into balance (id, user_id) VALUES ($1, $2) returning id`

	rows, err := b.pg.Pool.Query(ctx, query, balance.ID, balance.UserID)
	if err != nil {
		log.Println("cannot execute query: %w", err)
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		log.Println("error in creating new empty balance")
		return uuid.Nil, fmt.Errorf("error in creating new empty balance")
	}
	var balanceID uuid.UUID
	err = rows.Scan(&balanceID)
	if err != nil {
		log.Println("error in pasrsing balance id of creating balance")
		return uuid.Nil, fmt.Errorf("error in pasrsing balance id of creating balance")
	}
	return balanceID, nil
}

func (b *BalanceRepo) BalanceExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	query := `select * from balance where user_id = $1`

	rows, err := b.pg.Pool.Query(ctx, query, userID)
	if err != nil {
		log.Println("cannot execute query: %w", err)
		return false, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (b *BalanceRepo) GetBalanceByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	query := `select value from balance where user_id = $1`

	rows, err := b.pg.Pool.Query(ctx, query, userID)
	if err != nil {
		log.Println("cannot execute query: %w", err)
		return -1, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var value int64
	if rows.Next() {
		err = rows.Scan(&value)
		if err != nil {
			log.Println("error in parsing value balance by user id: %w", err)
			return -1, fmt.Errorf("error in parsing value balance by user id: %w", err)
		}
	} else {
		log.Printf("balance user (user_id = %s) not found", userID.String())
		return -1, fmt.Errorf("balance user (user_id = %s) not found", userID.String())
	}
	return value, nil
}

func (b *BalanceRepo) ReplenishmentBalance(ctx context.Context, replenishment models.Replenishment) (uuid.UUID, error) {
	query := `update balance set value = value + $1 where user_id = $2 returning id`

	rows, err := b.pg.Pool.Query(ctx, query, replenishment.Value, replenishment.UserID)
	if err != nil {
		log.Println("cannot execute query: %w", err)
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var balanceID uuid.UUID
	if !rows.Next() {
		log.Printf("error in replenishipment balance")
		return uuid.Nil, fmt.Errorf("error in replenishipment balance")
	}
	err = rows.Scan(&balanceID)
	if err != nil {
		log.Printf("error in parsing id balance of update balance")
		return uuid.Nil, fmt.Errorf("error in parsing id balance of update balance: %w", err)
	}
	return balanceID, nil
}

func (b *BalanceRepo) GetBalanceIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	query := `select id from balance where user_id = $1`

	rows, err := b.pg.Pool.Query(ctx, query, userID)
	if err != nil {
		log.Println("cannot execute query: %w", err)
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var balanceID uuid.UUID
	if rows.Next() {
		err = rows.Scan(&balanceID)
		if err != nil {
			log.Printf("error in parsing id balance by user id")
			return uuid.Nil, fmt.Errorf("error in parsing id balance by user id: %w", err)
		}
	} else {
		log.Printf("balance user (user_id = %s) not found", userID.String())
		return uuid.Nil, fmt.Errorf("balance user (user_id = %s) not found", userID.String())
	}
	return balanceID, nil
}

func (b *BalanceRepo) TransferBalance(ctx context.Context, balanceID uuid.UUID, value int64) error {
	query := `update balance set value = value - $1 where id = $2`

	rows, err := b.pg.Pool.Query(ctx, query, value, balanceID)
	if err != nil {
		log.Println("cannot execute query: %w", err)
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (b *BalanceRepo) ReturnMoneyFromReserve(ctx context.Context, balanceID uuid.UUID, value int64) error {
	query := `update balance set value = value + $1 where id=$2`

	rows, err := b.pg.Pool.Query(ctx, query, value, balanceID)
	if err != nil {
		log.Println("cannot execute query: %w", err)
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
