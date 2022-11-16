package v1

import "fmt"

func validate(orderBy string, limit, page int) error {
	switch {
	case limit < 1:
		return fmt.Errorf("invalid linit value, limit must be a positive")
	case page < 1:
		return fmt.Errorf("invalid page value, limit must be a positive")
	case orderBy != "value" && orderBy != "date":
		return fmt.Errorf("invalid order by - order by must be a \"value\" or \"date\"")
	default:
		return nil
	}
}
