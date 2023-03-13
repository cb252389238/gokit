package typedef

type Request struct {
	Data any `json:"data"` //数据内容
}

// 响应结构体
type Response struct {
	Code      int    `json:"code"`      //返回码 0成功 http相关返回码1000-1999 tcp相关返回码2000-2999 服务相关错误码3000-
	RequestId string `json:"requestId"` //唯一请求ID
	Msg       string `json:"msg"`
}
