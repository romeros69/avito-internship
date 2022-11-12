package v1

import (
	"avito-internship/internal/myapp/balance/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, bc usecase.BalanceContract) {
	h := handler.Group("/api/v1")
	{
		newBalanceRoutes(h, bc)
	}
}
