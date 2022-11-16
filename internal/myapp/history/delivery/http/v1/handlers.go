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

func (h *historyHandlers) GetTransactionInfoByUserID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	orderBy := c.Param("orderBy")
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
	var response []historyResponseDTO
	for i, v := range historyTransfer.Histories {
		response = append(response, historyToDTO(historyTransfer.ServiceNames[i], v))
	}
	c.JSON(http.StatusOK, response)
}
