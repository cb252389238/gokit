package coreTool_rate

import (
	"sync"
	"time"
)

// 内存方式解决请求速率方案 用户和权限 多对多关系
var (
	userLimiters sync.Map // 用户速率限制器
)

type UserRateLimiter struct {
	Rules map[string]*RateLimiter
	sync.RWMutex
}

func NewRateLimiter(maxReq int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		Requests:    0,
		MaxRequests: maxReq,
		Window:      window,
	}
}

type RateLimiter struct {
	Requests    int           //请求总数
	Timestamp   time.Time     //开始计时时间
	MaxRequests int           //计时时间内最大请求次数
	Window      time.Duration //计时单位
	sync.Mutex
}

func GetUserRateLimiter(key string) (*UserRateLimiter, bool) {
	limiterInterface, exists := userLimiters.Load(key)
	if exists {
		return limiterInterface.(*UserRateLimiter), true
	}

	// 如果用户的速率限制器不存在，初始化它
	limiter := &UserRateLimiter{
		Rules: make(map[string]*RateLimiter),
	}
	return limiter, false
}

// 增加使用速率的规则，并且存储
func (u *UserRateLimiter) SetUserRules(key string, item map[string]*RateLimiter) {
	u.Lock()
	defer u.Unlock()
	for k, v := range item {
		u.Rules[k] = v
	}
	userLimiters.LoadOrStore(key, u)
}

// 根据用户和路由获取速率限制器
func (ul *UserRateLimiter) GetRateLimiter(path string) *RateLimiter {
	ul.RLock()
	defer ul.RUnlock()
	limiter, ok := ul.Rules[path]
	if !ok {
		return nil
	}
	return limiter
}

func (rl *RateLimiter) AllowRequest() bool {
	rl.Lock()
	defer rl.Unlock()

	now := time.Now()
	if now.Sub(rl.Timestamp) > rl.Window {
		rl.Timestamp = now
		rl.Requests = 0
	}

	if rl.Requests < rl.MaxRequests {
		rl.Requests++
		return true
	}
	return false
}
