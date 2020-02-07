package register_rgs

import (
    "gin_template/router_groups"
    "github.com/gin-gonic/gin"
)

func RegisterRouterGroups(rg *gin.RouterGroup) {
    // 注册 user 路由
    userRg := rg.Group("user")
    router_groups.UserRouterGroups(userRg)
}
