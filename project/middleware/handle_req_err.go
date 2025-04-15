package middleware

import (
	"github.com/gin-gonic/gin"
)

func HandleRequestError(c *gin.Context) {
	c.AbortWithStatusJSON(400, gin.H{"msg": "NOT FOUND"})
}
