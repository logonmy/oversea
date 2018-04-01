package stdout


const (
	HttpStatusSuccess  = 200
	HttpStatusFail     = 200
	HttpStatusNotFound = 404
	HttpNotAuthorization = 401
)

const (
	Success = 0
	// 数据库错误 或更新修改等操作失败
	DBError = 1000
	// 参数错误
	ParamsError = 2000
	NotAuthorizationError = 2001
)

// 后台errMsg
const (
	// MsgSuccess ...
	MsgSuccess = "success"
	// MsgUnknown ...
	MsgUnknown = "unknown error"
	UserOrderPasswordError = "账户或密码错误"
	SystemError = "系统错误"
	UserIsDisenbled = "账号已经被禁用"
	FieldsLengthMustMoreThanOne = "更新字段不能为空"
	CaptchaError = "验证码错误"
	UserIsExists = "用户已存在"
	UsernameOrPasswdEmptyError = "用户名或密码不能为空"
	UsernameEmptyError = "用户名不能为空"
	RepeatPasswordError = "两次输入密码不一致"
	AddAdminFailError = "管理员添加失败"
	EmailAddressError = "邮箱地址格式错误"
	MobilePhoneError = "手机号码格式错误"
	UpdateError = "更新失败"
	NotAuthorizationErrorMsg = "你还没未登录，请先登陆"
)



