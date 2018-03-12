package constant

const (
	HttpStatusSuccess  = 200
	HttpStatusFail     = 200
	HttpStatusNotFound = 404
)
const (
	NoLoginError = -1
	// Success 没有错误
	Success = 0
	// TestError 测试错误
	TestError = 100
	//参数错误
	ParamsError = 201
	// NetError 网络错误
	NetError = 1000
	// DBError 数据库错误
	DBError = 2000
	// CaptchaError 验证码错误
	CaptchaError = 9000
	// user is exists
	UserIsExistsError = 9001
	UserOrPasswordError = 9002
	UserIsNoExistsError = 9003
	ChangePasswordError = 9004
	SamePasswordError = 9005
	ChangeInfoError = 9006
)
const (
	// MsgSuccess ...
	MsgSuccess = "success"
	// MsgUnknown ...
	MsgUnknown = "unknown error"
)

func ErrCode2Msg(code int) string {
	switch code {
	case NoLoginError:
		return "未登录"
	case TestError:
		return "测试错误"
	case NetError:
		return "网络错误"
	case DBError:
		return "数据库错误"
	case CaptchaError:
		return "验证码错误"
	case UserIsExistsError:
		return "账号已存在"
	case UserOrPasswordError:
		return "账号或密码错误"
	case UserIsNoExistsError:
		return "账号不存在"
	case ChangePasswordError:
		return "修改密码失败"
	case SamePasswordError:
		return "新旧密码不能一致"
	case ParamsError:
		return "参数错误"
	case ChangeInfoError:
		return "修改个人信息失败"

	}
	return MsgUnknown
}
