package oerror

type ErrCode = int

var (
	SuccessCode ErrCode = 0 //成功
)

var (
	ErrorCodeOperationFrequent ErrCode = 99997 //操作频繁
	ErrorCodeSystemBusy        ErrCode = 99998 //系统繁忙
	ErrorCodeUnknown           ErrCode = 99999 //未知错误
)
