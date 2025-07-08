package auth

import (
	"github.com/gin-gonic/gin"
)

func RouterGroup(e *gin.RouterGroup) {
	authrouter := e.Group("/auth")
	authrouter.POST("/login", loginHandler)
	authrouter.POST("/refresh_token", refreshTokenHandler)
}
