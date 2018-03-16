package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"oversea/app/backend/stdout"
	"oversea/app/backend/services"
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

	beego.ReadFromRequest(&c.Controller)
	if c.isPost() {
		phone := c.GetString("username")
		password := c.GetString("password")
		captchaId := c.GetString("captchaId")
		captchaValue := c.GetString("captcha")

		remember := c.GetString("rememberme")

		if !captcha.VerifyString(captchaId, captchaValue) {
			c.StdoutError(stdout.ParamsError, stdout.CaptchaError, c.getCaptchaMap())
		}

		if phone != "" && password != "" {
			token, err := services.BackAuthService.Login(phone, password)

			if err != nil {
				c.StdoutError(stdout.ParamsError, err.Error())
			}

			if remember == "yes" {
				c.Ctx.SetCookie("auth", token, 7*86400)
			} else {
				c.Ctx.SetCookie("auth", token)
			}
			services.BackActionLogService.Login()
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
	services.BackActionLogService.Logout()
	services.BackAuthService.Logout()
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor(".Login"))
}
