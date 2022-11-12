package v1

import (
	"avito-internship/internal/myapp/balance/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type balanceRoutes struct {
	bc usecase.BalanceContract
}

type balanceResponseDTO struct {
	Rubles  string `json:"rubles"`
	Pennies string `json:"pennies"`
}

func newBalanceRoutes(handler *gin.RouterGroup, bc usecase.BalanceContract) {
	br := &balanceRoutes{
		bc: bc,
	}
	handler.GET("/balance/:id", br.getBalanceByUserID)
}

func (br *balanceRoutes) getBalanceByUserID(c *gin.Context) {
	stringID := c.Param("id")
	userID, err := uuid.Parse(stringID)
	if err != nil || userID == uuid.Nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	value, err := br.bc.GetBalanceByUserID(c.Request.Context(), userID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, balanceToDTO(value))
}
