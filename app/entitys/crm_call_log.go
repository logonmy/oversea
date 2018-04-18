package entitys

import "time"

type CrmCallLog struct {
	Id             int       `orm:"column(id);pk"`
	Tel            string    `orm:"column(tel);size(20);null" description:"固定电话/手机号"`
	Content        string    `orm:"column(content);size(255);null" description:"来电内容"`
	Uid            int       `orm:"column(uid)" description:"处理人"`
	Status         int       `orm:"column(status)" description:"处理状态:0-待处理，1-处理中，2-已处理"`
	ProcessTime    time.Time `orm:"column(process_time);type(timestamp);null" description:"处理时间"`
	CallTypeId     int       `orm:"column(call_type_id)" description:"来电类型"`
	CustomerId     int       `orm:"column(customer_id)" description:"客户"`
	ProcessContext string    `orm:"column(process_context);size(255);null" description:"处理信息"`
	LinkmanId      int       `orm:"column(linkman_id)" description:"联系人"`
	CreateTime     time.Time `orm:"column(create_time);type(timestamp);auto_now_add" description:"创建时间"`
	UpdateTime     time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
}

func (t *CrmCallLog) TableName() string {
	return "call_log"
}