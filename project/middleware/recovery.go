package middleware

import (
	"fmt"
	"gin_template/project/utils/logger"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// TODO: 日志优化, 打印错误栈信息
				logger.Logger.Error(fmt.Sprintf("%v", err))
				c.AbortWithStatusJSON(400, gin.H{"msg": "ERROR"})
			}
		}()
		c.Next()
	}
}
