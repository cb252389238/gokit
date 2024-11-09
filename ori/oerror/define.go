package oerror

type ErrCode = int

var (
	SuccessCode ErrCode = 0 //成功
)

var (
	ErrorCodeParam             ErrCode = 99996 //参数错误
	ErrorCodeOperationFrequent ErrCode = 99997 //操作频繁
	ErrorCodeSystemBusy        ErrCode = 99998 //系统繁忙
	ErrorCodeUnknown           ErrCode = 99999 //未知错误
)
