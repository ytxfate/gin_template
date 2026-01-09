package loginsrvc

import (
	"errors"
	"gin_template/internal/backend/jwttool"
	"gin_template/internal/backend/webconfig"
	"gin_template/internal/pkg/commresp"
	"gin_template/internal/pkg/custerr"
	"gin_template/internal/pkg/repository/user"

	"github.com/golang-jwt/jwt/v5"
)

func Login(username, password string) (tk string, refreshTk string, custErr *custerr.CustErr) {
	// TODO: 用户信息验证
	_, err := user.FindUserByUserPwd(username, password)
	if err != nil {
		custErr = &custerr.CustErr{
			Code: commresp.JwtCreateError,
			Err:  errors.New("用户名密码错误"),
		}
		return
	}
	tk, refreshTk, err = jwttool.GenerateJWTAndRefreshJWT(&jwttool.JWTInfo{UserName: username}, webconfig.Cfg.Web.SecretKey)
	if err != nil {
		custErr = &custerr.CustErr{
			Code: commresp.JwtCreateError,
			Err:  errors.New("JWT 信息生成异常"),
		}
	}
	return
}

func RefreshToken(tk, refreshTk string) (newTk string, newRefreshTk string, custErr *custerr.CustErr) {
	_, jwtInfo, err := jwttool.ValidateJWT(tk, webconfig.Cfg.Web.SecretKey, jwt.WithoutClaimsValidation()) // 获取 jwt 存储信息, 不校验是否失效
	if err != nil {
		custErr = &custerr.CustErr{
			Code: commresp.UserLogout,
			Err:  errors.New("登录失效"),
		}
		return
	}
	stat := jwttool.ValidateRefreshJWT(tk, refreshTk, webconfig.Cfg.Web.SecretKey)
	if !stat {
		custErr = &custerr.CustErr{
			Code: commresp.UserNoLogin,
			Err:  errors.New("refresh token 验证失败"),
		}
		return
	}
	newTk, newRefreshTk, err = jwttool.GenerateJWTAndRefreshJWT(&jwtInfo, webconfig.Cfg.Web.SecretKey)
	if err != nil {
		custErr = &custerr.CustErr{
			Code: commresp.JwtCreateError,
			Err:  errors.New("JWT 信息生成异常"),
		}
		return
	}
	return
}
