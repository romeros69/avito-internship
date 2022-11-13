package models

import (
	"github.com/google/uuid"
	"time"
)

type Report struct {
	ID        uuid.UUID `json:"id"`
	ServiceID uuid.UUID `json:"service_id"`
	OrderID   uuid.UUID `json:"order_id"`
	Value     int64     `json:"value"`
	Date      time.Time `json:"date"`
}
