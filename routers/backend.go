package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"oversea/app/controllers/backend"
	"oversea/utils"
	"oversea/middleware"
	"github.com/astaxie/beego/logs"
	"oversea/app/stdout"
	"time"
	"github.com/astaxie/beego/context"
)

func adminRouters() {

	// 验证码路由
	beego.Handler("/captcha/*.png", captcha.Server(130, 40))

	ns :=
		beego.NewNamespace("/api",
			beego.NSNamespace("/v1",
				beego.NSRouter("/private/login", &backend.MainController{}, "post:Login"),
				beego.NSNamespace("/backend",
					// 后续添加验证 黑白名单，反扒之类的
					beego.NSCond(func(ctx *context.Context) bool {
						ip := utils.GetRemoteIpAddress(ctx.Request)
						blacklist := &middleware.IpBlacklist{}
						logs.Info("ip =======", ip)
						if blacklist.IsIPInBacklist(ip) {
							//ctx.ResponseWriter.WriteHeader(stdout.HttpForbidded)
							type stdoutJson struct {
								ErrCode int    `json:"errorCode"`
								ErrMsg  string `json:"errorMsg"`
								Time    int64  `json:"time"`
							}
							s := &stdoutJson{
								ErrCode: stdout.HttpNotAuthorization,
								ErrMsg:  "禁止访问",
								Time:    time.Now().Unix(),
							}
							ctx.Output.JSON(s, true, false)

							return false
						}
						return true
					}),
					beego.NSRouter("/user/getAll", &backend.AdminUserController{}, "get:GetAllUserList"),
					beego.NSRouter("/user/getAll", &backend.AdminUserController{}, "get:GetAllUserList"),
					beego.NSRouter("/user/getInfo", &backend.AdminUserController{}, "get:GetInfo"),
					beego.NSRouter("/user/list", &backend.AdminUserController{}, "get,post:List"),
					beego.NSRouter("/logout", &backend.MainController{}, "get:Logout"),
					beego.NSRouter("/user/add", &backend.AdminUserController{}, "post:Add"),
					beego.NSRouter("/user/edit", &backend.AdminUserController{}, "post:Edit"),
					beego.NSRouter("/user/profile", &backend.AdminUserController{}, "get:GetMyInfo"),
					beego.NSRouter("/user/changePwd", &backend.AdminUserController{}, "post:ChangePassword"),
					// 客户
					beego.NSRouter("/customer/list", &backend.CustomerController{}, "post:List"),
					beego.NSRouter("/customer/getInfo", &backend.CustomerController{}, "get:GetInfo"),
					beego.NSRouter("/customer/add", &backend.CustomerController{}, "post:AddCrmCustomer"),
					beego.NSRouter("/customer/update", &backend.CustomerController{}, "post:UpdateCrmCustomerById"),
					beego.NSRouter("/customer/delete", &backend.CustomerController{}, "get:DeleteCrmCustomer"),
					beego.NSRouter("/customer/export", &backend.CustomerController{}, "get:Export"),
					beego.NSRouter("/customer/me/export", &backend.CustomerController{}, "get:MyCustomerExport"),

					// 联系人
					beego.NSRouter("/linkman/getInfo", &backend.CrmLinkmanController{}, "get:GetCrmLinkmanById"),
					beego.NSRouter("/linkman/all/list", &backend.CrmLinkmanController{}, "post:GetAllCrmLinkmanList"),
					beego.NSRouter("/linkman/me/list", &backend.CrmLinkmanController{}, "post:GetMyCrmLinkmanList"),
					beego.NSRouter("/linkman/update", &backend.CrmLinkmanController{}, "post:UpdateCrmLinkmanById"),
					beego.NSRouter("/linkman/add", &backend.CrmLinkmanController{}, "post:AddCrmLinkman"),
					beego.NSRouter("/linkman/delete", &backend.CrmLinkmanController{}, "get:DeleteCrmLinkman"),

					beego.NSRouter("/contact/add", &backend.ContactController{}, "post:AddContactNote"),
					beego.NSRouter("/contact/list", &backend.ContactController{}, "post:GetAllContactNoteList"),

					// 客户来源
					beego.NSRouter("/customer/assignTo", &backend.CustomerController{}, "post:AssignTo"),

					beego.NSRouter("/customer/source/list", &backend.CrmCustomerSourceController{}, "get:GetAllSource"),
					beego.NSRouter("/customer/me/list", &backend.CustomerController{}, "post:MyCustomerList"),

					// 上级
					beego.NSRouter("/customer/chance", &backend.CrmChanceController{}, "get:GetCustomerChanceList"),
				),
			),
		)

	beego.AddNamespace(ns)

}
