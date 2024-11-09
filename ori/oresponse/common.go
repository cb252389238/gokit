package oresponse

import (
	"net/http"
	"ori/core/oriLog"
	"ori/oerror"
	"strconv"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// 响应结构体
type Response struct {
	Code      int          `json:"code"`
	RequestId string       `json:"requestId"` //唯一请求ID
	Msg       string       `json:"msg"`
	Data      any          `json:"data,omitempty"`
	c         *gin.Context `json:"-"`
}

func JsonResponse(c *gin.Context, err any, data any) {
	if err != nil {
		failResponse(c, err)
		return
	}
	successResponse(c, data)
}

func successResponse(c *gin.Context, data any) {
	res := &Response{
		Code:      oerror.SuccessCode,
		RequestId: c.GetString("requestId"),
		c:         c,
		Msg: ginI18n.MustGetMessage(
			c,
			&i18n.LocalizeConfig{
				MessageID:    strconv.Itoa(oerror.SuccessCode),
				TemplateData: nil,
			}),
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func failResponse(c *gin.Context, err any) {
	code := 0
	msg := ""
	switch err.(type) {
	case oerror.Error: //自定义错误
		realError := err.(oerror.Error)
		code = realError.Code
		msg = ginI18n.MustGetMessage(
			c,
			&i18n.LocalizeConfig{
				MessageID:    strconv.Itoa(code),
				TemplateData: realError.Msg,
			})
	case oerror.ErrCode:
		code = err.(oerror.ErrCode)
		msg = ginI18n.MustGetMessage(
			c,
			&i18n.LocalizeConfig{
				MessageID:    strconv.Itoa(code),
				TemplateData: nil,
			})
	default: //默认错误
		realError := err.(error)
		msg = realError.Error()
		oriLog.Error("err:%+v", realError)
		code = oerror.ErrorCodeUnknown
		msg = ginI18n.MustGetMessage(
			c,
			&i18n.LocalizeConfig{
				MessageID:    strconv.Itoa(code),
				TemplateData: nil,
			},
		)
	}
	res := &Response{
		Code:      code,
		RequestId: c.GetString("requestId"),
		c:         c,
		Msg:       msg,
	}
	c.JSON(http.StatusOK, res)
}
