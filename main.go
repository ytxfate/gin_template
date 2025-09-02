package main

import (
	"context"
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

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/v1.0/auth/login
// @scope {}
// @description OAuth protects our entity endpoints
func main() {
	config.InitConfig(config.NewWeb(config.WithAddr("0.0.0.0:8081")), config.DEV)
	logger.InitLogger(config.Cfg.Env == config.PROD)
	logger.Debugf("%#v", config.Cfg)
	middleware.InitValidator()

	// 数据库连接初始化
	err := operatemongodb.InitMongoDB(config.MgConf)
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
