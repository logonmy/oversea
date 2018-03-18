package entitys

type OzImmigrantNation struct {
	Id         int    `orm:"column(nation_id);auto" description:"国家或地区id"`
	NationName string `orm:"column(nation_name);size(30)" description:"国家或地区名称"`
	Flag       string `orm:"column(flag);size(500);null" description:"国旗图标"`
	Desc       string `orm:"column(desc);null" description:"简介描素"`
	Status     int8   `orm:"column(status);null" description:"状态，0正常 -1禁用"`
}

func (t *OzImmigrantNation) TableName() string {
	return "immigrant_nation"
}

