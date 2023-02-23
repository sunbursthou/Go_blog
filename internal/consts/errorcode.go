package consts

// 错误码
const (
	SUCCESS = 200
	ERROR   = 500
	// code 100...用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_USER_NO_RIGHT  = 1004

	// 鉴权相关错误
	ERROR_TOKEN_NOT_EXIST  = 1101
	ERROR_TOKEN_RUNTIME    = 1102
	ERROR_TOKEN_WRONG      = 1103
	ERROR_TOKEN_TYPE_WRONG = 1104
	ERROR_TOKEN_CREATE     = 1105
	ERROR_PERMI_DENIED     = 1106
	FORCE_OFFLINE          = 1107
	LOGOUT                 = 1108
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",
	// 用户
	ERROR_USERNAME_USED:  "用户名已存在",
	ERROR_PASSWORD_WRONG: "用户密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_USER_NO_RIGHT:  "该用户无权限",
	// 权限
	ERROR_TOKEN_NOT_EXIST:  "Token不存在，请重新登录",
	ERROR_TOKEN_RUNTIME:    "Token已过期，请重新登录",
	ERROR_TOKEN_WRONG:      "Token错误， 请重新登录",
	ERROR_TOKEN_TYPE_WRONG: "Token格式错误， 请重新登录",
	ERROR_TOKEN_CREATE:     "Token生成失败",
	ERROR_PERMI_DENIED:     "权限不足",
	FORCE_OFFLINE:          "您已被强制下线",
	LOGOUT:                 "您已退出登录",
}

func GetCodeMsg(code int) string {
	return codeMsg[code]
}
