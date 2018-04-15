package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"oversea/app/controllers/backend"
)

func adminRouters()  {

	// 验证码路由
	beego.Handler("/captcha/*.png", captcha.Server(130, 40))

	beego.Router("/api/v1/user/getInfo", &backend.AdminUserController{}, "get:GetInfo")
	beego.Router("/api/v1/user/list", &backend.AdminUserController{}, "get,post:List")
	beego.Router("/api/v1/login", &backend.MainController{}, "post:Login")
	beego.Router("/api/v1/logout", &backend.MainController{}, "get:Logout")
	beego.Router("/api/v1/user/add", &backend.AdminUserController{}, "post:Add")
	beego.Router("/api/v1/user/edit", &backend.AdminUserController{}, "post:Edit")
	beego.Router("/api/v1/user/profile", &backend.AdminUserController{}, "get:GetMyInfo")
	beego.Router("/api/v1/user/changePwd", &backend.AdminUserController{}, "post:ChangePassword")
    // 客户
	beego.Router("/api/v1/customer/list", &backend.CustomerController{}, "post:List")
	beego.Router("/api/v1/customer/getInfo", &backend.CustomerController{}, "get:GetInfo")
	beego.Router("/api/v1/customer/add", &backend.CustomerController{}, "post:AddCrmCustomer")
	beego.Router("/api/v1/customer/update", &backend.CustomerController{}, "post:UpdateCrmCustomerById")
	beego.Router("/api/v1/customer/delete", &backend.CustomerController{}, "get:DeleteCrmCustomer")
	beego.Router("/api/v1/customer/export", &backend.CustomerController{}, "get:Export")
    // 联系人
	beego.Router("/api/v1/linkman/getInfo", &backend.CrmLinkmanController{}, "get:GetCrmLinkmanById")
	beego.Router("/api/v1/linkman/all/list", &backend.CrmLinkmanController{}, "post:GetAllCrmLinkmanList")
	beego.Router("/api/v1/linkman/me/list", &backend.CrmLinkmanController{}, "post:GetMyCrmLinkmanList")
	beego.Router("/api/v1/linkman/update", &backend.CrmLinkmanController{}, "post:UpdateCrmLinkmanById")
	beego.Router("/api/v1/linkman/add", &backend.CrmLinkmanController{}, "post:AddCrmLinkman")
	beego.Router("/api/v1/linkman/delete", &backend.CrmLinkmanController{}, "get:DeleteCrmLinkman")

	beego.Router("/api/v1/contact/add", &backend.ContactController{}, "post:AddContactNote")
	beego.Router("/api/v1/contact/list", &backend.ContactController{}, "post:GetAllContactNoteList")

	// 客户来源
	beego.Router("/api/v1/customer/source/list", &backend.CrmCustomerSourceController{}, "get:GetAllSource")
	beego.Router("/api/v1/customer/me/list", &backend.CustomerController{}, "post:MyCustomerList")
}


