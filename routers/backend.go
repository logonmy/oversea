package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"oversea/app/controllers/backend"
)

func adminRouters()  {

	beego.SetStaticPath("/admin/fonts","static/backend/fonts")
	beego.SetStaticPath("/admin/images","static/backend/images")
	beego.SetStaticPath("/admin/css","static/backend/css")
	beego.SetStaticPath("/admin/js","static/backend/js")

	// 验证码路由
	beego.Handler("/captcha/*.png", captcha.Server(130, 40))

	beego.Router("/admin/home/index", &backend.MainController{}, "*:Index")
	beego.Router("/admin/user/list", &backend.AdminUserController{}, "*:List")
	beego.Router("/admin/login", &backend.MainController{}, "*:Login")
	beego.Router("/admin/logout", &backend.MainController{}, "*:Logout")
	beego.Router("/admin/user/add", &backend.AdminUserController{}, "*:Add")
	beego.Router("/admin/user/edit", &backend.AdminUserController{}, "*:Edit")

}


