package entitys



type OzProjectType struct {
	Id              int    `orm:"column(project_type_id);auto" description:"项目类型id"`
	ProjectTypeName string `orm:"column(project_type_name);size(30)" description:"项目类型名称"`
	Status          int8   `orm:"column(status);null" description:"状态，0正常 -1禁用"`
}

func (t *OzProjectType) TableName() string {
	return "project_type"
}
