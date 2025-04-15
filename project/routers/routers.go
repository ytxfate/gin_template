package routers

import (
	"gin_template/project/middleware"
	"gin_template/project/modules/auth"
	"gin_template/project/modules/user"

	"github.com/gin-gonic/gin"
)

type option func(*gin.RouterGroup)

// 注册app的路由配置
func include(opts ...option) (options []option) {
	options = append(options, opts...)
	return
}

// 初始化
func Init() *gin.Engine {
	engine := gin.New()

	engine.NoMethod(middleware.HandleRequestError)
	engine.NoRoute(middleware.HandleRequestError)
	// 中间件注册
	engine.Use(middleware.Logger(), middleware.Recovery())

	api := engine.Group("/api")
	options := include(auth.RouterGroup, user.RouterGroup)
	for _, opt := range options {
		opt(api)
	}
	return engine
}
