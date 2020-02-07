package user_controller

import (
    "github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
    var retMap = make(map[string]interface{})
    if c.Request.Method == "GET" {
        id := c.Query("id")
        name := c.DefaultQuery("name", "default_name")
        retMap["id"] = id
        retMap["name"] = name
    } else if c.Request.Method == "POST" {
        id := c.PostForm("id")
        name := c.DefaultPostForm("name", "default_name2")
        //POST和PUT主体参数优先于URL查询字符串值。
        //name := c.Request.FormValue("name")
        //返回POST并放置body参数，URL查询参数被忽略
        //name := c.Request.PostFormValue("name")
        retMap["id2"] = id
        retMap["name2"] = name
    }
    c.JSON(200, retMap)
    return
}