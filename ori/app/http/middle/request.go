package middle

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"ori/core/oriLog"
	"ori/core/oriSnowflake"
	"time"
)

// CustomResponseWriter 封装 gin ResponseWriter 用于获取回包内容。
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := oriSnowflake.GetSnowId()
		c.Set("requestId", requestId)
		// 记录请求时间
		start := time.Now()
		// 使用自定义 ResponseWriter
		crw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = crw

		// 打印请求信息
		reqBody, _ := c.GetRawData()
		// 请求包体写回。
		if len(reqBody) > 0 {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}
		oriLog.LogInfo("【Request】:| %s | %s | %s | %s", c.ClientIP(), c.Request.Method, c.Request.RequestURI, reqBody)
		// 执行请求处理程序和其他中间件函数
		c.Next()
		// 记录回包内容和处理时间
		end := time.Now()
		latency := end.Sub(start)
		respBody := string(crw.body.Bytes())
		oriLog.LogInfo("【Response】:| %s | %s | %s | %s | (%v)\n", c.ClientIP(), c.Request.Method, c.Request.RequestURI, respBody, latency)
	}
}
