package entitys

import "time"

type CmsUsers struct {
	Id             int       `orm:"column(id);auto"`
	UserLogin      string    `orm:"column(user_login);size(60)" description:"用户名"`
	UserPass       string    `orm:"column(user_pass);size(64)" description:"登录密码；sp_password加密"`
	UserNicename   string    `orm:"column(user_nicename);size(50)" description:"用户美名"`
	UserEmail      string    `orm:"column(user_email);size(100)" description:"登录邮箱"`
	UserUrl        string    `orm:"column(user_url);size(100)" description:"用户个人网站"`
	Avatar         string    `orm:"column(avatar);size(255);null" description:"用户头像，相对于upload/avatar目录"`
	Sex            int16     `orm:"column(sex);null" description:"性别；0：保密，1：男；2：女"`
	Birthday       time.Time `orm:"column(birthday);type(date);null" description:"生日"`
	Signature      string    `orm:"column(signature);size(255);null" description:"个性签名"`
	LastLoginIp    string    `orm:"column(last_login_ip);size(16);null" description:"最后登录ip"`
	LastLoginTime  time.Time `orm:"column(last_login_time);type(datetime);null" description:"最后登录时间"`
	ValidationType string    `orm:"column(validation_type);size(50);null" description:"验证类型(用户激活,重置密码,邮箱激活)"`
	ValidationKey  string    `orm:"column(validation_key);size(100);null" description:"验证KEY"`
	Salt           string    `orm:"column(salt);size(32);null" description:"加密混淆码"`
	UserStatus     int       `orm:"column(user_status)" description:"用户状态 0：禁用； 1：正常 ；2：未验证"`
	Mobile         string    `orm:"column(mobile);size(20)" description:"手机号"`
	QqOpenid       string    `orm:"column(qq_openid);size(64);null" description:"qq openid"`
	WeiboUid       string    `orm:"column(weibo_uid);size(64);null" description:"weibo uid"`
	WeixinOpenid   string    `orm:"column(weixin_openid);size(64);null" description:"weixin openid"`
	CreateTime     time.Time `orm:"column(create_time);type(timestamp);auto_now_add" description:"注册时间"`
	UpdateTime     time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
}

func (t *CmsUsers) TableName() string {
	return "users"
}
