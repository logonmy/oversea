package routers

import (
	"github.com/astaxie/beego"
	"oversea/app/backend/controllers"
)

func adminRouters()  {
	beego.Router("/admin/home/index", &controllers.HomeController{}, "*:Index")
}
