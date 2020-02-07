package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        var statusColor, methodColor, resetColor string
        if param.IsOutputColor() {
           statusColor = param.StatusCodeColor()
           methodColor = param.MethodColor()
           resetColor = param.ResetColor()
        }
        // 你的自定义格式
        return fmt.Sprintf("%s - %15s \"|%s %s %s %s %s |%s %3d %s| %s %s\"\n",
            param.TimeStamp.Format("2006-01-02 15:04:05"),
            param.ClientIP,
            methodColor, param.Method, resetColor,
            param.Path,
            param.Request.Proto,
            statusColor, param.StatusCode, resetColor,
            param.Latency,
            param.ErrorMessage,
        )
    })
}
