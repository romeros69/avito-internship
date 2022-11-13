package models

import (
	"github.com/google/uuid"
	"time"
)

type History struct {
	ID                  uuid.UUID `json:"id"`
	BalanceID           uuid.UUID `json:"balance_id"`
	TypeHistory         string    `json:"type_history"`
	ReserveID           uuid.UUID `json:"reserve_id"`
	ReportID            uuid.UUID `json:"report_id"`
	SourceReplenishment string    `json:"source_replenishment"`
	Date                time.Time `json:"date"`
}
