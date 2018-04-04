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
	this.pageSize = 1

	this.initAuth()
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
	this.auth = services.SysAuthService
	this.auth.Init(token)
	this.userId = this.auth.GetUserId()


	this.Data["auth"] = this.auth
	this.Data["adminEntity"] = this.auth.GetUser()
	if !this.auth.IsLogined() {
		if this.actionName != "logout" && this.actionName != "login" {
			    this.Ctx.ResponseWriter.WriteHeader(stdout.HttpNotAuthorization)
				this.StdoutError(stdout.NotAuthorizationError,  stdout.NotAuthorizationErrorMsg)
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

// StdoutSuccess 输出带分页查询信息的结构-完成
func (this *AdminBaseController) StdoutQuerySuccess(offset, limit int, count int64, data interface{}) {
	s :=  this.makeStdJSON(stdout.Success)
	s.ErrMsg = stdout.MsgSuccess
	if data != nil {
		s.Data = data
	}
	s.Query(offset, limit, count)
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

func (this *AdminBaseController) makeStdJSON(code int) *StdJSON {
	return &StdJSON{
		ErrCode: code,
		Time:    time.Now().Unix(),
	}
}

func (this *AdminBaseController) checkError(err error) {
	if err != nil {
		this.StdoutError(-1, err.Error())
	}
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

func (s *StdJSON) Query(offset, limit int, count int64) *StdJSON {
	s.DataQuery = map[string]interface{}{
		"offset": offset,
		"limit":  limit,
		"count":  count,
	}
	return s
}

func (s *StdJSON) Extra(key string, data interface{}) *StdJSON {
	if key == "tagMap" {
		s.TagMap = data
		return s
	}
	if s.DataExtra == nil {
		s.DataExtra = make(map[string]interface{}, 2)
	}
	s.DataExtra[key] = data
	return s
}

