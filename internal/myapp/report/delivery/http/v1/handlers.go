package v1

import (
	"avito-internship/internal/myapp/middleware"
	"avito-internship/internal/myapp/report"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type reportHandlers struct {
	reportUC report.UseCase
}

func NewReportHandlers(reportUC report.UseCase) report.Handlers {
	return &reportHandlers{
		reportUC: reportUC,
	}
}

var _ report.Handlers = (*reportHandlers)(nil)

// GetReport godoc
// @Summary GetReportByYearMonth
// @Tags report
// @Description Getting report by year and month
// @ID get-report
// @Accept json
// @Produce json
// @Param year query string true "Enter year (number)"
// @Param month query string true "Enter month (number)"
// @Success 200 {object} linkReportDTO
// @Failure 400 {object} middleware.errResponse
// @Failure 500 {object} middleware.errResponse
// @Router /api/v1/report [get]
func (r *reportHandlers) GetReport(c *gin.Context) {
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = validateYearMonth(year, month)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = r.reportUC.GetReport(c.Request.Context(), year-1, month-1)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, linkReportDTO{"localhost:9000/api/v1/report.csv"})
}
