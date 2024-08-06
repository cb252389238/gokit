package middle

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"ori/core/oriLog"
	"ori/typedef/code"
	"ori/typedef/response"
	"runtime"
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
				oriLog.LogError("异常:%v", msg)
				errorCode := code.ERROR_CODE_UNKNNOWN
				switch v := r.(type) {
				case code.ErrCode:
					errorCode = v
					break
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
