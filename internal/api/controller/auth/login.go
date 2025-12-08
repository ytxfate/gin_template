package auth

import (
	"gin_template/internal/api/middleware"
	commresp "gin_template/pkg/comm_resp"
	loginsrvc "gin_template/pkg/service/login_srvc"
	"net/http"

	"github.com/gin-gonic/gin"
)

// authInfo 登录信息
type authInfo struct {
	Username string `form:"username" json:"username" binding:"required,min=6"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

// LoginHandler 登录接口
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
func LoginHandler(c *gin.Context) {
	var ai authInfo
	err := c.ShouldBind(&ai)
	if err != nil {
		commresp.CommResp(c, commresp.ParameterError, err, "参数异常")
		return
	}
	tk, refreshTk, custErr := loginsrvc.Login(ai.Username, ai.Password)
	if custErr != nil {
		commresp.CommResp(c, custErr.Code, nil, custErr.Error())
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

// RefreshTokenHandler 刷新token接口
// @Summary 刷新token接口
// @Description 刷新 jwt
// @Tags 认证
// @Accept application/json
// @Produce application/json
// @Param object body refreshInfo true "刷新tokenn信息"
// @Success 200 {object} nil
// @Router /auth/refresh_token [post]
// @Security OAuth2Password
func RefreshTokenHandler(c *gin.Context) {
	tk := middleware.GetHeaderAuthToken(c)
	var ri refreshInfo
	err := c.ShouldBind(&ri)
	if err != nil {
		commresp.CommResp(c, commresp.ParameterError, err, "参数异常")
		return
	}

	newTk, newRefreshTk, custErr := loginsrvc.RefreshToken(tk, ri.RefreshJwt)
	if custErr != nil {
		commresp.CommResp(c, custErr.Code, nil, custErr.Error())
		return
	}

	commresp.CommResp(c, commresp.Success, gin.H{
		"jwt":         newTk,
		"refresh_jwt": newRefreshTk,
	}, "OK")
}
