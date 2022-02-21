package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pledge-backend/api/common/statecode"
	"pledge-backend/log"
)

//--------------------------------------------------------------------------------------------
type Gin struct {
	Res *gin.Context
}

//--------------------------------------------------------------------------------------------
type ResponsePage struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

//--------------------------------------------------------------------------------------------
// 响应统一分页格式
func (g *Gin) ResponsePages(c *gin.Context, code int, totalCount int, data interface{}) {
	lang := statecode.LANG_ZH
	langInf, hasLang := c.Get("lang")
	if hasLang {
		lang = langInf.(int)
	}
	rsp := ResponsePage{
		Code: code,
		Msg:  statecode.GetMsg(code, lang),
		Data: data,
	}
	g.Res.JSON(200, rsp)
	return
}

//--------------------------------------------------------------------------------------------
// 响应统一格式
func (g *Gin) Response(c *gin.Context, code int, data interface{}, httpStatus ...int) {
	lang := statecode.LANG_EN
	langInf, hasLang := c.Get("lang")
	if hasLang {
		lang = langInf.(int)
	}
	rsp := Response{
		Code: code,
		Msg:  statecode.GetMsg(code, lang),
		Data: data,
	}
	HttpStatus := 200
	if len(httpStatus) > 0 {
		HttpStatus = httpStatus[0]
	}
	g.Res.JSON(HttpStatus, rsp)
	log.Logger.Info(fmt.Sprintf("----- response ----- %+v", rsp))
	return
}

//--------------------------------------------------------------------------------------------
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

//--------------------------------------------------------------------------------------------
type RspPhoneMsgVerifyCode struct {
	//保留为空,为了开发阶段暂时增加如下两个字段作为测试联调
	TelephoneNO     string `json:"telephoneNO"`
	PhoneVerifyCode string `json:"phoneVerifyCode"`
}

//--------------------------------------------------------------------------------------------
type RspEmailVerifyCode struct {
	//保留为空
}

//--------------------------------------------------------------------------------------------
type RspOpenGoogleVerifyCode struct {
	TokenKey string `json:"tokenKey"`
	Qrcode   string `json:"qrcode"`
}

//--------------------------------------------------------------------------------------------
type RspConfirmOpenGoogleVerify struct {
	//保留为空
}

//--------------------------------------------------------------------------------------------
type RspCloseGoogleVerify struct {
	//保留为空
}

//--------------------------------------------------------------------------------------------
