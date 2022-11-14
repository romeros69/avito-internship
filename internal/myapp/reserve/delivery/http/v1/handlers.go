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

// @Summary ReserveBalance
// @Tags reserve
// @Description The method of reserving funds from the main balance in a separate account
// @ID reserve-balance
// @Accept json
// @Produce json
// @Param input body reserveBalanceDTO true "Enter user id, value, service id, order id"
// @Success 200 {object} nil
// @Failure 400 {object} middleware.errResponse
// @Failure 500 {object} middleware.errResponse
// @Router /api/v1/reserve [post]
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
	err = reserveInfoEntity.Validate()
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

// @Summary AcceptReserve
// @Tags reserve
// @Description Revenue recognition method - writes off money from the reserve, adds data to the report for accounting
// @ID accept-reverse
// @Accept json
// @Produce json
// @Param input body reserveBalanceDTO true "Enter user id, value, service id, order id"
// @Success 200 {object} nil
// @Failure 400 {object} middleware.errResponse
// @Failure 500 {object} middleware.errResponse
// @Router /api/v1/reserve/accept [post]
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
	err = reserveInfoEntity.Validate()
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
