package middle

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"ori/core/oriLog"
	"ori/oerror"
	"ori/oresponse"
	"runtime"
)

func printStackTrace(err any) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v", err)
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d\n", file, line)
	}
	return buf.String()
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				msg := printStackTrace(r)
				oriLog.Error("异常:%v", msg)
				var errorCode any
				switch v := r.(type) {
				case oerror.Error:
					errorCode = v
				case oerror.ErrCode: //解决除了api多语言接口返回信息错误被重置为未知错误
					errorCode = v
				default:
					errorCode = oerror.ErrorCodeUnknown
				}
				oresponse.JsonResponse(c, errorCode, nil)
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		c.Next()
	}
}
