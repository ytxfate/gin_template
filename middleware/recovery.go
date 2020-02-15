package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Recovery() gin.HandlerFunc {
	var logger *log.Logger
	logger = log.New(gin.DefaultErrorWriter, "\x1b[31m", log.LstdFlags)
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Printf("%s \t=>  %v\x1b[0m", c.Request.RequestURI, err)
				c.AbortWithStatusJSON(400, gin.H{"msg": "ERROR"})
			}
		}()
		c.Next()
	}
}
