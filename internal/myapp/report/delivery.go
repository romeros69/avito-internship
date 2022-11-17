package report

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetReport(c *gin.Context)
}
