package main

import (
	"context"
	"fmt"
	"gin_template/project/config"
	"gin_template/project/middleware"
	"gin_template/project/routers"
	"gin_template/project/utils/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/v1.0/auth/login
// @scope {}
// @description OAuth protects our entity endpoints
func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	logger.InitLogger()
	middleware.InitValidator()

	logger.Logger.Debug(fmt.Sprintf("%#v", config.Cfg))

	engine := routers.Init()
	/* =============== 优雅关停 =============== */
	srv := &http.Server{
		Addr:    config.Cfg.Web.Addr,
		Handler: engine,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatal(err.Error())
	}
	logger.Logger.Info("server exiting")
	/* ======================================= */
}
