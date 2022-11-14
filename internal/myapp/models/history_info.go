package models

import "github.com/google/uuid"

type HistoryInfo struct {
	BalanceID   uuid.UUID
	ServiceID   uuid.UUID
	OrderID     uuid.UUID
	TypeHistory string
}
