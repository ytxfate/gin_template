package middleware

import (
	"fmt"
	"gin_template/project/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		logger.Info(fmt.Sprintf("%-15s | %s %s %s | %3d | %s %s",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			c.Writer.Status(),
			cost,
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		))
	}
}
