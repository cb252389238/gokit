package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ori/typedef/code"
)

// 响应结构体
type Response struct {
	Code      int    `json:"code"`
	RequestId string `json:"requestId"` //唯一请求ID
	Msg       string `json:"msg"`
	Data      any    `json:"data"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{
		Msg:       "ok",
		RequestId: c.GetString("requestId"),
	}
}

func (r *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func (r *Response) SetRequestId(requestId string) {
	r.RequestId = requestId
}

func (r *Response) SetCode(code code.ErrCode) {
	r.Code = int(code)
	r.Msg = code.ErrorMsg()
}

func (r *Response) SetMsg(msg string) {
	r.Msg = msg
}

func (r *Response) SetData(data any) {
	r.Data = data
}
