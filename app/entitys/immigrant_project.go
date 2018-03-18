package entitys

import "time"

type OzImmigrantProject struct {
	Id                   int       `orm:"column(project_id);auto" description:"项目id"`
	ProjectName          string    `orm:"column(project_name);size(30)" description:"项目名称"`
	NationId             int32     `orm:"column(nation_id)" description:"移民国家id"`
	ProjectTypeId        int32     `orm:"column(project_type_id)" description:"移民类型id"`
	InvestmentAmount     string    `orm:"column(investment_amount);size(50)" description:"投资金额"`
	Complexity           int8      `orm:"column(complexity);null" description:"手续复杂度，0-简单; 1-普通; 2-困难"`
	ResidencyRequirement string    `orm:"column(residency_requirement);size(50)" description:"居住要求"`
	JobRequirement       string    `orm:"column(job_requirement);size(50)" description:"工作要求"`
	CycleTime            string    `orm:"column(cycle_time);size(20)" description:"办理周期"`
	VisaType             string    `orm:"column(visa_type);size(20)" description:"签证类型"`
	ProjectDesc          string    `orm:"column(project_desc);null" description:"项目介绍"`
	ApplyRequirement     string    `orm:"column(apply_requirement);null" description:"申请条件描述"`
	PolicyAdvantage      string    `orm:"column(policy_advantage);null" description:"政策优势描述"`
	HandlingProcedures   string    `orm:"column(handling_procedures);null" description:"办理流程描述"`
	ApplyList            string    `orm:"column(apply_list);null" description:"申请材料清单"`
	ChargeList           string    `orm:"column(charge_list);null" description:"费用清单"`
	Status               int8      `orm:"column(status);null" description:"状态，0正常 -1禁用"`
	CreateTime           time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
	UpdateTime           time.Time `orm:"column(update_time);type(timestamp);auto_now"`
}

func (t *OzImmigrantProject) TableName() string {
	return "immigrant_project"
}