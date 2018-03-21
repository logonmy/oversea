package backend

import (
	"oversea/app/services"
	"github.com/astaxie/beego"
	"oversea/app/libs"
	"oversea/app/stdout"
	"oversea/utils"
)

type AdminUserController struct {
	AdminBaseController
}

// 列表
func (this *AdminUserController) List() {

	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}
	filters := make([]interface{}, 0)

	id := this.GetString("id")
	if id != "" {
		filters = append(filters, "id", id)
	}

	userList, count := services.SysUserService.GetAdminUsersList(page, this.pageSize, filters...)

	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("HomeController.List",
		filters...),
		true).ToString()
	this.Data["pageTitle"] = "管理员列表"
	this.Data["userList"] = userList
	this.display()
}

func (this *AdminUserController) Add() {

	if this.isPost() {
		realName := this.GetString("real_name")
		phone := this.GetString("phone")
		email := this.GetString("email")
		password := this.GetString("password")
		repassword := this.GetString("repassword")
		sex, _ := this.GetInt("sex", 1)

		if utils.IsEmpty(realName) || utils.IsEmpty(password) {
			this.StdoutError(stdout.ParamsError, stdout.UsernameOrPasswdEmptyError)
		}

		if !utils.IsEmpty(email) && !utils.IsEmail(email) {
			this.StdoutError(stdout.ParamsError, stdout.EmailAddressError)
		}

		if !utils.IsEmpty(phone) && !utils.IsMobilePhone(phone) {
			this.StdoutError(stdout.ParamsError, stdout.MobilePhoneError)
		}

		if password != repassword {
			this.StdoutError(stdout.ParamsError, stdout.RepeatPasswordError)
		}

		_, e := services.SysUserService.GetUserByName(realName)
		if e == nil {
			this.StdoutError(stdout.DBError, stdout.UserIsExists)
		}

		_, err := services.SysUserService.AddUser(realName, phone, email, password, sex)
		if err != nil {
			this.StdoutError(stdout.DBError, stdout.AddAdminFailError)
		}

		this.StdoutSuccess(map[string]string{"href": beego.URLFor("AdminUserController.List")})
	}
	this.Data["pageTitle"] = "添加管理员账号"
	this.display()
}

func (this *AdminUserController) Edit() {

	id, _ := this.GetInt("id")
	user, err := services.SysUserService.GetUser(id)
	this.checkError(err)

	if this.isPost() {
		phone := this.GetString("phone")
		email := this.GetString("email")

		if !utils.IsEmpty(email) && !utils.IsEmail(email) {
			this.StdoutError(stdout.ParamsError, stdout.EmailAddressError)
		}

		if !utils.IsEmpty(phone) && !utils.IsMobilePhone(phone) {
			this.StdoutError(stdout.ParamsError, stdout.MobilePhoneError)
		}

		user.Phone = phone
		user.Email = email

		err := services.SysUserService.UpdateAdminUser(user, "Phone", "email")

		if err != nil {
			this.StdoutError(stdout.DBError, stdout.UpdateError)
		}
		this.StdoutSuccess(nil)
	}

	this.Data["userInfo"] = user
	this.Data["pageTitle"] = "编辑管理员账号"
	this.display()
}