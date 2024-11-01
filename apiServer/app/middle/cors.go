package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域请求的头部
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, X-Request-Id, X-Forwarded-For, X-Real-IP, appVersion, channel, machineCode, platform")

		// 允许发送Cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 允许前端直接读取
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")

		// 响应OPTIONS请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next() // 继续处理请求
		}
	}
}
