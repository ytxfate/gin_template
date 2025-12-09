package jwttool

import (
	"crypto/sha256"
	"encoding/hex"
	webconfig "gin_template/internal/api/web-config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtValidTime time.Duration = time.Minute * 30

// jwt 存储信息
type JWTInfo struct {
	UserName string   `json:"username"`
	Scopes   []string `json:"scopes"`
}

type jwtBodyInfo struct {
	jwt.RegisteredClaims
	Data JWTInfo `json:"data"`
}

type refreshJWTBodyInfo struct {
	jwt.RegisteredClaims
	EncryptJwt string `json:"encrypt_jwt"`
}

func jwtEncryptHandler(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

// 创建 JWT
func createJWT(claims jwt.Claims) (string, error) {
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString([]byte(webconfig.Cfg.Web.SecretKey))
}

func CreateJWTAndRefreshJWT(jwtInfo *JWTInfo) (tk string, refreshTk string, err error) {
	nowTime := time.Now()
	jwtClaims := &jwtBodyInfo{
		Data: *jwtInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(nowTime.Add(jwtValidTime)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(nowTime),                   // 签发时间
			Issuer:    "hsd",                                         // 签发人
		},
	}
	tk, err = createJWT(jwtClaims)
	if err != nil {
		return
	}
	refreshClaims := &refreshJWTBodyInfo{
		EncryptJwt: jwtEncryptHandler(tk),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(nowTime.Add(jwtValidTime * 2)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(nowTime),                       // 签发时间
			Issuer:    "hsd",                                             // 签发人
		},
	}
	refreshTk, err = createJWT(refreshClaims)
	return
}

// 解析/校验 JWT
func ValidateJWT(tk string, options ...jwt.ParserOption) (stat bool, jwtInfo JWTInfo, err error) {
	// 解析token
	var (
		jbi   jwtBodyInfo
		token *jwt.Token
	)
	token, err = jwt.ParseWithClaims(tk, &jbi,
		func(token *jwt.Token) (interface{}, error) { return []byte(webconfig.Cfg.Web.SecretKey), nil },
		options...,
	)
	if err != nil { // 解析token失败
		return
	}
	return token.Valid, jbi.Data, nil
}

func ValidateRefreshJWT(priTk, refreTk string) bool {
	// 解析token
	var (
		rjbi refreshJWTBodyInfo
		err  error
	)
	_, err = jwt.ParseWithClaims(refreTk, &rjbi, func(token *jwt.Token) (interface{}, error) {
		return []byte(webconfig.Cfg.Web.SecretKey), nil
	})
	if err == nil && rjbi.EncryptJwt == jwtEncryptHandler(priTk) {
		return true
	}
	return false
}
