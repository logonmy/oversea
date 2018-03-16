package routers

import (
	"github.com/astaxie/beego"
	"oversea/app/frontend/controllers"
)

func frontendRouters()  {

	beego.SetStaticPath("/fonts","static/frontend/fonts")
	beego.SetStaticPath("/images","static/frontend/images")
	beego.SetStaticPath("/css","static/frontend/css")
	beego.SetStaticPath("/js","static/frontend/js")

	beego.Router("/", &controllers.IndexController{}, "get:Home")
}