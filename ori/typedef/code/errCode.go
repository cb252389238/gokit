package code

// 1-999 成功状态码
// 10000-19999 参数校验错误
// 20000-29999 数据库错误
// 30000-39999 业务执行错误
// 40000-49999 异常错误

type ErrCode int

func (errCode ErrCode) ErrorMsg() string {
	if msg, ok := ErrMsg[errCode]; ok {
		return msg
	}
	return ErrMsg[ERROR_CODE_UNKNNOWN]
}

var (
	SUCCESS_CODE ErrCode = 0 //成功
)

var (
	ERROR_CODE_PARAM ErrCode = 10000

	ERROR_CODE_UNKNNOWN ErrCode = 99999
)