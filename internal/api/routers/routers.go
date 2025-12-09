package routers

import (
	"gin_template/internal/api/controller/auth"
	"gin_template/internal/api/controller/user"

	"github.com/gin-gonic/gin"
)

func routerGroup(e *gin.RouterGroup) {
	authRouter := e.Group("/auth")
	{
		authRouter.POST("/login", auth.LoginHandler)
		authRouter.POST("/refresh_token", auth.RefreshTokenHandler)
	}

	userRouter := e.Group("/user")
	{
		userRouter.GET("/", user.UserHandler)
		userRouter.GET("/2", user.User2Handler)
	}
}
