package entity

import "github.com/google/uuid"

type Balance struct {
	ID     uuid.UUID `json:"id"`
	Value  int64     `json:"value"`
	UserID uuid.UUID `json:"user_id"`
}
