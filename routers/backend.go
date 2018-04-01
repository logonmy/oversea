package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"oversea/app/controllers/backend"
)

func adminRouters()  {

	// 验证码路由
	beego.Handler("/captcha/*.png", captcha.Server(130, 40))

	beego.Router("/api/v1/user/list", &backend.AdminUserController{}, "get:List")
	beego.Router("/api/v1/login", &backend.MainController{}, "post:Login")
	beego.Router("/api/v1/logout", &backend.MainController{}, "get:Logout")
	beego.Router("/api/v1/user/add", &backend.AdminUserController{}, "post:Add")
	beego.Router("/api/v1/user/edit", &backend.AdminUserController{}, "post:Edit")

}


