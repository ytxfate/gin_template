package main

import (
    "context"
    "gin_template/middleware"
    "gin_template/register_rgs"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

var router *gin.Engine

func main() {
    //router = gin.Default()
    router := gin.New()

    router.NoMethod(middleware.HandleRequestError)
    router.NoRoute(middleware.HandleRequestError)

    router.Use(middleware.Logger(), middleware.Recovery())

    // 注册路由
    api := router.Group("/api")
    register_rgs.RegisterRouterGroups(api)

    /* =============== 优雅关停 =============== */
    srv := &http.Server{
        Addr: ":8080",
        Handler: router,
    }
    go func() {
        err:=srv.ListenAndServe()
        if err != nil && err != http.ErrServerClosed {
            log.Fatalf("Listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("shutdown server...")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx);err != nil {
        log.Fatal("server shutdown:", err)
    }
    log.Println("server exiting")
    /* ======================================= */
}