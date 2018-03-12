package controllers

import (
	"github.com/astaxie/beego"
	//"time"
	"github.com/dchest/captcha"
	"oversea/app/backend/constant"
	//"oversea/utils"
)

type MainController struct {
	AdminBaseController
}

// 登录
func (c *MainController) Login() {
	if c.userId > 0 {
		c.redirect("/")
	}
	beego.ReadFromRequest(&c.Controller)
	if c.isPost() {
		//phone := c.GetString("phone")
		//password := c.GetString("password")
		captchaId := c.GetString("captchaId")
		captchaValue := c.GetString("captcha")
		newCaptcha := c.getCaptchaMap()

		if !captcha.VerifyString(captchaId, captchaValue) {
			c.StdoutError (constant.CaptchaError, nil, newCaptcha)
		}

		//m , _ := models.GetMemberByPhone(phone)
		//if m == nil {
		//	c.StdoutError(constant.UserIsNoExistsError, nil, newCaptcha)
		//} else if  m.Password != utils.MD5(password) {
		//	c.StdoutError(constant.PasswordError, nil, newCaptcha)
		//}
		//
		//c.SetSession("adminid", m.MemberId)
		//c.SetSession("name", m.Phone)
		//c.SetSession("phone", m.Phone)
		//
		//m.LastLoginTime = time.Now().Unix()
		//_ = models.UpdateMember(m, "LastLoginTime")
		//
		//c.Data["json"] = utils.StdoutSuccess(m)
		c.ServeJSON()
	}

	c.TplName = "main/login.html"
}

// 退出登录
func (this *MainController) Logout() {
	//service.ActionService.Logout(this.auth.GetUser().UserName, this.auth.GetUserId(), this.getClientIp())
	//this.auth.Logout()
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor(".Login"))
}

