package webconfig

import (
	"gin_template/pkg/config"
	"sync"
	"time"
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
	Env        config.DeployEnv
	Web        *Web `yaml:"web"`
	nacosCfg   *config.NacosServerConfig
	nacosToken string
}

func InitConfig(webCfg *Web, env config.DeployEnv) {
	once.Do(func() {
		initConfig(webCfg, env)
	})
}

func initConfig(webCfg *Web, env config.DeployEnv) {
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
	if Cfg.Env != config.PROD {
		if config.NacosHostLookup(config.ProdNacosHost, time.Second) {
			Cfg.Env = config.PROD
			Cfg.nacosCfg = config.NewNacosServerConfigProd()
		}
	}

	if Cfg.nacosCfg == nil {
		Cfg.nacosCfg = config.NewNacosServerConfigTest()
	}

	// 获取 Nacos AccessToken
	var err error
	Cfg.nacosToken, err = Cfg.nacosCfg.NacosLogin()
	if err != nil {
		panic(err)
	}

	// NOTE: 初始化所有中间件配置
	config.InitAllDBConfig(Cfg.Env, *Cfg.nacosCfg, Cfg.nacosToken)
}
