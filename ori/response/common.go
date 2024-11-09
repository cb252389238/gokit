package response

import (
	"net/http"
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

func SuccessResponse(c *gin.Context, data any) {
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

func FailResponse(c *gin.Context, anyCode any) {
	code := 0
	msg := ""
	switch anyCode.(type) {
	case oerror.I18nError:
		realCode := anyCode.(oerror.I18nError)
		code = realCode.Code
		msg = ginI18n.MustGetMessage(
			c,
			&i18n.LocalizeConfig{
				MessageID:    strconv.Itoa(code),
				TemplateData: realCode.Msg,
			})
	default:
		realCode := anyCode.(oerror.ErrCode)
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
