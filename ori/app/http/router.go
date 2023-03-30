package http

import (
	"github.com/gin-gonic/gin"
	"ori/app/http/handle"
	"ori/app/http/middle"
	"ori/core/oriEngine"
)

func SetupRouter(oriEngine *oriEngine.OriEngine) *gin.Engine {
	router := gin.New() //实例化gin框架
	router.Use(middle.Auth())
	router.Any("/test1", handle.Test1(oriEngine)) //any 代表各种请求都可以进来

	router.GET("/test2", handle.Test2(oriEngine)) //get请求

	router.POST("/test3", handle.Test3(oriEngine)) //post请求

	//路由分组
	ad := router.Group("/v1")
	{
		ad.Any("/test4", handle.Test4(oriEngine))
		ad.Any("/test5", handle.Test5(oriEngine))
	}

	//例子
	router.GET("/userinfo", handle.UserInfo(oriEngine))
	return router
}
