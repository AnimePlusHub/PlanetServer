package signalInfo

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "AToken鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "AToken已超时",
	ERROR_AUTH_TOKEN:               "AToken生成失败",
	ERROR_AUTH:                     "AToken错误",
	ERROR_RAUTH_TIMEOUT:            "RToken已超时",
	MISS_TOKEN:                     "缺少access token",

	ERROR_USER_NOT_FOUND: "未找到用户",
	ERROR_WRONG_PWD:      "密码错误",

	REDIS_SUCCESS:       "Redis操作成功",
	REDIS_FAIL:          "Redis操作失败",
	REDIS_KEY_NOT_EXIST: "Redis Key值不存在",

	EMAIL_FOMAT_FALSE: "邮箱格式错误",
	VALIDATE_PASS:     "验证码正确",
	VALIDATE_WRONG:    "验证码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
