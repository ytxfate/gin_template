package user

import (
	commresp "gin_template/project/utils/comm_resp"
	"gin_template/project/utils/logger"

	"github.com/gin-gonic/gin"
)

func userHandler(c *gin.Context) {
	logger.Logger.Info("请求 userHandler")
	commresp.CommResp(c, commresp.Success, gin.H{"path": "user"}, "ok")
}

func user2Handler(c *gin.Context) {
	logger.Logger.Info("请求 user2Handler")
	panic("主动 panic")
}
