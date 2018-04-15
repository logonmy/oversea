package entitys


type CrmCustomerSource struct {
	Id      int    `orm:"column(sid);auto"`
	Source  string `orm:"column(source);size(20)" description:"来源名称"`
	Comment string `orm:"column(comment);null" description:"备注"`
}

func (t *CrmCustomerSource) TableName() string {
	return "customer_source"
}
