package v1

import (
	"avito-internship/internal/myapp/models"
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

func reserveInfoToEntity(dto reserveBalanceDTO) (models.ReserveInfo, error) {
	userID, err := uuid.Parse(dto.UserID)
	if err != nil {
		return models.ReserveInfo{}, fmt.Errorf("error parsing user id in mapping dto reserve info: %w", err)
	}
	serviceID, err := uuid.Parse(dto.ServiceID)
	if err != nil {
		return models.ReserveInfo{}, fmt.Errorf("error parsing service id in mapping dto reserve info: %w", err)
	}
	orderID, err := uuid.Parse(dto.OrderID)
	if err != nil {
		return models.ReserveInfo{}, fmt.Errorf("error parsing irder id in mapping dto reserve info: %w", err)
	}
	value, err := strconv.ParseInt(dto.Value, 10, 64)
	if err != nil {
		return models.ReserveInfo{}, fmt.Errorf("error parsing value in mapping dto reserve info: %w", err)
	}
	return models.ReserveInfo{
		UserID:    userID,
		ServiceID: serviceID,
		OrderID:   orderID,
		Value:     value,
	}, nil
}
