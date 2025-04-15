package user

import "github.com/gin-gonic/gin"

func userHandler(c *gin.Context) {
	c.JSON(200, gin.H{"path": "user"})
}
