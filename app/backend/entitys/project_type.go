package entitys



type projectType struct {
	Id         int  // 项目类型id
	PrjectTypeName   string    `orm:"unique;size(20)"`             // 项目类型名称
	Status     int       `orm:"default(0)"`                  // 状态，0正常 -1禁用
}

// 表名
func (m *projectType) TableName() string {
	return "project_type"
}

