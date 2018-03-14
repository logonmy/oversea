package routers

import (
	"github.com/astaxie/beego"
	"oversea/app/backend/controllers"
	"github.com/dchest/captcha"
)

func adminRouters()  {
	// 验证码路由
	beego.Handler("/captcha/*.png", captcha.Server(130, 34))

	beego.Router("/admin/home/index", &controllers.MainController{}, "*:Index")
	beego.Router("/admin/user/list", &controllers.AdminUserController{}, "*:List")
	beego.Router("/admin/login", &controllers.MainController{}, "*:Login")
	beego.Router("/admin/user/add", &controllers.AdminUserController{}, "*:Add")
	beego.Router("/admin/user/edit", &controllers.AdminUserController{}, "*:Edit")

}


