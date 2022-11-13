package models

import "github.com/google/uuid"

type Replenishment struct {
	UserID uuid.UUID
	Value  int64
	Source string
}
