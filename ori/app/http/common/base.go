package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ori/core/oriLog"
	"ori/oerror"
)

// BindJson 绑定 JSON 数据并验证
func BindJson[T any](c *gin.Context, req *T) {
	// 绑定 JSON 数据到结构体
	if err := c.BindJSON(req); err != nil {
		oriLog.Error("BindJson err:%+v", err)
		panic(oerror.Error{
			Code: oerror.ErrorCodeParam,
			Msg:  nil,
		})
	}
	validate := validator.New()
	// 验证结构体
	if err := validate.Struct(req); err != nil {
		oriLog.Error("BindJson validate err:%+v", err)
		panic(oerror.Error{
			Code: oerror.ErrorCodeParam,
			Msg:  nil,
		})
	}
}

// BindBody body绑定参数，返回错误码
func BindBody(c *gin.Context, req any) {
	if err := c.ShouldBind(req); err != nil {
		oriLog.Error("BindBody err:%+v", err)
		panic(oerror.Error{
			Code: oerror.ErrorCodeParam,
			Msg:  nil,
		})
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		oriLog.Error("BindBody validate err:%+v", err)
		panic(oerror.Error{
			Code: oerror.ErrorCodeParam,
			Msg:  nil,
		})
	}
}

// BindQuery query绑定参数，返回错误码
func BindQuery(c *gin.Context, req any) {
	if err := c.ShouldBind(req); err != nil {
		oriLog.Error("BindQuery err:%+v", err)
		panic(oerror.Error{
			Code: oerror.ErrorCodeParam,
			Msg:  nil,
		})
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		oriLog.Error("BindQuery validate err:%+v", err)
		panic(oerror.Error{
			Code: oerror.ErrorCodeParam,
			Msg:  nil,
		})
	}
}
