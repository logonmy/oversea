package frontend

import (
	"github.com/astaxie/beego"
	"strings"
)

type FrontendBaseController struct {
	beego.Controller
	controllerName string // 控制器名
	actionName     string // 动作名
}

func (this *FrontendBaseController) Prepare() {
	//this.Ctx.Output.Header("X-Powered-By", "GoPub/"+beego.AppConfig.String("version"))
	this.Ctx.Output.Header("X-Author-By", "weilanzhuan")
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	this.Data["website"] = beego.AppConfig.String("frontend_website")
	this.Data["xsrf_token"] = this.XSRFToken()
}

func (this *FrontendBaseController) display(tpl ...string) {
	var tplname string
	tpldir := "frontend/"
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}

	this.Layout = tpldir + "layout/layout.html"
	this.TplName = tpldir + tplname

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = tpldir + "layout/sections/header.html"
	this.LayoutSections["Footer"] = tpldir + "layout/sections/footer.html"
	this.LayoutSections["Navbar"] = tpldir + "layout/sections/navbar.html"
}