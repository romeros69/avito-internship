package v1

import (
	"avito-internship/internal/myapp/report"
	"github.com/gin-gonic/gin"
)

func MapReportRoutes(handler *gin.RouterGroup, h report.Handlers) {
	handler.GET("", h.GetReport)
}
