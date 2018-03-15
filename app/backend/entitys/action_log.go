package entitys

import "time"

// 用户动作
type ActionLog struct {
	Id         int
	UserId     int                                     // 用户id
	Action     string    `orm:"size(20)"`                // 动作类型
	Actor      string    `orm:"size(20)"`                // 操作角色
	ObjectType string    `orm:"size(20)"`                // 操作对象类型
	ObjectId   int       `orm:"default(0)"`              // 操作对象id
	Extra      string    `orm:"size(1000)"`              // 额外信息
	Comment    string    `orm:"size(1000)"`              // 备注信息
	CreateTime time.Time `orm:"auto_now;type(datetime)"` // 更新时间
	Read       int
	Message    string    `orm:"-"` // 格式化后的消息
}

// 表名
func (m *ActionLog) TableName() string {
	return "action_log"
}