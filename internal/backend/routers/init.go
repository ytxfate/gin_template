package routers

import (
	"gin_template/internal/backend/docs"
	"gin_template/internal/backend/middleware"
	"gin_template/internal/backend/webconfig"
	"gin_template/pkg/deployenv"
	"gin_template/pkg/logger"

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

	prefix := webconfig.Cfg.Web.ApiPrefixPath + "/" + webconfig.Cfg.Web.Version
	api := engine.Group(prefix)
	options := include(routerGroup)
	for _, opt := range options {
		opt(api)
	}
	if webconfig.Cfg.Env != deployenv.PROD {
		// 设置 swagger 文档信息
		docs.SwaggerInfo.Title = webconfig.Cfg.Web.Title
		docs.SwaggerInfo.Description = webconfig.Cfg.Web.Description
		docs.SwaggerInfo.BasePath = prefix
		docs.SwaggerInfo.Version = webconfig.Cfg.Web.Version
		docs.SwaggerInfo.Host = webconfig.Cfg.Web.Addr
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		logger.Debugf("swagger docs: http://%s%s/swagger/index.html", webconfig.Cfg.Web.Addr, prefix)
	}
	return engine
}
