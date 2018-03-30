package backend

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"oversea/app/stdout"
	"oversea/app/services"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	AdminBaseController
}

// 列表
func (this *MainController) Index() {
	this.Data["pageTitle"] = "后台首页"
	this.display()
}

// 登录
func (c *MainController) Login() {
	if c.userId > 0 {
		c.redirect("/admin/home/index")
	}

	//beego.ReadFromRequest(&c.Controller)
	if c.isPost() {
		phone := c.GetString("username")
		if phone == `` {
			phone = c.GetString("userName")
		}
		password := c.GetString("password")
		captchaId := c.GetString("captchaId")
		captchaValue := c.GetString("captcha")

		remember := c.GetString("rememberme")


		if !captcha.VerifyString(captchaId, captchaValue) {
			//c.StdoutError(stdout.ParamsError, stdout.CaptchaError, c.getCaptchaMap())
		}

		if phone != "" && password != "" {
			token, err := services.SysAuthService.Login(phone, password)

			if err != nil {
				c.StdoutError(stdout.ParamsError, err.Error())
			}

			if remember == "yes" {
				c.Ctx.SetCookie("auth", token, 7*86400)
			} else {
				c.Ctx.SetCookie("auth", token)
			}
			services.SysActionLogService.Login()
			c.StdoutSuccess(nil)
		} else {
			c.StdoutError(stdout.ParamsError, stdout.UsernameOrPasswdEmptyError, c.getCaptchaMap())
		}
	}
	c.Data["captcha"] = c.getCaptchaMap()
	c.setTpl()
}

// 退出登录
func (this *MainController) Logout() {
	services.SysActionLogService.Logout()
	services.SysAuthService.Logout()
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor(".Login"))
}
