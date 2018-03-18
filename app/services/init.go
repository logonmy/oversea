package services

var (
	BackUserService *AdminUserService // 用户服务
	BackAuthService *AuthService
	BackActionLogService *actionLogService
)

func InitServices() {
	BackUserService = &AdminUserService{}
	BackAuthService = &AuthService{}
	BackActionLogService = &actionLogService{}
}
