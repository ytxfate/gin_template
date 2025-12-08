package commresp

import (
	webconfig "gin_template/internal/web_config"
	zhtrans "gin_template/internal/zh_trans"
	"gin_template/pkg/config"
	"gin_template/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StatusCode int
type apiData interface{}

type commRespBody struct {
	Code StatusCode `json:"code"`
	Resp apiData    `json:"resp"`
	Msg  string     `json:"msg"`
}

var (
	// ===================  基本 HTTP response code  ===================
	Success        StatusCode = 200 // 成功
	ExceptionError StatusCode = 400 // 异常错误

	// ===================  其他 HTTP response code  ===================
	ParameterError           StatusCode = 1000 // 参数异常错误
	DataCheckError           StatusCode = 1001 // 数据比对出错(数据库中不存在此数据 或 此数据已存在于数据库中)
	DataInsertError          StatusCode = 1002 // 数据写入数据库出错
	DataUpdateError          StatusCode = 1003 // 数据库数据更新出错
	DataDeleteError          StatusCode = 1004 // 数据库数据删除出错
	DocumentsAreNotSupported StatusCode = 1005 // 不支持的文件上传格式
	FileNotFound             StatusCode = 1006 // 文件不存在
	DataCreateError          StatusCode = 1007 // 数据生成异常
	// jwt 相关
	JwtCreateError StatusCode = 1101 // jwt 生成异常
	JwtParseError  StatusCode = 1102 // jwt 解析异常
	// 认证相关
	UserNoLogin     StatusCode = 1200 // 用户未登录
	UserLogout      StatusCode = 1201 // 用户登出
	UserNoRoles     StatusCode = 1202 // 用户没有角色
	UserNoAuthority StatusCode = 1203 // 用户没有此接口权限
	// 接口相关
	ApiLimit StatusCode = 1300 // 接口限流
)

func CommResp(ctx *gin.Context, code StatusCode, resp apiData, msg string) {
	// 参数校验失败时错误信息放于resp中, 生产环境应该关闭参数错误提示
	switch err := resp.(type) {
	case validator.ValidationErrors:
		e := err.Translate(zhtrans.Trans)
		logger.Infof("%#v", e)
		if webconfig.Cfg.Env != config.PROD {
			ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: e, Msg: msg})
		} else {
			ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: nil, Msg: msg})
		}
		return
	case error:
		logger.Info(err.Error())
		if webconfig.Cfg.Env != config.PROD {
			ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: err.Error(), Msg: msg})
		} else {
			ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: nil, Msg: "请求异常"})
		}
		return
	}
	ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: resp, Msg: msg})
}
