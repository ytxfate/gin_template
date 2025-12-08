package user

import (
	commresp "gin_template/project/utils/comm_resp"
	"gin_template/project/utils/logger"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户接口
// @Summary 用户接口
// @Description 用户模拟接口
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Success 200 {object} nil
// @Router /user/ [get]
func UserHandler(c *gin.Context) {
	logger.Info("请求 userHandler")
	commresp.CommResp(c, commresp.Success, gin.H{"path": "user"}, "ok")
}

// User2Handler 用户接口2
// @Summary 用户接口2
// @Description 模拟panic
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Success 200 {object} nil
// @Router /user/2 [get]
func User2Handler(c *gin.Context) {
	logger.Info("请求 user2Handler")
	panic("主动 panic")
}
