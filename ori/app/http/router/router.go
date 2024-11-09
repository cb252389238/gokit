package router

import (
	"github.com/gin-gonic/gin"
	"ori/app/http/middle"
	"sync"
)

func SetupRouter(reqEntityWg *sync.WaitGroup) *gin.Engine {
	router := gin.New() //实例化gin框架
	router.Use(middle.Cors())
	router.Use(middle.Recover())
	router.Use(middle.RequstLogger())
	group := router.Group("/api")
	SetLUserRouter(group, reqEntityWg)
	return router
}

func BaseRouter(reqEntityWg *sync.WaitGroup, handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		reqEntityWg.Add(1)
		defer reqEntityWg.Done()
		handlerFunc(context)
	}
}
