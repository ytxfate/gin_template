package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequestError(c *gin.Context) {
	var logger *log.Logger
	logger = log.New(gin.DefaultErrorWriter, "\x1b[31m", log.LstdFlags)
	logger.Printf("%s \t=>  not found\x1b[0m", c.Request.RequestURI)
	c.AbortWithStatusJSON(400, gin.H{"msg": "NOT FOUND"})
}
