package v1

import (
	"avito-internship/internal/myapp/history"
	"avito-internship/internal/myapp/middleware"
	"avito-internship/internal/myapp/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type historyHandlers struct {
	historyUC history.UseCase
}

func NewHistoryHandlers(historyUC history.UseCase) history.Handlers {
	return &historyHandlers{
		historyUC: historyUC,
	}
}

var _ history.Handlers = (*historyHandlers)(nil)

// GetTransactionInfoByUserID godoc
// @Summary GetHistoryByUserID
// @Tags history
// @Description Getting history transactions by user id
// @ID get-history
// @Accept json
// @Produce json
// @Param id path string true "Enter user id"
// @Param limit query string true "Enter limit"
// @Param page query string true "Enter number of page"
// @Param orderBy query string true "Enter sort type (date or value)"
// @Success 200 {object} centralHistoryResponseDTO
// @Failure 400 {object} middleware.errResponse
// @Failure 500 {object} middleware.errResponse
// @Router /api/v1/history/{id} [get]
func (h *historyHandlers) GetTransactionInfoByUserID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	orderBy := c.Query("orderBy")
	err = validate(orderBy, limit, page)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	historyTransfer, err := h.historyUC.GetHistoryByUserID(
		c.Request.Context(),
		models.Pagination{
			Size:    limit,
			Page:    page,
			OrderBy: orderBy,
		}, userID)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var historyList []historyResponseDTO
	for i, v := range historyTransfer.Histories {
		historyList = append(historyList, historyToDTO(historyTransfer.ServiceNames[i], v))
	}
	c.JSON(http.StatusOK, centralHistoryResponseDTO{
		History:   historyList,
		HasMore:   GetHasMore(int64(page), historyTransfer.Count, int64(limit)),
		TotalPage: GetTotalPage(historyTransfer.Count, int64(limit)),
	})
}
