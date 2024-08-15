package coreTool_rate

import (
	"sync"
	"testing"
	"time"
)

// 基准测试：测试速率限制器在高负载下的表现
func BenchmarkRateLimiter(b *testing.B) {
	// 设置速率限制器的参数
	maxRequests := 10
	window := time.Second
	limiter := NewRateLimiter(maxRequests, window)

	// 记录开始时间
	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 模拟并发请求
			limiter.AllowRequest()
		}()
	}
	wg.Wait()

	// 输出性能信息
	elapsed := time.Since(start)
	b.Logf("Elapsed time: %s", elapsed)
}

func TestRateLimiter(t *testing.T) {

	// 创建用户速率限制器
	userID := "user1"
	userLimiter, exists := GetUserRateLimiter(userID)
	if !exists {
		// 定义速率限制规则
		limiter1 := NewRateLimiter(10, 10*time.Second)
		limiter2 := NewRateLimiter(10, 10*time.Second)

		// 设置用户速率限制规则
		rules := map[string]*RateLimiter{
			"/api/test":    limiter1,
			"/api/another": limiter2,
		}
		userLimiter.SetUserRules(userID, rules)
	}

	// 获取并测试速率限制器
	limiter := userLimiter.GetRateLimiter("/api/test")
	if limiter == nil {
		t.Fatalf("Expected rate limiter for path /api/test but got none")
	}

	// 测试 AllowRequest 函数
	for i := 0; i < 10; i++ {
		if !limiter.AllowRequest() {
			t.Errorf("Request %d should be allowed but was denied", i)
		}
	}

	// 应该超出限制
	if limiter.AllowRequest() {
		t.Error("Request after limit should be denied")
	}

	// 测试时间窗口
	time.Sleep(15 * time.Second) // 等待时间窗口刷新

	// 应该允许新请求
	if !limiter.AllowRequest() {
		t.Error("Request after window reset should be allowed")
	}

	// 测试不同路径的速率限制器
	limiter = userLimiter.GetRateLimiter("/api/another")
	if limiter == nil {
		t.Fatalf("Expected rate limiter for path /api/another but got none")
	}

	for i := 0; i < 10; i++ {
		if !limiter.AllowRequest() {
			t.Errorf("Request %d for /api/another should be allowed but was denied", i)
		}
	}

	// 应该超出限制
	if limiter.AllowRequest() {
		t.Error("Request after limit for /api/another should be denied")
	}
}
