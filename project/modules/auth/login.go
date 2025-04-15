package auth

import "github.com/gin-gonic/gin"

func loginHandler(c *gin.Context) {
	c.JSON(200, gin.H{"username": "xxx", "password": "yyy"})
}
