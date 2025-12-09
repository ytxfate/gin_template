package webconfig

import (
	"gin_template/configs"
	"gin_template/pkg/deployenv"
	"gin_template/pkg/nacos"
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
	Env deployenv.DeployEnv
	Web *Web `yaml:"web"`
}

func InitConfig(webCfg *Web, env deployenv.DeployEnv) {
	once.Do(func() {
		initConfig(webCfg, env)
	})
}

func initConfig(webCfg *Web, env deployenv.DeployEnv) {
	if !env.IsValid() {
		panic("env enum not match")
	}
	env = nacos.InitNacos(env)
	Cfg = &Config{
		Env: env,
		Web: webCfg,
	}
	// 根据主
	// NOTE: 初始化所有中间件配置
	err := configs.InitAllDBConfig(Cfg.Env)
	if err != nil {
		panic(err)
	}
}
