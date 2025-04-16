package user

import (
	"gin_template/project/utils/logger"

	"github.com/gin-gonic/gin"
)

func userHandler(c *gin.Context) {
	logger.Logger.Info("请求 userHandler")
	c.JSON(200, gin.H{"path": "user"})
}

func user2Handler(c *gin.Context) {
	logger.Logger.Info("请求 user2Handler")
	panic("主动 panic")
}
