package v1

import (
	"strconv"
)

func balanceToDTO(value int64) balanceResponseDTO {
	return balanceResponseDTO{
		Rubles:  strconv.FormatInt(value/100, 10),
		Pennies: strconv.FormatInt(value%100, 10),
	}
}
