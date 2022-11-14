package models

import (
	"fmt"
	"github.com/google/uuid"
	"unicode/utf8"
)

type Replenishment struct {
	UserID uuid.UUID
	Value  int64
	Source string
}

func (r *Replenishment) Validate() error {
	switch {
	case r.Value < 0:
		return fmt.Errorf("value could not be a negative")
	case utf8.RuneCountInString(r.Source) != 16:
		return fmt.Errorf("invalid source card number")
	}
	return nil
}
