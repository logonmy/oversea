package backend

import (
	"github.com/dchest/captcha"
	"oversea/app/stdout"
	"oversea/app/services"
)

type MainController struct {
	AdminBaseController
}


// 登录
func (c *MainController) Login() {

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

// 退出登录
func (this *MainController) Logout() {
	services.SysActionLogService.Logout()
	services.SysAuthService.Logout()
	this.Ctx.SetCookie("auth", "")
}
