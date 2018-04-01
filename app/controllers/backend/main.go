package backend

import (
	"oversea/app/stdout"
	"oversea/app/services"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"oversea/app/form/backend"
	"github.com/astaxie/beego/validation"
)

type MainController struct {
	AdminBaseController
}


// 登录
func (c *MainController) Login() {

	var loginForm backend.LoginForm
	json.Unmarshal(c.Ctx.Input.RequestBody,&loginForm)
	valid := validation.Validation{}
	_, err := valid.Valid(&loginForm)

	if err != nil{
		c.StdoutError(stdout.ParamsError, err.Error())
	}

	logs.Error(string(c.Ctx.Input.RequestBody))
	logs.Error(loginForm)
	token, err := services.SysAuthService.Login(loginForm.UserName, loginForm.Password)

	if err != nil {
		c.StdoutError(stdout.ParamsError, err.Error())
	}

	if loginForm.Remember == "yes" {
		c.Ctx.SetCookie("auth", token, 7*86400)
	} else {
		c.Ctx.SetCookie("auth", token)
	}
	services.SysActionLogService.Login()
	c.StdoutSuccess(nil)

}

// 退出登录
func (this *MainController) Logout() {
	services.SysActionLogService.Logout()
	services.SysAuthService.Logout()
	this.Ctx.SetCookie("auth", "")
}
