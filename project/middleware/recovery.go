package middleware

import (
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(400, gin.H{"msg": "ERROR"})
			}
		}()
		c.Next()
	}
}
