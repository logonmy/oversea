package entitys

import "time"

type CmsGuestbook struct {
	Id           int       `orm:"column(id);auto"`
	FullName     string    `orm:"column(full_name);size(50)" description:"留言者姓名"`
	Email        string    `orm:"column(email);size(100)" description:"留言者邮箱"`
	Title        string    `orm:"column(title);size(255);null" description:"留言标题"`
	Msg          string    `orm:"column(msg)" description:"留言内容"`
	Createtime   time.Time `orm:"column(createtime);type(timestamp);auto_now_add" description:"留言时间"`
	ReplyMsg     string    `orm:"column(reply_msg);size(255)" description:"回复信息"`
	Replier      string    `orm:"column(replier);size(50)" description:"答复者"`
	ReplyTime    time.Time `orm:"column(reply_time);type(timestamp);null" description:"答复时间"`
	IsEmailReply int16     `orm:"column(is_email_reply);null" description:"是否同时发送答复邮件，1：发送，0：不发送"`
	Status       int16     `orm:"column(status)" description:"留言状态，1：正常，0：删除"`
}

func (t *CmsGuestbook) TableName() string {
	return "guestbook"
}
