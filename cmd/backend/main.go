package main

import (
	"context"
	"flag"
	"fmt"
	"gin_template/internal/backend/middleware"
	"gin_template/internal/backend/routers"
	"gin_template/internal/backend/webconfig"
	"gin_template/internal/configs"
	"gin_template/pkg/deployenv"
	"gin_template/pkg/logger"
	"gin_template/pkg/mongodb"
	"gin_template/pkg/nacos"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var ( // 构建信息记录
	gitHash      string
	buildTime    string
	goVersion    string
	remoteOrigin string
)

// NOTE: 从命令行初始化部分参数 及 初始化部分默认配置数值
var addr = flag.String("a", "0.0.0.0:8081", "服务监听地址[ip:port],缺失ip可能导致swagger文档异常")
var secretKey = flag.String("s", "xxxxxxxx", "密钥")
var version = flag.String("v", "v1.0", "版本")
var apiPrefixPath = flag.String("p", "/api", "接口前缀")
var env = flag.String("e", "DEV", "运行环境[生产环境判断优先]{ DEV/SIT/UAT/PRE_PROD/PROD }")
var buildInfo = flag.Bool("version", false, "输出版本构建信息")

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/v1.0/auth/login
// @scope {}
// @description OAuth protects our entity endpoints
func main() {
	flag.Parse()
	if *buildInfo {
		fmt.Printf("Build TimeStamp   : %s\n", buildTime)
		fmt.Printf("GoLang Version    : %s\n", goVersion)
		fmt.Printf("Git Remote Origin : %s\n", remoteOrigin)
		fmt.Printf("Git Commit Hash   : %s\n", gitHash)
		return
	}

	realEnv, err := deployenv.IsDeployEnv(*env)
	if err != nil {
		realEnv = deployenv.DEV
	}
	// 根据 Nacos 判断运行环境
	realEnv = nacos.InitNacos(realEnv)
	// 初始化 web 配置
	webconfig.InitConfig(webconfig.NewWeb(
		// NOTE: 从命令行初始化部分配置, 若不需要可以注释掉
		webconfig.WithAddr(*addr),
		webconfig.WithSecretKey(*secretKey),
		webconfig.WithVersion(*version),
		webconfig.WithApiPrefixPath(*apiPrefixPath),
	), realEnv)
	// 初始化日志服务
	logger.InitLogger(webconfig.Cfg.Env == deployenv.PROD)
	logger.Debugf("%s", webconfig.Cfg)
	// 初始化所有中间件配置
	err = configs.InitAllDBConfig(realEnv)
	if err != nil {
		panic(err)
	}
	// 参数校验中文翻译
	err = middleware.InitValidator("zh")
	if err != nil {
		logger.Fatal(err.Error())
	}

	// 数据库连接初始化
	err = mongodb.InitMongoDB(configs.MgConf)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// err = gaussdb.InitGaussDB(config.GaussCfg)
	// if err != nil {
	// 	logger.Fatal(err.Error())
	// }

	if webconfig.Cfg.Env == deployenv.PROD {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := routers.Init()
	/* =============== 优雅关停 =============== */
	srv := &http.Server{
		Addr:    webconfig.Cfg.Web.Addr,
		Handler: engine,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 关闭数据库连接
	err = mongodb.Close()
	if err != nil {
		logger.Error(err.Error())
	}
	// err = gaussdb.Close()
	// if err != nil {
	// 	logger.Error(err.Error())
	// }

	logger.Info("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("server exiting")
	/* ======================================= */
}
