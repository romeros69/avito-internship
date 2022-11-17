package v1

import (
	"fmt"
	"time"
)

func validateYearMonth(year, month int) error {
	switch {
	case year > time.Now().Year():
		return fmt.Errorf("invalid year: the entered year is greater than the current one")
	case month > 12:
		return fmt.Errorf("invalid month: maximum month number - 12")
	case month > int(time.Now().Month()):
		return fmt.Errorf("invalid month: the entered year is greater than the current one")
	}
	return nil
}
