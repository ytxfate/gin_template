package jwttool

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwt 存储信息
type JWTInfo struct {
	UserName string   `json:"username"`
	Scopes   []string `json:"scopes"`
}

const defaultValidTime time.Duration = time.Minute * 30

type claims struct {
	jwt.RegisteredClaims
	Data JWTInfo `json:"data"`
}

type refreshClaims struct {
	jwt.RegisteredClaims
	EncryptJwt string `json:"encrypt_jwt"`
}

func sha256Encrypt(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

// 创建 JWT
func generateJWT(c jwt.Claims, secret string) (string, error) {
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 生成签名字符串
	return token.SignedString([]byte(secret))
}

func GenerateJWTAndRefreshJWT(info *JWTInfo, secret string) (tk string, refreshTk string, err error) {
	now := time.Now()
	jwtClaims := &claims{
		Data: *info,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(defaultValidTime)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(now),                       // 签发时间
			Issuer:    "hsd",                                         // 签发人
		},
	}
	tk, err = generateJWT(jwtClaims, secret)
	if err != nil {
		return
	}
	refreshClaims := &refreshClaims{
		EncryptJwt: sha256Encrypt(tk),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(defaultValidTime * 2)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(now),                           // 签发时间
			Issuer:    "hsd",                                             // 签发人
		},
	}
	refreshTk, err = generateJWT(refreshClaims, secret)
	return
}

// 解析/校验 JWT
func ValidateJWT(tk, secret string, options ...jwt.ParserOption) (stat bool, jwtInfo JWTInfo, err error) {
	// 解析token
	var (
		c     claims
		token *jwt.Token
	)
	token, err = jwt.ParseWithClaims(tk, &c,
		func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil },
		options...,
	)
	if err != nil { // 解析token失败
		return
	}
	return token.Valid, c.Data, nil
}

func ValidateRefreshJWT(priTk, refreTk, secret string) bool {
	// 解析token
	var (
		rc  refreshClaims
		err error
	)
	_, err = jwt.ParseWithClaims(refreTk, &rc, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err == nil && rc.EncryptJwt == sha256Encrypt(priTk) {
		return true
	}
	return false
}
