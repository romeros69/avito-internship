package models

import "github.com/google/uuid"

type ReserveInfo struct {
	UserID    uuid.UUID
	ServiceID uuid.UUID
	OrderID   uuid.UUID
	Value     int64
}
