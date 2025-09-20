package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindJSON[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set("req", req)
		c.Next()
	}
}
