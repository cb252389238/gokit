package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ori/typedef/code"
)

// 响应结构体
type Response struct {
	Code      int          `json:"code"`
	RequestId string       `json:"requestId"` //唯一请求ID
	Msg       string       `json:"msg"`
	Data      any          `json:"data"`
	c         *gin.Context `json:"-"`
}

func SuccessResponse(c *gin.Context, data any) {
	res := &Response{
		Code:      int(code.SUCCESS_CODE),
		RequestId: c.GetString("requestId"),
		c:         c,
		Msg:       code.SUCCESS_CODE.ErrorMsg(),
		Data:      data,
	}
	c.JSON(http.StatusOK, res)
}

func FailResponse(c *gin.Context, code code.ErrCode) {
	res := &Response{
		Code:      int(code),
		RequestId: c.GetString("requestId"),
		c:         c,
		Msg:       code.ErrorMsg(),
	}
	c.JSON(http.StatusOK, res)
}
