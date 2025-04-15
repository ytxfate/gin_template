package main

import (
	"context"
	"gin_template/project/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
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
