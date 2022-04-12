package response

import (
	"github.com/gin-gonic/gin"
	"pledge-backend/api/common/statecode"
)

type Gin struct {
	Res *gin.Context
}

type Page struct {
	Code  int         `json:"code"`
	Msg   string      `json:"message"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

// ResponsePages
// 响应统一分页格式
func (g *Gin) ResponsePages(c *gin.Context, code int, totalCount int, data interface{}) {
	lang := statecode.LangZh
	langInf, hasLang := c.Get("lang")
	if hasLang {
		lang = langInf.(int)
	}
	rsp := Page{
		Code:  code,
		Msg:   statecode.GetMsg(code, lang),
		Total: totalCount,
		Data:  data,
	}
	g.Res.JSON(200, rsp)
	return
}

// Response  响应统一格式
func (g *Gin) Response(c *gin.Context, code int, data interface{}, httpStatus ...int) {
	lang := statecode.LangEn
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
	return
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
