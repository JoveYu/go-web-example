package base

type ErrorCode int

const (
	OK         ErrorCode = 0
	ERR_LOGIN  ErrorCode = 1000
	ERR_PARAMS ErrorCode = 1100
	ERR_DB     ErrorCode = 1200
	ERR_USER   ErrorCode = 1300
)

var ERR_MSG = map[ErrorCode]string{
	OK:         "请求成功",
	ERR_LOGIN:  "用户未登录",
	ERR_PARAMS: "参数错误",
	ERR_DB:     "数据库错误",
	ERR_USER:   "用户错误",
}

type Response struct {
	Code    ErrorCode `json:"code"`    // 错误码
	Error   string    `json:"error"`   // 错误描述
	Message string    `json:"message"` // 提示信息
}

func NewResponse(code ErrorCode, s ...string) Response {
	r := Response{
		Code: code,
	}
	if v, ok := ERR_MSG[code]; ok {
		r.Message = v
	} else {
		r.Message = "未知错误"
	}

	if len(s) > 0 {
		r.Error = s[0]
	}

	if len(s) > 1 {
		r.Message = s[1]
	}

	return r
}
