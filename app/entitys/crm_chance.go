package entitys

import "time"

type CrmChance struct {
	Id                int       `orm:"column(id);pk" description:"编号"`
	ChanceTitle       string    `orm:"column(chance_title);size(255)" description:"机会主题"`
	CustId            int       `orm:"column(cust_id)" description:"客户ID"`
	FoundTime         time.Time `orm:"column(found_time);type(date)" description:"发现时间"`
	Demand            string    `orm:"column(demand);null" description:"客户需求"`
	EstimatedSignTime time.Time `orm:"column(estimated_sign_time);type(date)" description:"预计签单时间"`
	EstimatedCost     float64   `orm:"column(estimated_cost)" description:"预计金额"`
	Possable          string    `orm:"column(possable);size(255)" description:"可能性描素"`
	Status            int       `orm:"column(status)" description:"状态"`
	Creator           int       `orm:"column(creator)" description:"创建人"`
	Remark            string    `orm:"column(remark);size(255)" description:"备注"`
	LastContactTime   time.Time `orm:"column(last_contact_time);type(datetime);null" description:"最后联系时间"`
	Attachment        string    `orm:"column(attachment);size(255);null" description:"附件地址"`
	CreateTime        time.Time `orm:"column(create_time);type(timestamp);auto_now_add" description:"创建时间"`
	UpdateTime        time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
}

func (t *CrmChance) TableName() string {
	return "crm_chance"
}
