package entitys

import (
	"time"
)

type AdminUser struct {
	Id         int
	UserName   string    `orm:"unique;size(20)"`             // 用户名
	Password   string    `orm:"size(32)"`                    // 密码
	Salt       string    `orm:"size(255)"`                    // 密码盐
	Sex        int       `orm:"default(0)"`                  // 性别
	Email      string    `orm:"size(50)"`                    // 邮箱
	Phone      string    `orm:"size(12)"`                    // 手机号
	LastLogin  time.Time `orm:"null;type(datetime)"`         // 最后登录时间
	LastIp     string    `orm:"size(15)"`                    // 最后登录IP
	Status     int       `orm:"default(0)"`                  // 状态，0正常 -1禁用
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`     // 更新时间
}

// 表名
func (m *AdminUser) TableName() string {
	return "admin_user"
}


type User struct {
	Id         int
	UserName   string    `orm:"unique;size(20)"`             // 用户名
	Sex        int       `orm:"default(0)"`                  // 性别
	Email      string    `orm:"size(50)"`                    // 邮箱
	Phone      string    `orm:"size(12)"`                    // 手机号
	LastLogin  time.Time `orm:"null;type(datetime)"`         // 最后登录时间
	LastIp     string    `orm:"size(15)"`                    // 最后登录IP
	Status     int       `orm:"default(0)"`                  // 状态，0正常 -1禁用
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`     // 更新时间
	Avatar     string
}
