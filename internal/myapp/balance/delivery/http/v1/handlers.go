package v1

import (
	"avito-internship/internal/myapp/balance"
	"avito-internship/internal/myapp/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type balanceHandlers struct {
	balanceUC balance.UseCase
}

func NewBalanceHandlers(balanceUC balance.UseCase) balance.Handlers {
	return &balanceHandlers{
		balanceUC: balanceUC,
	}
}

var _ balance.Handlers = (*balanceHandlers)(nil)

func (b *balanceHandlers) GetBalanceByUserID(c *gin.Context) {
	stringID := c.Param("id")
	userID, err := uuid.Parse(stringID)
	if err != nil || userID == uuid.Nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	value, err := b.balanceUC.GetBalanceByUserID(c.Request.Context(), userID)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, balanceToDTO(value))
}

func (b *balanceHandlers) ReplenishmentBalance(c *gin.Context) {
	req := new(replenishmentRequestDTO)
	if err := c.ShouldBindJSON(req); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	replenishmentEntity, err := replenishmentToEntity(*req)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = b.balanceUC.ReplenishmentBalance(c.Request.Context(), replenishmentEntity)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
