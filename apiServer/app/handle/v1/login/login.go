package login

import (
	"apiServer/app/handle"
	request_login "apiServer/typedef/request/login"

	"apiServer/typedef/response"
	_ "apiServer/typedef/response/login"

	"github.com/gin-gonic/gin"
)

// LoginByPass
//
//	@Summary	手机号密码登录
//	@Schemes
//	@Description	手机号密码登录
//	@Tags			登录相关
//	@Param			req	body	request_login.LoginPassReq	true	"登录参数"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	login.LoginCodeRes
//	@Router			/v1/loginByPass [post]
func LoginByPass(context *gin.Context) {
	req := new(request_login.LoginPassReq)
	handle.BindBody(context, req)
	handle.RepeatSubmitPost(context)
	response.SuccessResponse(context, nil)
}
