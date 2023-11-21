package signalInfo

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "AToken鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "AToken已超时",
	ERROR_AUTH_TOKEN:               "AToken生成失败",
	ERROR_AUTH:                     "AToken错误",
	ERROR_RAUTH_TIMEOUT:            "RToken已超时",
	PAGE_NOT_FOUND:                 "Page not found",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
