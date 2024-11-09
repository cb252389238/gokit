package handle_v1

import (
	"github.com/gin-gonic/gin"
	"ori/app/http/common"
	"ori/internal/logic"
	"ori/oresponse"
	"ori/typedef/request/user"
	user2 "ori/typedef/response/user"
)

func GetUserInfo(c *gin.Context) {
	req := new(user.UserInfoReq)
	common.BindQuery(c, req)
	logicSer := new(logic.User)
	resp := user2.UserInfo{}
	err := logicSer.GetUserInfo(req, &resp)
	oresponse.JsonResponse(c, err, resp)
}
