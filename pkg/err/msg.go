package err

// 错误码
const (
	SUCCESS = 200
	ERROR   = 500
)

// 错误信息
var errMsg = map[int]string{
	SUCCESS: "success",
	ERROR:   "error",
}

func GetErrMsg(code int) string {
	msg, ok := errMsg[code]
	if ok {
		return msg
	}
	return errMsg[ERROR]
}
