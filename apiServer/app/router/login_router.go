package router

import (
	v1_login "apiServer/app/handle/v1/login"
	"apiServer/internal/engine"
	"github.com/gin-gonic/gin"
	"sync"
)

func SetLoginRouter(oriEngine *engine.OriEngine, router *gin.RouterGroup, wg *sync.WaitGroup) {
	group := router.Group("/v1")
	{
		group.POST("/loginByPass", BaseRouter(oriEngine, wg, v1_login.LoginByPass))
	}
}
