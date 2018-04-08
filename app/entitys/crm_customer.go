package entitys

import "time"

type CrmCustomer struct {
	Id            int       `orm:"column(cust_id);auto"`
	Name          string    `orm:"column(name);size(64)" description:"客户姓名"`
	Source        int       `orm:"column(source);null" description:"客户来源"`
	Level         int       `orm:"column(level);null" description:"客户等级"`
	Website       string    `orm:"column(website);size(256);null" description:"客户个人网站地址"`
	Mobile        string    `orm:"column(mobile);size(64)" description:"客户手机号码"`
	Tel           string    `orm:"column(tel);size(256);null" description:"客户电话号码"`
	Fax           string    `orm:"column(fax);size(256);null" description:"客户传真"`
	Email         string    `orm:"column(email);size(256);null" description:"客户邮箱地址"`
	Wechat        string    `orm:"column(wechat);size(256);null" description:"微信"`
	Status        int       `orm:"column(status);null" description:"客户状态: 0-正常，1-禁用"`
	Sex           int       `orm:"column(sex);null" description:"性别: 0-未知，1-男， 2-女"`
	NativePlace   string    `orm:"column(native_place);size(10);null" description:"籍贯"`
	Address       string    `orm:"column(address);size(255);null" description:"家庭住址"`
	LinkAddress   string    `orm:"column(link_address);size(255);null" description:"联系地址"`
	Birthday      string    `orm:"column(birthday);type(date);null" description:"生日"`
	Idcard        string    `orm:"column(idcard);size(30);null" description:"身份证"`
	Qq            string    `orm:"column(qq);size(20);null" description:"QQ"`
	Age           int       `orm:"column(age);" description:"年龄"`
	Capital       string    `orm:"column(capital);size(30);null" description:"资本描素"`
	Intension     string    `orm:"column(intension);size(100);null" description:"移民意向"`
	Intro         string    `orm:"column(intro);null" description:"客户简介"`
	CreateBy      int       `orm:"column(create_by)" description:"创建者"`
	AssignTo      int       `orm:"column(assign_to)" description:"指派给"`
	AssignStatus  int       `orm:"column(assign_status)" description:"指派状态: 0-未指派，1-已指派，2-无需指派"`
	AssignTime    time.Time `orm:"column(assign_time);type(timestamp);null" description:"指派日期"`
	ContactedDate time.Time `orm:"column(contacted_date);type(datetime);null" description:"拜访日期"`
	NextDate      time.Time `orm:"column(next_date);type(datetime);null" description:"下次拜访日期"`
	CreateTime    time.Time `orm:"column(create_time);type(timestamp);auto_now_add" description:"创建时间"`
	UpdateTime    time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
}

func (t *CrmCustomer) TableName() string {
	return "customer"
}
