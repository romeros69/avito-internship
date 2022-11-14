package models

import "github.com/google/uuid"

type Reserve struct {
	ID        uuid.UUID `json:"id"`
	BalanceID uuid.UUID `json:"balance_id"`
	Value     int64     `json:"value"`
}
