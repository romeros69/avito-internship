package v1

import (
	"avito-internship/internal/myapp/balance"
	"github.com/gin-gonic/gin"
)

func MapBalanceRoutes(handler *gin.RouterGroup, h balance.Handlers) {
	handler.GET("/:id", h.GetBalanceByUserID)
	handler.POST("", h.ReplenishmentBalance)
}
