package response

import (
	error2 "apiServer/i18n/error"
	"net/http"
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

func SuccessResponse(c *gin.Context, data any) {
	res := &Response{
		Code:      error2.SuccessCode,
		RequestId: c.GetString("requestId"),
		c:         c,
		//Msg:       code2.SuccessCode.ErrorMsg(),
		Msg: ginI18n.MustGetMessage(
			c,
			&i18n.LocalizeConfig{
				MessageID:    strconv.Itoa(error2.SuccessCode),
				TemplateData: nil,
			}),
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func FailResponse(c *gin.Context, anyCode any) {
	code := 0
	msg := ""
	switch anyCode.(type) {
	case error2.I18nError:
		realCode := anyCode.(error2.I18nError)
		code = realCode.Code
		msg = ginI18n.MustGetMessage(
			c,
			&i18n.LocalizeConfig{
				MessageID:    strconv.Itoa(code),
				TemplateData: realCode.Msg,
			})
	default:
		realCode := anyCode.(error2.ErrCode)
		code = realCode
		msg = ginI18n.MustGetMessage(
			c,
			&i18n.LocalizeConfig{
				MessageID: strconv.Itoa(code),
			})
	}
	res := &Response{
		Code:      code,
		RequestId: c.GetString("requestId"),
		c:         c,
		Msg:       msg,
	}
	c.JSON(http.StatusOK, res)
}
