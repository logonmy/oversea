package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	frontendRouters()
	adminRouters()

	// 允许跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  false,
		AllowOrigins:[]string{"http://localhost:8000"},
		AllowMethods:     []string{"PUT","POST","GET","DELETE"},
		//AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", " Access-Control-Request-Method", "Access-Control-Allow-Origin",
		"Access-Control-Request-Headers", "mode"},
		AllowCredentials: true,
	}))
}
