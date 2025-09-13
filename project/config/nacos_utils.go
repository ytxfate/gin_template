package config

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gin_template/project/utils/request"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	nacosHost      string = "nacos-server-standalone"
	localNacosHost string = "test-nacos-server" // 本地测试使用
)

type nacosServerConfig struct {
	addr      string
	namespace string
	group     string
	username  string
	password  string
}

func NewNacosServerConfigTest() *nacosServerConfig {
	cfg := &nacosServerConfig{
		addr:      base64Decoder("MTI3LjAuMC4xOjg4NDg="), // 127.0.0.1:8848
		namespace: "public",
		group:     "DEFAULT_GROUP",
		username:  base64Decoder("bmFjb3M="), // nacos
		password:  base64Decoder("bmFjb3M="), // nacos
	}

	// 临时用于本地判断, 可移除
	if NacosHostLookup(localNacosHost, time.Second) {
		cfg.addr = localNacosHost + ":8848"
	}
	return cfg
}

func NewNacosServerConfigProd() *nacosServerConfig {
	return &nacosServerConfig{
		addr:      fmt.Sprintf("%s:8848", nacosHost),
		namespace: "public",
		group:     "DEFAULT_GROUP",
		username:  base64Decoder("bmFjb3M="), // nacos
		password:  base64Decoder("bmFjb3M="), // nacos
	}
}

func base64Decoder(s string) string {
	res, _ := base64.StdEncoding.DecodeString(s)
	return string(res)
}

type loginRes struct {
	AccessToken string `json:"accessToken"`
	// TokenTtl    int64  `json:"tokenTtl"`
	// GlobalAdmin bool   `json:"globalAdmin"`
	// Username    string `json:"username"`
}

// nacos 认证获取 token
func (nacosCfg *nacosServerConfig) nacosLogin() (token string, err error) {
	// curl -X POST 'http://127.0.0.1:8848/nacos/v3/auth/user/login' -d 'username=xxx&password=xxx'
	resp, err := request.NewRequest(
		fmt.Sprintf("http://%s/nacos/v3/auth/user/login", nacosCfg.addr),
		request.WithMethod(http.MethodPost),
		request.WithTimeout(time.Second*5),
		request.WithParams([][2]string{
			{"username", nacosCfg.username},
			{"password", nacosCfg.password},
		}),
	).Do()
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var logRes loginRes
	err = json.Unmarshal(data, &logRes)
	if err != nil {
		return
	}
	if logRes.AccessToken == "" {
		err = errors.New("request success, but no token return")
		return
	}
	token = logRes.AccessToken
	return
}

type cfgContent struct {
	Content string `json:"content"` //	配置内容
	// EncryptedDataKey string `json:"encryptedDataKey"` //	配置的加解密密钥，仅在使用配置加解密插件时有此值
	ContentType string `json:"contentType"` //	配置的类型，如TEXT,JSON等
	// Md5              string `json:"md5"`              //	配置的md5值
	// LastModified     int64  `json:"lastModified"`     //	配置的最后修改时间
	// Beta             bool   `json:"beta"`             //	配置是否有灰度配置
}

type confResp struct {
	// Code    int32      `json:"code"`
	// Message string     `json:"message"`
	Data cfgContent `json:"data"`
}

// 获取 Nacos 配置
func (nacosCfg *nacosServerConfig) getNacosConfig(accessToken, dataId string) (config string, configType string, err error) {
	// curl -X GET 'http://127.0.0.1:8848/nacos/v3/client/ns/instance/list?serviceName=quickstart.test.service&accessToken=eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ2b3lhZ2VyIiwiZXhwIjoxNzU2NzEwNjM1fQ.t7V7uLFL0y8bHSeZ-tMWykI6jlr0pcNpnR-b_LbpEis'
	resp, err := request.NewRequest(
		fmt.Sprintf(
			"http://%s/nacos/v3/client/cs/config?accessToken=%s&namespaceId=%s&groupName=%s&dataId=%s",
			nacosCfg.addr,
			accessToken,
			nacosCfg.namespace,
			nacosCfg.group,
			dataId,
		),
		request.WithTimeout(time.Second*5),
	).Do()
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var cr confResp
	err = json.Unmarshal(data, &cr)
	if err != nil {
		return
	}
	if cr.Data.Content == "" {
		err = errors.New("request success, but no config return")
		return
	}
	config = cr.Data.Content
	configType = cr.Data.ContentType
	return
}

func NacosHostLookup(hostname string, timeout time.Duration) bool {
	if timeout == 0 {
		timeout = time.Second
	}
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	_, err := net.DefaultResolver.LookupHost(ctx, hostname)
	return err == nil
}
