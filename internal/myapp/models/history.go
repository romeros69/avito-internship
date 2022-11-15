package models

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type History struct {
	ID                  uuid.UUID `json:"id"`                   // зачем это нужно пользователю?
	BalanceID           uuid.UUID `json:"balance_id"`           // надо всегда
	TypeHistory         string    `json:"type_history"`         // надо всегда
	OrderID             uuid.UUID `json:"order_id"`             // надо при reserve и confirmation
	ServiceID           uuid.UUID `json:"service_id"`           // надо при reserve и confirmation
	Value               int64     `json:"value"`                // надо всегда
	SourceReplenishment string    `json:"source_replenishment"` // надо при replenishment
	Date                time.Time `json:"date"`                 // надо всегда
}

func (h *History) StringTransaction(serviceName string) string {
	switch {
	case h.TypeHistory == "replenishment":
		return fmt.Sprintf("Date: %s, type of transaction: %s balance, Sourse replenishment: %s, value: %s", h.Date.String(), h.TypeHistory, h.SourceReplenishment, strconv.FormatInt(h.Value, 10))
	case h.TypeHistory == "reserve":
		return fmt.Sprintf("Date: %s, type of transaction: transfer money to %s , service name: %s, order id: %s, value: %s, ", h.Date.String(), h.TypeHistory, serviceName, h.OrderID.String(), strconv.FormatInt(h.Value, 10))
	case h.TypeHistory == "confirmation":
		return fmt.Sprintf("Date: %s, type of transaction: %s reserved money, service name: %s, order id: %s, value: %s", h.Date.String(), h.TypeHistory, serviceName, h.OrderID.String(), strconv.FormatInt(h.Value, 10))
	default:
		return ""
	}
}
