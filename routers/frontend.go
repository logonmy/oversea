package routers

import (
	"github.com/astaxie/beego"
	"oversea/app/controllers/frontend"
)

func frontendRouters()  {

	beego.SetStaticPath("/fonts","static/frontend/fonts")
	beego.SetStaticPath("/images","static/frontend/images")
	beego.SetStaticPath("/css","static/frontend/css")
	beego.SetStaticPath("/js","static/frontend/js")
	beego.SetStaticPath("/excel","static/frontend/excel")

	beego.Router("/", &frontend.IndexController{}, "get:Home")
	beego.Router("/travel", &frontend.TravelController{}, "get:Index")
	beego.Router("/about", &frontend.AboutController{}, "get:Index")
	beego.Router("/contact", &frontend.ContactController{}, "get:Index")
	beego.Router("/news", &frontend.NewsController{}, "get:Index")
}