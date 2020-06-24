package base

type ErrorCode int

const (
	ERR_LOGIN  ErrorCode = 1000
	ERR_PARAMS ErrorCode = 1100
)

type Response struct {
	ErrorCode ErrorCode `json:"error_code"` // 错误码
	Error     string    `json:"error"`      // 错误描述
	Msg       string    `json:"msg"`        // 提示信息
}
