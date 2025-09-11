package config

import (
	"net"
	"sync"
)

var (
	Cfg  *Config
	once sync.Once
)

type Web struct {
	Title         string `yaml:"title"`
	Description   string `yaml:"description"`
	Addr          string `yaml:"addr"`
	SecretKey     string `yaml:"secretKey"`
	Version       string `yaml:"version"`
	ApiPrefixPath string `yaml:"apiPrefixPath"`
}

type webOption func(*Web)

func NewWeb(opts ...webOption) *Web {
	web := &Web{
		Title:         "gin_template",
		Description:   "gin 模板",
		Addr:          "0.0.0.0:8080",
		SecretKey:     "xxxxxxxx",
		Version:       "v1.0",
		ApiPrefixPath: "/api",
	}
	for _, opt := range opts {
		opt(web)
	}
	return web
}

func WithTitle(title string) webOption {
	return func(w *Web) { w.Title = title }
}
func WithDescription(description string) webOption {
	return func(w *Web) { w.Description = description }
}
func WithAddr(addr string) webOption {
	return func(w *Web) { w.Addr = addr }
}
func WithSecretKey(secretKey string) webOption {
	return func(w *Web) { w.SecretKey = secretKey }
}
func WithVersion(version string) webOption {
	return func(w *Web) { w.Version = version }
}
func WithApiPrefixPath(apiPrefixPath string) webOption {
	return func(w *Web) { w.ApiPrefixPath = apiPrefixPath }
}

type Config struct {
	Env        deployEnv
	Web        *Web `yaml:"web"`
	nacosCfg   *nacosServerConfig
	nacosToken string
}

func InitConfig(webCfg *Web, env deployEnv) {
	once.Do(func() {
		initConfig(webCfg, env)
	})
}

func initConfig(webCfg *Web, env deployEnv) {
	if !env.IsValid() {
		panic("env enum not match")
	}
	Cfg = &Config{
		Env:        env,
		Web:        webCfg,
		nacosCfg:   nil,
		nacosToken: "",
	}
	// 根据主机名自动判断一次运行环境(优先级最高)
	if Cfg.Env != PROD {
		_, err := net.LookupHost(nacosHost)
		if err == nil {
			Cfg.Env = PROD
			Cfg.nacosCfg = NewNacosServerConfigProd()
		}
	}

	if Cfg.nacosCfg == nil {
		Cfg.nacosCfg = NewNacosServerConfigTest()
	}

	// 获取 Nacos AccessToken
	var err error
	Cfg.nacosToken, err = Cfg.nacosCfg.nacosLogin()
	if err != nil {
		panic(err)
	}

	// NOTE: 初始化所有中间件配置
	initAllConfig()
}
