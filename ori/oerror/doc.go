package oerror

// 多语言错误信息处理
type I18nError struct {
	Code int
	Msg  map[string]interface{}
}
