package entitys

import "time"

type CrmCustomer struct {
	Id         int       `orm:"column(cust_id);auto"`
	Name       string    `orm:"column(name);size(64)" description:"客户姓名"`
	Source     int       `orm:"column(source);null" description:"客户来源"`
	Level      int       `orm:"column(level);null" description:"客户等级"`
	Website    string    `orm:"column(website);size(256);null" description:"客户个人网站地址"`
	Mobile     string    `orm:"column(mobile);size(64)" description:"客户手机号码"`
	Tel        string    `orm:"column(tel);size(256);null" description:"客户电话号码"`
	Fax        string    `orm:"column(fax);size(256);null" description:"客户传真"`
	Email      string    `orm:"column(email);size(256);null" description:"客户邮箱地址"`
	Status     string    `orm:"column(status);size(256);null" description:"客户状态"`
	Intro      string    `orm:"column(intro);null" description:"客户简介"`
	CreateBy   int       `orm:"column(create_by)" description:"创建者"`
	AssignTo   int       `orm:"column(assign_to)" description:"指派给"`
	AssignTime time.Time `orm:"column(assign_time);type(timestamp);null" description:"指派日期"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
}

func (t *CrmCustomer) TableName() string {
	return "customer"
}