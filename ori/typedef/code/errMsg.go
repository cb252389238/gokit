package code

// 1-999 成功状态码
// 10000-19999 参数校验错误
// 20000-29999 数据库错误
// 30000-39999 业务执行错误
// 40000-49999 异常错误

var ErrMsg = map[ErrCode]string{
	SUCCESS_CODE:        "ok",
	ERROR_CODE_PARAM:    "参数错误",
	ERROR_CODE_UNKNNOWN: "未知错误",
}
