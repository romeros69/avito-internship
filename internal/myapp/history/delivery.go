package history

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetTransactionInfoByUserID(c *gin.Context)
}
