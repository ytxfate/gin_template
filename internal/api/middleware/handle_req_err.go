package middleware

import (
	commresp "gin_template/internal/comm_resp"

	"github.com/gin-gonic/gin"
)

func HandleRequestError(c *gin.Context) {
	commresp.CommResp(c, commresp.ExceptionError, nil, "NOT FOUND")
	c.Abort()
}
