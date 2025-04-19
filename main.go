package main

import (
	"context"
	"gin_template/project/routers"
	"gin_template/project/utils/logger"
	"log"
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

	engine := routers.Init()
	/* =============== 优雅关停 =============== */
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting")
	/* ======================================= */
}
