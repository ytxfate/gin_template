package middleware

import (
	"gin_template/internal/pkg/commresp"

	"github.com/gin-gonic/gin"
)

func HandleRequestError(c *gin.Context) {
	commresp.CommResp(c, commresp.ExceptionError, nil, "NOT FOUND")
	c.Abort()
}
