package balance

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetBalanceByUserID(c *gin.Context)
}
