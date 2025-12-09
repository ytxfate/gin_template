package loginsrvc

import (
	"errors"
	"gin_template/internal/pkg/commresp"
	"gin_template/internal/pkg/customerrors"
	"gin_template/internal/pkg/jwttool"
	"gin_template/internal/repository/user"

	"github.com/golang-jwt/jwt/v5"
)

func Login(username, password string) (tk string, refreshTk string, custErr *customerrors.CustErr) {
	// TODO: 用户信息验证
	_, err := user.FindUserByUserPwd(username, password)
	if err != nil {
		custErr = &customerrors.CustErr{
			Code: commresp.JwtCreateError,
			Err:  errors.New("用户名密码错误"),
		}
		return
	}
	tk, refreshTk, err = jwttool.CreateJWTAndRefreshJWT(&jwttool.JWTInfo{UserName: username})
	if err != nil {
		custErr = &customerrors.CustErr{
			Code: commresp.JwtCreateError,
			Err:  errors.New("JWT 信息生成异常"),
		}
	}
	return
}

func RefreshToken(tk, refreshTk string) (newTk string, newRefreshTk string, custErr *customerrors.CustErr) {
	_, jwtInfo, err := jwttool.ValidateJWT(tk, jwt.WithoutClaimsValidation()) // 获取 jwt 存储信息, 不校验是否失效
	if err != nil {
		custErr = &customerrors.CustErr{
			Code: commresp.UserLogout,
			Err:  errors.New("登录失效"),
		}
		return
	}
	stat := jwttool.ValidateRefreshJWT(tk, refreshTk)
	if !stat {
		custErr = &customerrors.CustErr{
			Code: commresp.UserNoLogin,
			Err:  errors.New("refresh token 验证失败"),
		}
		return
	}
	newTk, newRefreshTk, err = jwttool.CreateJWTAndRefreshJWT(&jwtInfo)
	if err != nil {
		custErr = &customerrors.CustErr{
			Code: commresp.JwtCreateError,
			Err:  errors.New("JWT 信息生成异常"),
		}
		return
	}
	return
}
