package backend

import (
	"oversea/app/services"
	"oversea/app/stdout"
	"oversea/utils"
)

type AdminUserController struct {
	AdminBaseController
}

// 列表
func (this *AdminUserController) List() {

	page, _ := this.GetInt("page")
	pageSize, _:= this.GetInt("pageSize")
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = this.pageSize
	}

	filters := make([]interface{}, 0)

	id := this.GetString("id")
	if id != "" {
		filters = append(filters, "id", id)
	}

	userList, count := services.SysUserService.GetAdminUsersList(page, pageSize, filters...)

	this.StdoutQuerySuccess(page, page, count, userList)
}

func (this *AdminUserController) Add() {

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

	this.StdoutSuccess(nil)

}

func (this *AdminUserController) Edit() {

	id, _ := this.GetInt("id")
	user, err := services.SysUserService.GetUser(id)
	this.checkError(err)

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

	err = services.SysUserService.UpdateAdminUser(user, "Phone", "email")

	if err != nil {
		this.StdoutError(stdout.DBError, stdout.UpdateError)
	}
	this.StdoutSuccess(nil)
}