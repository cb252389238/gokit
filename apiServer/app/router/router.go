package router

import (
	"apiServer/app/middle"
	"apiServer/i18n"
	"apiServer/internal/engine"
	"github.com/gin-gonic/gin"
	"sync"
)

func SetupRouter(oriEngine *engine.OriEngine, wg *sync.WaitGroup) *gin.Engine {
	router := gin.New() //实例化gin框架
	router.Use(middle.Cors())
	router.Use(middle.Recover())
	router.Use(middle.Logger())
	router.Use(i18n.I18nHandler())
	group := router.Group("/api")
	SetLoginRouter(oriEngine, group, wg)
	return router
}
func BaseRouter(oriEngine *engine.OriEngine, wg *sync.WaitGroup, handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		wg.Add(1)
		defer wg.Done()
		handlerFunc(context)
	}
}
