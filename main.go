package main

import (
	"context"
	"gin_template/project/middleware"
	"gin_template/project/routers"
	"gin_template/project/utils/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title gin_template
// @version 1.0
// @description gin 模板
// @host 127.0.0.1:8080
// @BasePath /api
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/auth/login
// @scope {}
// @description OAuth protects our entity endpoints
func main() {
	logger.InitLogger()
	middleware.InitValidator()

	engine := routers.Init()
	/* =============== 优雅关停 =============== */
	srv := &http.Server{
		Addr:    ":8080",
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
