package v1

import (
	"avito-internship/internal/myapp/history"
	"github.com/gin-gonic/gin"
)

func MapHistoryRoutes(handler *gin.RouterGroup, h history.Handlers) {
	handler.GET("/:id", h.GetTransactionInfoByUserID)
}
