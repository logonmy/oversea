package backend

import "time"

type UserForm struct {
	PageSize   int
	Page       int
	Id         int
	Password   string
	OldPassword string
	UserName   string    `orm:"unique;size(20)"`             // 用户名
	Sex        int       `orm:"default(0)"`                  // 性别
	Email      string    `orm:"size(50)"`                    // 邮箱
	Phone      string    `orm:"size(12)"`                    // 手机号
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`     // 更新时间
}
