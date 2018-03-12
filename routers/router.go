package routers

import "github.com/astaxie/beego"

func init() {
	beego.SetStaticPath("/admin/fonts","static/backend/fonts")
	beego.SetStaticPath("/admin/images","static/backend/images")
	beego.SetStaticPath("/admin/css","static/backend/css")
	beego.SetStaticPath("/admin/js","static/backend/js")

	adminRouters()
}
