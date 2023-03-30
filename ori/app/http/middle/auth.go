package middle

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		log.Println("请求耗时:", time.Since(t))
	}
}
