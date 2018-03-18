package backend

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/dchest/captcha"
	"time"
	"oversea/app/stdout"
	"oversea/app/services"
)

type AdminBaseController struct {
	beego.Controller
	auth           *services.AuthService // 验证服务
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

func (this *AdminBaseController) Prepare() {
	//this.Ctx.Output.Header("X-Powered-By", "GoPub/"+beego.AppConfig.String("version"))
	this.Ctx.Output.Header("X-Author-By", "weilanzhuan")
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.pageSize = 10

	this.initAuth()
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

func (this *AdminBaseController) setTpl(tpl ...string) {
	var tplname string
	tpldir := "backend/"
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}
	this.TplName = tpldir + tplname

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

//登录验证
func (this *AdminBaseController) initAuth() {
	token := this.Ctx.GetCookie("auth")
	this.auth = services.BackAuthService
	this.auth.Init(token)
	this.userId = this.auth.GetUserId()


	this.Data["auth"] = this.auth
	this.Data["adminEntity"] = this.auth.GetUser()
	if !this.auth.IsLogined() {
		if this.controllerName != "main" ||
			(this.controllerName == "main" && this.actionName != "logout" && this.actionName != "login") {
			this.redirect(beego.URLFor("MainController.Login"))
		}

	} else {
		// 进行权限判断
	}
}

// StdoutSuccess 输出结构-完成
func (this *AdminBaseController) StdoutSuccess(data interface{}) {
	s :=  this.makeStdJSON(stdout.Success)
	s.ErrMsg = stdout.MsgSuccess
	if data != nil {
		s.Data = data
	}
	this.Data["json"] = s
	this.ServeFormatted()
	this.StopRun()
}

// StdoutError 输出结构-失败
func (this *AdminBaseController) StdoutError(code int, errMsg string, data ...interface{}) {
	s := this.makeStdJSON(code)
	s.ErrMsg = errMsg
	if len(data) != 0 {
		s.Data = data[0]
	}
	this.Data["json"] = s
	this.ServeFormatted()
	this.StopRun()
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

func (this *AdminBaseController) makeStdJSON(code int) *StdJSON {
	return &StdJSON{
		ErrCode: code,
		Time:    time.Now().Unix(),
	}
}

func (this *AdminBaseController) checkError(err error) {
	if err != nil {
		if this.IsAjax() {
			this.StdoutError(-1, err.Error())
		}
		this.Ctx.WriteString(err.Error()) //后续用错误页面替换
	}
}