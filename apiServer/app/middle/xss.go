package middle

import (
	"html"
	"strings"

	"github.com/gin-gonic/gin"
)

// 过滤用户输入，防止XSS攻击
func FilterXSS(input string) string {
	// 将特殊字符转义成HTML实体
	input = html.EscapeString(input)

	// 移除一些危险的HTML标签和属性
	input = strings.ReplaceAll(input, "<script", "&lt;script")
	input = strings.ReplaceAll(input, "</script", "&lt;/script")
	input = strings.ReplaceAll(input, "<iframe", "&lt;iframe")
	input = strings.ReplaceAll(input, "</iframe", "&lt;/iframe")
	input = strings.ReplaceAll(input, "javascript:", "")

	return input
}

// XSS拦截中间件
func XSSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取输入参数
		for _, v := range c.Request.URL.Query() {
			v[0] = FilterXSS(v[0])
			c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, v[0], FilterXSS(v[0]))
		}
		for k, v := range c.Request.PostForm {
			v[0] = FilterXSS(v[0])
			c.Request.PostForm.Set(k, v[0])
		}
		for k, v := range c.Request.PostForm {
			c.Request.Form.Set(k, FilterXSS(v[0]))
		}

		c.Next()
	}
}
