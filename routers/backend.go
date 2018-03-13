package routers

import (
	"github.com/astaxie/beego"
	"oversea/app/backend/controllers"
	"github.com/dchest/captcha"
)

func adminRouters()  {
	// 验证码路由
	beego.Handler("/captcha/*.png", captcha.Server(130, 34))

	beego.Router("/admin/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/admin/home/list", &controllers.HomeController{}, "*:List")
	beego.Router("/admin/login", &controllers.MainController{}, "*:Login")


}


