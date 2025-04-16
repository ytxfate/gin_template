package auth

import (
	"gin_template/project/utils/logger"

	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	logger.Logger.Info("请求 loginHandler")
	c.JSON(200, gin.H{"username": "xxx", "password": "yyy"})
}
