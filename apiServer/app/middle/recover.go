package middle

import (
	"apiServer/core/coreLog"
	error2 "apiServer/i18n/error"
	"apiServer/typedef/response"
	"bytes"
	"fmt"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func printStackTrace(err any) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%dv \n", file, line)
	}
	return buf.String()
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				msg := printStackTrace(r)
				coreLog.LogError("异常:%v", msg)
				var errorCode any
				switch v := r.(type) {
				case error2.I18nError:
					errorCode = v
				case error2.ErrCode: //解决除了api多语言接口返回信息错误被重置为未知错误
					errorCode = v
				default:
					errorCode = error2.ErrorCodeUnknown
				}
				response.FailResponse(c, errorCode)
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		c.Next()
	}
}

var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
