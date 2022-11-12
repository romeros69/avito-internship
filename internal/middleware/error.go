package middleware

import "github.com/gin-gonic/gin"

type errResponse struct {
	Error string `json:"error"`
}

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, errResponse{msg})
}
