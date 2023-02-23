package douyin_core

const (
	SUCCESS               = 0
	LOGIN_FAILD           = 1
	LOGIN_NO_REQUEST      = 2
	LOGIN_ERROR_PASS_USER = 3
	LOGIN_ERROR_UNION     = 4
	ERR_PARMER            = 5
	UPLODA_ERROR          = 6
	FAILD                 = 7
	REPEAT_ACTION         = 8
	ERR_USERNAME_PASSWORD = 9
)

var (
	statusMessage = map[int]string{
		SUCCESS:               "success",
		LOGIN_FAILD:           "faild",
		LOGIN_NO_REQUEST:      "请输入完整信息",
		LOGIN_ERROR_PASS_USER: "用户名或密码错误",
		LOGIN_ERROR_UNION:     "用户已存在,请重新换一个=￣ω￣=",
		ERR_PARMER:            "参数错误",
		UPLODA_ERROR:          "上传失败",
		FAILD:                 "未知错误",
		REPEAT_ACTION:         "请勿重复操作",
		ERR_USERNAME_PASSWORD: "请检查用户名密码是否正确",
	}
)

func GetStatusMsg(status int) *string {
	res := statusMessage[status]
	return &res
}
