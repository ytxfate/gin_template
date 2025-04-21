package commresp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type statusCode int
type apiData interface{}

type commRespBody struct {
	Code statusCode `json:"code"`
	Resp apiData    `json:"resp"`
	Msg  string     `json:"msg"`
}

var (
	// ===================  基本 HTTP response code  ===================
	Success        statusCode = 200 // 成功
	ExceptionError statusCode = 400 // 异常错误

	// ===================  其他 HTTP response code  ===================
	ParameterError           statusCode = 1000 // 参数异常错误
	DataCheckError           statusCode = 1001 // 数据比对出错(数据库中不存在此数据 或 此数据已存在于数据库中)
	DataInsertError          statusCode = 1002 // 数据写入数据库出错
	DataUpdateError          statusCode = 1003 // 数据库数据更新出错
	DataDeleteError          statusCode = 1004 // 数据库数据删除出错
	DocumentsAreNotSupported statusCode = 1005 // 不支持的文件上传格式
	FileNotFound             statusCode = 1006 // 文件不存在
	DataCreateError          statusCode = 1007 // 数据生成异常
	// jwt 相关
	JwtCreateError statusCode = 1101 // jwt 生成异常
	JwtParseError  statusCode = 1102 // jwt 解析异常
	// 认证相关
	UserNoLogin     statusCode = 1200 // 用户未登录
	UserLogout      statusCode = 1201 // 用户登出
	UserNoRoles     statusCode = 1202 // 用户没有角色
	UserNoAuthority statusCode = 1203 // 用户没有此接口权限
	// 接口相关
	ApiLimit statusCode = 1300 // 接口限流
)

func CommResp(ctx *gin.Context, code statusCode, resp apiData, msg string) {
	// TODO: 参数校验失败时错误信息放于resp中, 生产环境应该关闭参数错误提示
	if err, ok := resp.(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: err.Error(), Msg: msg})
		return
	}
	ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: resp, Msg: msg})
}
