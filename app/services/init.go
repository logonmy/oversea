package services

var (
	BackUserService *AdminUserService // 用户服务
	BackAuthService *AuthService
	BackActionLogService *actionLogService
	OzProjectTypeService	*ozProjectTypeService
	OzImmigrantNationService *ozImmigrantNationService
	OzImmigrantProjectService *ozImmigrantProjectService
	CrmCustomerService *crmCustomerService
)

func InitServices() {
	BackUserService = &AdminUserService{}
	BackAuthService = &AuthService{}
	BackActionLogService = &actionLogService{}
	OzProjectTypeService = &ozProjectTypeService{}
	OzImmigrantNationService = &ozImmigrantNationService{}
	OzImmigrantProjectService = &ozImmigrantProjectService{}
	CrmCustomerService = &crmCustomerService{}
}
