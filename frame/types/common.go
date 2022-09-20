package types

type Request struct {
	Data ReqData `json:"data"` //数据内容
}

type ReqData struct {
	Msg string `json:"text" binding:"required,max=1024"` //检测文本
}

// 响应结构体
type Response struct {
	Code      int    `json:"code"`      //返回码 0成功 http相关返回码1000-1999 tcp相关返回码2000-2999 服务相关错误码3000-
	Message   string `json:"message"`   //返回描述
	RequestId string `json:"requestId"` //唯一请求ID
	Msg       string `json:"msg"`
}
