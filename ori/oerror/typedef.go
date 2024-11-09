package oerror

// 创建错误
func NewError(code int, msg ...map[string]any) Error {
	data := make(map[string]any)
	if len(msg) > 0 {
		data = msg[0]
	}
	return Error{
		Code: code,
		Msg:  data,
	}
}

// 错误类型
type Error struct {
	Code int
	Msg  map[string]any
}

func (e Error) Error() string {
	return ""
}
