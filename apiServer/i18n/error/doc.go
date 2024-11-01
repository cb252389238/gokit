package i18n_err

// 1-999 成功状态码
// 10000-19999 参数校验错误
// 20000-29999 数据库错误
// 30000-39999 业务执行错误
// 40000-49999 异常错误

// 多语言错误信息处理
type I18nError struct {
	Code int
	Msg  map[string]interface{}
}

// func test() {
// 	panic(&I18nError{Code: 1112, Msg: nil})
// 	panic(&I18nError{Code: 1111, Msg: map[string]interface{}{
// 		"times": 3,
// 		"expr":  10 * time.Second,
// 	}})

// }
