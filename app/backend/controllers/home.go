package controllers

import (
	"oversea/app/backend/services"
	"github.com/astaxie/beego"
	"oversea/app/backend/libs"
)

type HomeController struct {
	AdminBaseController
}

// 列表
func (this *HomeController) Index() {
	this.Data["pageTitle"] = "跳板机列表"
	this.display()
}

// 列表
func (this *HomeController) List() {

	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}
	filters := make([]interface{}, 0)

	id := this.GetString("id")
	if id != "" {
		filters = append(filters, "id", id)
	}

	userList, count := services.BackUserService.GetAdminUsersList(page, this.pageSize, filters...)

	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("HomeController.List",
		filters...),
		true).ToString()
	this.Data["pageTitle"] = "管理员列表"
	this.Data["userList"] = userList
	this.display()
}