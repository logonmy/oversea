package services

var (
	BackUserService       *AdminUserService       // 用户服务
	BackAuthService           *AuthService
)
func InitServices()  {
	BackUserService = &AdminUserService{}
	BackAuthService = &AuthService{}
}