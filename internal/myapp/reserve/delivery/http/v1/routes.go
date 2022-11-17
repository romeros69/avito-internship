package v1

import (
	"avito-internship/internal/myapp/reserve"
	"github.com/gin-gonic/gin"
)

func MapReserveRoutes(handler *gin.RouterGroup, h reserve.Handlers) {
	handler.POST("", h.ReserveBalance)
	handler.POST("/accept", h.AcceptReserve)
	handler.POST("/cancel", h.CancelReserve)
}
