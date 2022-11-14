package v1

import (
	"avito-internship/internal/myapp/middleware"
	"avito-internship/internal/myapp/reserve"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reserveHandlers struct {
	reserveUC reserve.UseCase
}

func NewReserveHandlers(reserveUC reserve.UseCase) reserve.Handlers {
	return &reserveHandlers{
		reserveUC: reserveUC,
	}
}

var _ reserve.Handlers = (*reserveHandlers)(nil)

func (r *reserveHandlers) ReserveBalance(c *gin.Context) {
	req := new(reserveBalanceDTO)
	if err := c.ShouldBindJSON(req); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	reserveInfoEntity, err := reserveInfoToEntity(*req)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = r.reserveUC.ReserveBalance(c.Request.Context(), reserveInfoEntity)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (r *reserveHandlers) AcceptReserve(c *gin.Context) {
	req := new(reserveBalanceDTO)
	if err := c.ShouldBindJSON(req); err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	reserveInfoEntity, err := reserveInfoToEntity(*req)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = r.reserveUC.AcceptReserve(c.Request.Context(), reserveInfoEntity)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
