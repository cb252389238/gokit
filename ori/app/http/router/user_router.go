package router

import (
	"github.com/gin-gonic/gin"
	"ori/app/http/handle/v1"
	"ori/app/http/middle"
	"sync"
)

func SetLUserRouter(router *gin.RouterGroup, reqEntityWg *sync.WaitGroup) {
	group := router.Group("/v1/user")
	group.Use(middle.Auth())
	{
		group.POST("/userInfo", BaseRouter(reqEntityWg, handle_v1.GetUserInfo)) // 查询用户主页信息
	}
}
