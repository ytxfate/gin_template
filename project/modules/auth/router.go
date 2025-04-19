package auth

import (
	"gin_template/project/middleware"

	"github.com/gin-gonic/gin"
)

func RouterGroup(e *gin.RouterGroup) {
	authrouter := e.Group("/auth")
	authrouter.POST("/login", loginHandler)
	authrouter.POST("/refresh_token", middleware.AuthMiddleware(), refreshTokenHandler)
}
