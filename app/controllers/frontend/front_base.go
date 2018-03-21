package frontend

import (
	"github.com/astaxie/beego"
	"strings"
	"time"
	"oversea/app/stdout"
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

// StdoutSuccess 输出结构-完成
func (this *FrontendBaseController) StdoutSuccess(data interface{}) {
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
func (this *FrontendBaseController) StdoutError(code int, errMsg string, data ...interface{}) {
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

func (this *FrontendBaseController) makeStdJSON(code int) *StdJSON {
	return &StdJSON{
		ErrCode: code,
		Time:    time.Now().Unix(),
	}
}

func (this *FrontendBaseController) checkError(err error) {
	if err != nil {
		if this.IsAjax() {
			this.StdoutError(-1, err.Error())
		}
		this.Ctx.WriteString(err.Error()) //后续用错误页面替换
	}
}