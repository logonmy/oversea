package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.SetStaticPath("/admin/fonts","static/backend/fonts")
	beego.SetStaticPath("/admin/images","static/backend/images")
	beego.SetStaticPath("/admin/css","static/backend/css")
	beego.SetStaticPath("/admin/js","static/backend/js")

	adminRouters()

	// 允许跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}
