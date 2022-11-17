package reserve

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	ReserveBalance(c *gin.Context)
	AcceptReserve(c *gin.Context)
	CancelReserve(c *gin.Context)
}
