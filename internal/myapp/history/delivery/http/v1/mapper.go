package v1

import (
	"avito-internship/internal/myapp/models"
	"strconv"
)

func historyToDTO(serviceName string, history models.History) historyResponseDTO {

	return historyResponseDTO{
		Date:                history.Date.String(),
		TypeOfTransaction:   history.TypeHistory,
		SourceReplenishment: history.SourceReplenishment, // вот это проверить
		Value:               strconv.FormatInt(history.Value, 10),
		ServiceName:         serviceName,
		OrderID: func(typeTransactional string) string {
			if typeTransactional != "replenishment" {
				return history.OrderID.String()
			}
			return ""
		}(history.TypeHistory),
	}
}
