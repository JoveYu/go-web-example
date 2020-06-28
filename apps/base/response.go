package base

type ErrorCode int

const (
	ERR_LOGIN  ErrorCode = 1000
	ERR_PARAMS ErrorCode = 1100
)

var ERR_MSG = map[ErrorCode]string{
	ERR_LOGIN:  "登录失败",
	ERR_PARAMS: "参数错误",
}

type Response struct {
	ErrorCode ErrorCode `json:"error_code"` // 错误码
	Error     string    `json:"error"`      // 错误描述
	Msg       string    `json:"msg"`        // 提示信息
}
