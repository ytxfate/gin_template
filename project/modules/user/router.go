package user

import "github.com/gin-gonic/gin"

func RouterGroup(e *gin.RouterGroup) {
	authrouter := e.Group("/user")
	authrouter.GET("/", userHandler)
	authrouter.GET("/2", user2Handler)
}
