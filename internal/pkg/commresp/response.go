package commresp

import (
	"gin_template/internal/backend/webconfig"
	"gin_template/pkg/deployenv"
	"gin_template/pkg/logger"
	"gin_template/pkg/validatorzhtrans"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type respData interface{}

type commRespBody struct {
	Code StatusCode `json:"code"`
	Resp respData   `json:"resp"`
	Msg  string     `json:"msg"`
}

func CommResp(ctx *gin.Context, code StatusCode, resp respData, msg string) {
	// 参数校验失败时错误信息放于resp中, 生产环境应该关闭参数错误提示
	switch err := resp.(type) {
	case validator.ValidationErrors:
		e := err.Translate(validatorzhtrans.Trans)
		logger.Infof("%#v", e)
		if webconfig.Cfg.Env != deployenv.PROD {
			ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: e, Msg: msg})
		} else {
			ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: nil, Msg: msg})
		}
		return
	case error:
		logger.Info(err.Error())
		if webconfig.Cfg.Env != deployenv.PROD {
			ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: err.Error(), Msg: msg})
		} else {
			ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: nil, Msg: "请求异常"})
		}
		return
	}
	ctx.JSON(http.StatusOK, commRespBody{Code: code, Resp: resp, Msg: msg})
}
