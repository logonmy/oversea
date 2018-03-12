package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"oversea/app/backend/constant"
	"strings"
	"github.com/dchest/captcha"
)

type AdminBaseController struct {
	beego.Controller
	userId         int    // 当前登录的用户id
	controllerName string // 控制器名
	actionName     string // 动作名
	pageSize       int    // 默认分页大小
	lang           string // 当前语言环境
	menuList       []Menu // 当前菜单
}

type Menu struct {
	Name    string
	Route   string
	Icon    string
	Submenu []SubMenu
}

type SubMenu struct {
	Name   string
	Route  string
	Action string
}

// StdJSON 标准输出JSON格式
type StdJSON struct {
	ErrCode   int                    `json:"errorCode"`
	ErrMsg    string                 `json:"errorMsg"`
	ErrDesc   string                 `json:"errDesc,omitempty"`
	Data      interface{}            `json:"data,omitempty"`
	DataQuery interface{}            `json:"query,omitempty"`
	DataExtra map[string]interface{} `json:"extra,omitempty"`
	TagMap    interface{}            `json:"tagMap,omitempty"`
	Time      int64                  `json:"time"`
}

func (this *AdminBaseController) Prepare() {
	//this.Ctx.Output.Header("X-Powered-By", "GoPub/"+beego.AppConfig.String("version"))
	//this.Ctx.Output.Header("X-Author-By", "lisijie")
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.pageSize = 20
}

//渲染模版
func (this *AdminBaseController) display(tpl ...string) {
	var tplname string
	tpldir := "backend/"
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
	this.LayoutSections["Sidebar"] = tpldir + "layout/sections/sidebar.html"

	this.Data["website"] = beego.AppConfig.String("website")
	this.Data["xsrf_token"] = this.XSRFToken()
}

// 重定向
func (this *AdminBaseController) redirect(url string) {
	if this.IsAjax() {
		//this.showMsg("", MSG_REDIRECT, url)
	} else {
		this.Redirect(url, 302)
		this.StopRun()
	}
}

func (this *AdminBaseController) makeStdJSON(code int) *StdJSON {
	return &StdJSON{
		ErrCode: code,
		Time:    time.Now().Unix(),
	}
}

// StdoutSuccess 输出结构-完成
func (this *AdminBaseController) StdoutSuccess(data interface{}) {
	s := this.makeStdJSON(constant.Success)
	s.ErrMsg = constant.MsgSuccess
	if data != nil {
		s.Data = data
	}
	this.Data["json"] = s
	this.ServeFormatted()
	this.StopRun()
}

// StdoutError 输出结构-失败
func (this *AdminBaseController) StdoutError(code int, err error, data ...interface{}) {
	s := this.makeStdJSON(code)
	s.ErrMsg = constant.ErrCode2Msg(code)
	if err != nil {
		s.ErrDesc = err.Error()
	}
	if len(data) != 0 {
		s.Data = data[0]
	}
	this.Data["json"] = s
	this.ServeFormatted()
	this.StopRun()
}

// 获取验证码
func (this *AdminBaseController) getCaptchaMap() map[string]interface{} {

	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}

	captchaMap := map[string]interface{}{
		"captchaId":  d.CaptchaId,
		"captchaUrl": "/captcha/" + d.CaptchaId + ".png",
	}
	return captchaMap
}

func (this *AdminBaseController) isPost() bool {

	return this.Ctx.Input.IsPost()
}


func (this *AdminBaseController) isAjax() bool {

	return this.Ctx.Input.IsAjax()
}