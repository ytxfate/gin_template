package auth

import (
	"gin_template/project/middleware"
	commresp "gin_template/project/utils/comm_resp"
	"gin_template/project/utils/jwttool"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// authInfo 登录信息
type authInfo struct {
	Username string `form:"username" json:"username" binding:"required,min=6"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

// loginHandler 登录接口
// @Summary 登录接口
// @Description 登录获取 jwt
// @Description code == 1102 , 需刷新 jwt;
// @Description code == 1200 , 需重新登录后跳转;
// @Description code == 1101 , 再次请求; (基本不需要)
// @Tags 认证
// @Accept application/json
// @Produce application/json
// @Param object body authInfo true "登录信息"
// @Success 200 {object} nil
// @Router /auth/login [post]
func loginHandler(c *gin.Context) {
	var ai authInfo
	err := c.ShouldBind(&ai)
	if err != nil {
		commresp.CommResp(c, commresp.ParameterError, err, "参数异常")
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
	RefreshJwt string `json:"refresh_jwt" binding:"required,min=1"`
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
	_, jwtInfo, err := jwttool.ValidateJWT(middleware.GetHeaderAuthToken(c),
		jwt.WithoutClaimsValidation()) // 获取 jwt 存储信息, 不校验是否失效
	if err != nil {
		commresp.CommResp(c, commresp.UserLogout, nil, "登录失效")
		return
	}
	var ri refreshInfo
	err = c.ShouldBind(&ri)
	if err != nil {
		commresp.CommResp(c, commresp.ParameterError, err, "参数异常")
		return
	}
	stat := jwttool.ValidateRefreshJWT(middleware.GetHeaderAuthToken(c), ri.RefreshJwt)
	if !stat {
		commresp.CommResp(c, commresp.UserNoLogin, nil, "Refresh Token 验证失败")
		return
	}
	tk, refreshTk, err := jwttool.CreateJWTAndRefreshJWT(&jwtInfo)
	if err != nil {
		commresp.CommResp(c, commresp.JwtCreateError, nil, "JWT 信息生成异常")
		return
	}
	commresp.CommResp(c, commresp.Success, gin.H{
		"jwt":         tk,
		"refresh_jwt": refreshTk,
	}, "OK")
}
