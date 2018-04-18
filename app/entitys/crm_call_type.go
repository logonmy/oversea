package entitys

type CrmCallType struct {
	Id   int    `orm:"column(id);pk"`
	Name string `orm:"column(name);size(50);null"`
}

func (t *CrmCallType) TableName() string {
	return "call_type"
}

