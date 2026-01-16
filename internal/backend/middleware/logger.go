package middleware

import (
	"fmt"
	"gin_template/internal/backend/docs"
	"gin_template/pkg/logger"
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
		logger.Debugf("%+v",
			docs.GetApiTagsAndSummary(c.Request.URL.Path, c.Request.Method))
	}
}
