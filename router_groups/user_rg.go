package router_groups

import (
    "gin_template/controller/user_controller"
    "github.com/gin-gonic/gin"
)

func UserRouterGroups(userRg *gin.RouterGroup)  {
    userRg.Any("/login", user_controller.UserLogin)
}
