package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"oversea/utils"
	"fmt"
)

// 系统动态
type actionLogService struct{}

// 添加记录
func (this *actionLogService) Add(action, objectType string, objectId int, extra string, fileds ...int) bool {
	o := orm.NewOrm()
	act := new(entitys.ActionLog)
	act.Action = action
	act.ObjectType = objectType
	act.ObjectId = objectId
	act.Extra = extra
	act.Actor = SysAuthService.GetUserName()
	act.UserId = SysAuthService.GetUserId()
	len := len(fileds)
	if len > 0 {
		if len == 1 {
			act.Customer = fileds[0]
		} else {
			act.Customer = fileds[0]
			act.Contact = fileds[1]
		}
	}
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

// 添加沟通记录
func (this *actionLogService) AddContactLog(customerId, contactId int, extra string) {
	this.Add("addcontact", "contact", 2, extra, customerId, contactId)
}

// 获取动态列表
func (this *actionLogService) GetList(page, pageSize int, filters ...interface{}) ([]entitys.ActionLog, int64) {
	var list []entitys.ActionLog
	o := orm.NewOrm()
	query := o.QueryTable(new(entitys.ActionLog))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	num, err := query.OrderBy("-create_time").Offset((page - 1) * pageSize).Limit(
		pageSize).All(
		&list)
	if num > 0 && err == nil {
		for i := 0; i < int(num); i++ {
			this.format(&list[i])
		}
	}
	total, _ := query.Count()
	return list, total
}

// 格式化
func (this *actionLogService) format(action *entitys.ActionLog) {
	switch action.Action {
	case "login":
		action.Message = fmt.Sprintf("<b>%s</b> 登录系统，IP为 <b>%s</b>。", action.Actor, action.Extra)
	case "logout":
		action.Message = fmt.Sprintf("<b>%s</b> 退出系统。", action.Actor)
	case "update_profile":
		action.Message = fmt.Sprintf("<b>%s</b> 更新了个人资料。", action.Actor)
	case "create_task":
		action.Message = fmt.Sprintf("<b>%s</b> 创建了编号为 <b class='blue'>%d</b> 的发布单。", action.Actor, action.ObjectId)
	case "addcontact":
		user, err := SysUserService.GetUser(action.UserId)
		linkman, err1 := CrmLinkmanService.GetCrmLinkmanById(action.Contact)

		if err == nil && err1 == nil {
			action.Message = fmt.Sprintf("<b>%s</b>, <b>%s</b> 添加了沟通日志，联系人：<b>%s</b>，沟通内容：<b>%s</b>",
				action.CreateTime,user.UserName,
				linkman.Name, action.Extra)
		}

	}
}
