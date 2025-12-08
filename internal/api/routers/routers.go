package routers

import (
	"gin_template/internal/api/controller/auth"
	"gin_template/internal/api/controller/user"
	"gin_template/internal/api/docs"
	"gin_template/internal/api/middleware"
	"gin_template/project/config"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	engine.Use(middleware.Recovery())
	engine.Use(middleware.Logger())
	engine.Use(middleware.CORS())

	prefix := config.Cfg.Web.ApiPrefixPath + "/" + config.Cfg.Web.Version
	api := engine.Group(prefix)
	options := include(auth.RouterGroup, user.RouterGroup)
	for _, opt := range options {
		opt(api)
	}
	if config.Cfg.Env != config.PROD {
		// 设置 swagger 文档信息
		docs.SwaggerInfo.Title = config.Cfg.Web.Title
		docs.SwaggerInfo.Description = config.Cfg.Web.Description
		docs.SwaggerInfo.BasePath = prefix
		docs.SwaggerInfo.Version = config.Cfg.Web.Version
		docs.SwaggerInfo.Host = config.Cfg.Web.Addr
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return engine
}
