package backend

import (
	"oversea/app/services"
	"oversea/app/stdout"
	"oversea/utils"
	"encoding/json"
	"oversea/app/entitys"
	"oversea/app/form/backend"
)

type AdminUserController struct {
	AdminBaseController
}

// 列表
func (this *AdminUserController) List() {

	var userForm backend.UserForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &userForm)

	if userForm.Page < 1 {
		userForm.Page = 1
	}

	if userForm.PageSize < 1 {
		userForm.PageSize = this.pageSize
	}

	filters := make([]interface{}, 0)

	id := this.GetString("id")
	if id != "" {
		filters = append(filters, "id", id)
	}

	userList, count := services.SysUserService.GetAdminUsersList(userForm.Page, userForm.PageSize , filters...)

	this.StdoutQuerySuccess(userForm.Page, userForm.PageSize, count, userList)
}

func (this *AdminUserController) Add() {
	v := entitys.AdminUser{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		if utils.IsEmpty(v.UserName) || utils.IsEmpty(v.Password) {
			this.StdoutError(stdout.ParamsError, stdout.UsernameOrPasswdEmptyError)
		}

		if !utils.IsEmpty(v.Email) && !utils.IsEmail(v.Email) {
			this.StdoutError(stdout.ParamsError, stdout.EmailAddressError)
		}

		if !utils.IsEmpty(v.Phone) && !utils.IsMobilePhone(v.Phone) {
			this.StdoutError(stdout.ParamsError, stdout.MobilePhoneError)
		}

		_, e := services.SysUserService.GetUserByName(v.UserName)
		if e == nil {
			this.StdoutError(stdout.DBError, stdout.UserIsExists)
		}

		_, err := services.SysUserService.AddUser(v.UserName, v.Phone, v.Email, v.Password, v.Sex)
		if err != nil {
			this.StdoutError(stdout.DBError, stdout.AddAdminFailError)
		}

		this.StdoutSuccess(nil)
	}

}

func (this *AdminUserController) Edit() {

	var userForm backend.UserForm
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userForm)
	this.checkError(err)

	user, err := services.SysUserService.GetUser(userForm.Id)
	this.checkError(err)


	if !utils.IsEmpty(userForm.Email) && !utils.IsEmail(userForm.Email) {
		this.StdoutError(stdout.ParamsError, stdout.EmailAddressError)
	}

	if !utils.IsEmpty(userForm.Phone) && !utils.IsMobilePhone(userForm.Phone) {
		this.StdoutError(stdout.ParamsError, stdout.MobilePhoneError)
	}

	user.Phone = userForm.Phone
	user.Email = userForm.Email

	err = services.SysUserService.UpdateAdminUser(user, "Phone", "email")

	if err != nil {
		this.StdoutError(stdout.DBError, stdout.UpdateError)
	}
	this.StdoutSuccess(nil)
}

// 获取客户信息
func (this *AdminUserController) GetInfo() {

	uid, _ := this.GetInt("id", 0)
	if uid <= 0 {
		this.StdoutError(stdout.ParamsError, stdout.ParamsErrorMsg, nil)
	}

	customer, err := services.SysUserService.GetUser(uid)
	if err != nil {
		this.StdoutError(stdout.DBError, err.Error(), nil)
	}

	this.StdoutSuccess(customer)
}

// 我的信息
func (this *AdminUserController) GetMyInfo() {

	customer, err := services.SysUserService.GetUser(this.userId)
	if err != nil {
		this.StdoutError(stdout.DBError, err.Error(), nil)
	}

	this.StdoutSuccess(customer)
}

// 修改密码
func (this *AdminUserController) ChangePassword() {

	var userForm backend.UserForm
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userForm)
	this.checkError(err)

	user, err := services.SysUserService.GetUser(this.userId)
	this.checkError(err)

	oldPassword := utils.MD5(user.Password + user.Salt)
	confirmOldPassword := utils.MD5(userForm.Password + user.Salt)
	if oldPassword != confirmOldPassword {
		this.StdoutError(stdout.HttpStatusFail, stdout.OldPasswordError, nil)
	}

	user.Salt = utils.NewNoDashUUID()
	user.Password = utils.MD5(userForm.Password + user.Salt)

	err = services.SysUserService.UpdateAdminUser(user, "Password")

	if err != nil {
		this.StdoutError(stdout.DBError, stdout.UpdateError)
	}
	this.StdoutSuccess(nil)
}