package v1

import (
	"avito-internship/internal/myapp/models"
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

func balanceToDTO(value int64) balanceResponseDTO {
	return balanceResponseDTO{
		Balance: strconv.FormatInt(value, 10),
	}
}

func replenishmentToEntity(dto replenishmentRequestDTO) (models.Replenishment, error) {
	userID, err := uuid.Parse(dto.UserID)
	if err != nil {
		return models.Replenishment{}, fmt.Errorf("error parsing user id in mapping dto replenishment: %w", err)
	}
	value, err := strconv.ParseInt(dto.Value, 10, 64)
	if err != nil {
		return models.Replenishment{}, fmt.Errorf("error parsing value in mapping dto replenishment: %w", err)
	}
	return models.Replenishment{
		UserID: userID,
		Value:  value,
		Source: dto.Source,
	}, nil
}
