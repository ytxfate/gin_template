package main

import (
	"context"
	"flag"
	"gin_template/project/config"
	"gin_template/project/middleware"
	"gin_template/project/routers"
	"gin_template/project/utils/logger"
	operategaussdb "gin_template/project/utils/operate_gaussdb"
	operatemongodb "gin_template/project/utils/operate_mongodb"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// NOTE: 从命令行初始化部分参数 及 初始化部分默认配置数值
var addr = flag.String("a", "0.0.0.0:8081", "服务监听地址[ip:port],缺失ip可能导致swagger文档异常")
var secretKey = flag.String("s", "xxxxxxxx", "密钥")
var version = flag.String("v", "v1.0", "版本")
var apiPrefixPath = flag.String("p", "/api", "接口前缀")
var env = flag.String("e", "DEV", "运行环境[生产环境判断优先]{ DEV/SIT/UAT/PRE_PROD/PROD }")

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/v1.0/auth/login
// @scope {}
// @description OAuth protects our entity endpoints
func main() {
	flag.Parse()
	realEnv, err := config.IsDeployEnv(*env)
	if err != nil {
		realEnv = config.DEV
	}
	config.InitConfig(config.NewWeb(
		// NOTE: 从命令行初始化部分配置, 若不需要可以注释掉
		config.WithAddr(*addr),
		config.WithSecretKey(*secretKey),
		config.WithVersion(*version),
		config.WithApiPrefixPath(*apiPrefixPath),
	), realEnv)
	logger.InitLogger(config.Cfg.Env == config.PROD)
	logger.Debugf("%#v", config.Cfg)
	err = middleware.InitValidator("zh")
	if err != nil {
		logger.Fatal(err.Error())
	}

	// 数据库连接初始化
	err = operatemongodb.InitMongoDB(config.MgConf)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = operategaussdb.InitGaussDB(config.GaussCfg)
	if err != nil {
		logger.Fatal(err.Error())
	}

	if config.Cfg.Env == config.PROD {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := routers.Init()
	/* =============== 优雅关停 =============== */
	srv := &http.Server{
		Addr:    config.Cfg.Web.Addr,
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
	err = operatemongodb.Close()
	if err != nil {
		logger.Error(err.Error())
	}
	err = operategaussdb.Close()
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("server exiting")
	/* ======================================= */
}
