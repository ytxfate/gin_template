package auth

import (
	"gin_template/project/middleware"
	commresp "gin_template/project/utils/comm_resp"
	"gin_template/project/utils/jwttool"
	"net/http"

	"github.com/gin-gonic/gin"
)

// authInfo 登录信息
type authInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// loginHandler 登录接口
// @Summary 登录接口
// @Description 登录获取 jwt
// @Tags 认证
// @Accept application/json
// @Produce application/json
// @Param object body authInfo true "登录信息"
// @Success 200 {object} nil
// @Router /auth/login [post]
func loginHandler(c *gin.Context) {
	var ai authInfo
	err := c.ShouldBind(&ai)
	if err != nil || ai.Username == "" || ai.Password == "" {
		commresp.CommResp(c, commresp.ParameterError, nil, "无效的参数")
		return
	}
	// NOTE: 用户信息验证
	tk, refreshTk, err := jwttool.CreateJWTAndRefreshJWT(&jwttool.JWTInfo{UserName: ai.Username})
	if err != nil {
		commresp.CommResp(c, commresp.JwtCreateError, nil, "JWT 信息生成异常")
		return
	}
	// commresp.CommResp(c, commresp.Success, gin.H{
	// 	"jwt":         tk,
	// 	"refresh_jwt": refreshTk,
	// }, "OK")
	c.JSON(http.StatusOK, gin.H{"code": commresp.Success,
		"resp":         gin.H{"jwt": tk, "refresh_jwt": refreshTk},
		"msg":          "ok",
		"access_token": tk,
	})
}

// refreshInfo 刷新tokenn信息
type refreshInfo struct {
	RefreshJwt string `json:"refresh_jwt"`
}

// refreshTokenHandler 刷新token接口
// @Summary 刷新token接口
// @Description 刷新 jwt
// @Tags 认证
// @Accept application/json
// @Produce application/json
// @Param object body refreshInfo true "刷新tokenn信息"
// @Success 200 {object} nil
// @Router /auth/refresh_token [post]
// @Security OAuth2Password
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
