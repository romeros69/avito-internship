package models

import (
	"github.com/google/uuid"
	"time"
)

type History struct {
	ID                  uuid.UUID `json:"id"`
	BalanceID           uuid.UUID `json:"balance_id"`
	TypeHistory         string    `json:"type_history"`
	OrderID             uuid.UUID `json:"order_id"`
	ServiceID           uuid.UUID `json:"service_id"`
	Value               int64     `json:"value"`
	SourceReplenishment string    `json:"source_replenishment"`
	Date                time.Time `json:"date"`
}
