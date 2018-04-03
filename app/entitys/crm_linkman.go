package entitys

import "time"

type CrmLinkman struct {
	Id         int       `orm:"column(id);auto"`
	CustId     int       `orm:"column(cust_id);null" description:"客户id"`
	Name       string    `orm:"column(name);size(255);null" description:"联系人姓名"`
	Job        string    `orm:"column(job);size(255);null" description:"职业"`
	Call       string    `orm:"column(call);size(255);null" description:"固定电话"`
	Phone      string    `orm:"column(phone);size(255);null" description:"手机号"`
	Qq         string    `orm:"column(qq);size(255);null" description:"QQ"`
	Email      string    `orm:"column(email);size(255);null" description:"邮箱"`
	Main       uint64    `orm:"column(main);size(1);null"`
	Sex        int8      `orm:"column(sex);null"`
	Address    string    `orm:"column(address);size(255);null"`
	Intro      string    `orm:"column(intro);size(255);null"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
}

func (t *CrmLinkman) TableName() string {
	return "linkman"
}