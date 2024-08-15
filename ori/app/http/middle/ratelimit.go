package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	coreTool_rate "ori/core/oriTools/rateLimiter"
	"time"
)

var (
	Router = map[string]*RateType{
		"/api/v1/user/collect":     {MaxRequests: 10, Is: true, Window: 1 * time.Minute},
		"/api/v1/index/getCollect": {MaxRequests: 10, Is: true, Window: 1 * time.Minute},
		"api/v1/room/reporting":    {MaxRequests: 5, Is: true, Window: 1 * time.Minute},
	}
)

type RateType struct {
	MaxRequests int           //限制次数
	Is          bool          //是否启用路由限制
	Window      time.Duration //限时时间
}

// 请求速率限制
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		userId := ""
		//如果路由不存在检测里面，直接跳过
		val, ok := Router[path]
		if !ok || !val.Is || userId == "" {
			c.Next()
			return
		}

		userLimiter, ok := coreTool_rate.GetUserRateLimiter(userId)
		if !ok {
			rates := make(map[string]*coreTool_rate.RateLimiter)
			for k, v := range Router {
				item := coreTool_rate.NewRateLimiter(v.MaxRequests, v.Window)
				rates[k] = item
			}
			userLimiter.SetUserRules(userId, rates)
		}
		limiter := userLimiter.GetRateLimiter(path)
		if !limiter.AllowRequest() {
			c.JSON(http.StatusOK, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
