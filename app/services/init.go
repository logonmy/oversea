package services

var (
	SysUserService               *AdminUserService // 用户服务
	SysAuthService               *AuthService
	SysActionLogService          *actionLogService
	OzProjectTypeService         *ozProjectTypeService
	OzImmigrantNationService     *ozImmigrantNationService
	OzImmigrantProjectService    *ozImmigrantProjectService
	CrmCustomerService           *crmCustomerService
	SysConfigService             *sysConfigService
	CmsArticleService            *cmsArticleService
	CmsCategoryArticleRelService *cmsCategoryArticleRelService
	CmsCategoryService           *cmsCategoryService
	CmsGuestbookService          *cmsGuestbookService
	CmsUsersService              *cmsUsersService
	CrmLinkmanService            *crmLinkmanService
)

func InitServices() {
	SysUserService = &AdminUserService{}
	SysAuthService = &AuthService{}
	SysActionLogService = &actionLogService{}
	OzProjectTypeService = &ozProjectTypeService{}
	OzImmigrantNationService = &ozImmigrantNationService{}
	OzImmigrantProjectService = &ozImmigrantProjectService{}
	CrmCustomerService = &crmCustomerService{}
	SysConfigService = &sysConfigService{}
	CmsArticleService = &cmsArticleService{}
	CmsCategoryArticleRelService = &cmsCategoryArticleRelService{}
	CmsCategoryService = &cmsCategoryService{}
	CmsGuestbookService = &cmsGuestbookService{}
	CmsUsersService = &cmsUsersService{}
	CrmLinkmanService = &crmLinkmanService{}
}
