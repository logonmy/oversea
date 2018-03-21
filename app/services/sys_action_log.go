package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"oversea/utils"
)

// 系统动态
type actionLogService struct{}

// 添加记录
func (this *actionLogService) Add(action,objectType string, objectId int, extra string) bool {
	o := orm.NewOrm()
	act := new(entitys.ActionLog)
	act.Action = action
	act.ObjectType = objectType
	act.ObjectId = objectId
	act.Extra = extra
	act.Actor = SysAuthService.GetUserName()
	act.UserId = SysAuthService.GetUserId()
	o.Insert(act)
	return true
}

// 登录动态
func (this *actionLogService) Login() {
	ip := SysAuthService.GetLastIp()
	this.Add("login", "user", 1, ip)
}

// 退出登录
func (this *actionLogService) Logout() {
	ip := utils.GetIpAddress()
	this.Add("logout", "user", 1, ip)
}