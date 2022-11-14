package models

import (
	"fmt"
	"github.com/google/uuid"
)

type ReserveInfo struct {
	UserID    uuid.UUID
	ServiceID uuid.UUID
	OrderID   uuid.UUID
	Value     int64
}

func (r *ReserveInfo) Validate() error {
	if r.Value < 0 {
		return fmt.Errorf("value could not be a negative")
	}
	return nil
}
