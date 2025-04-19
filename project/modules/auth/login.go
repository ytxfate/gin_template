package auth

import (
	"gin_template/project/middleware"
	commresp "gin_template/project/utils/comm_resp"
	"gin_template/project/utils/jwttool"

	"github.com/gin-gonic/gin"
)

type authInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	var ai authInfo
	err := c.ShouldBind(&ai)
	if err != nil {
		commresp.CommResp(c, commresp.ParameterError, nil, "无效的参数")
		return
	}
	// NOTE: 用户信息验证
	tk, refreshTk, err := jwttool.CreateJWTAndRefreshJWT(&jwttool.JWTInfo{UserName: ai.Username})
	if err != nil {
		commresp.CommResp(c, commresp.JwtCreateError, nil, "JWT 信息生成异常")
		return
	}
	commresp.CommResp(c, commresp.Success, gin.H{
		"jwt":         tk,
		"refresh_jwt": refreshTk,
	}, "OK")
}

type refreshInfo struct {
	RefreshJwt string `json:"refresh_jwt"`
}

func refreshTokenHandler(c *gin.Context) {
	var ri refreshInfo
	err := c.ShouldBind(&ri)
	if err != nil || ri.RefreshJwt == "" {
		commresp.CommResp(c, commresp.ParameterError, nil, "无效的参数")
		return
	}
	stat := jwttool.ValidateRefreshJWT(middleware.GetHeaderAuthToken(c), ri.RefreshJwt)
	if !stat {
		commresp.CommResp(c, commresp.UserNoLogin, nil, "Refresh Token 验证失败")
		return
	}
	jwtInfo := c.MustGet("jwtInfo").(*jwttool.JWTInfo)
	tk, refreshTk, err := jwttool.CreateJWTAndRefreshJWT(jwtInfo)
	if err != nil {
		commresp.CommResp(c, commresp.JwtCreateError, nil, "JWT 信息生成异常")
		return
	}
	commresp.CommResp(c, commresp.Success, gin.H{
		"jwt":         tk,
		"refresh_jwt": refreshTk,
	}, "OK")
}
