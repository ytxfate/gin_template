package auth

import "github.com/gin-gonic/gin"

func RouterGroup(e *gin.RouterGroup) {
	authrouter := e.Group("/auth")
	authrouter.GET("/login", loginHandler)
}
