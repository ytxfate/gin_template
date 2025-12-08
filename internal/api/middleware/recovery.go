package middleware

import (
	"fmt"
	commresp "gin_template/project/utils/comm_resp"
	"gin_template/project/utils/logger"
	"runtime"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印错误栈信息
				buf := make([]byte, 1<<16) // 65536
				runtime.Stack(buf, false)
				logger.Error(fmt.Sprintf("%v\n%v", err, string(buf)))
				commresp.CommResp(c, commresp.ExceptionError, nil, "ERROR")
				c.Abort()
			}
		}()
		c.Next()
	}
}
