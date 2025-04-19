package middleware

import (
	commresp "gin_template/project/utils/comm_resp"
	"gin_template/project/utils/jwttool"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetHeaderAuthToken(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	return parts[1]
}

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := GetHeaderAuthToken(c)
		if token == "" {
			commresp.CommResp(c, commresp.UserNoLogin, nil, "未登录")
			c.Abort()
			return
		}
		stat, jwtInfo, err := jwttool.ValidateJWT(token)
		if !stat || err != nil {
			commresp.CommResp(c, commresp.UserNoLogin, nil, "未登录.")
			c.Abort()
			return
		}
		c.Set("jwtInfo", &jwtInfo)
		c.Next()
	}
}
