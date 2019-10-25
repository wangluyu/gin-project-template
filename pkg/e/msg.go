package e

// 错误码
const (
	Success               = 200
	Error                 = 500
	ErrorAuth             = 1001
	ErrorAuthToken        = 1002
	ErrorAuthTokenExpired = 1003
	ErrorAuthParams       = 1004
)

// 错误信息
var errMsg = map[int]string{
	Success:               "success",
	Error:                 "error",
	ErrorAuth:             "没有权限",
	ErrorAuthToken:        "无效Token",
	ErrorAuthTokenExpired: "Token已过期",
	ErrorAuthParams:       "参数错误",
}

func GetErrMsg(code int) string {
	msg, ok := errMsg[code]
	if ok {
		return msg
	}
	return errMsg[Error]
}
