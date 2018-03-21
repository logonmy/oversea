package entitys

type SysConfig struct {
	Id          int    `orm:"column(id);auto" description:"主键编号"`
	ConfigKey   string `orm:"column(config_key);size(60)" description:"配置名称"`
	ConfigValue string `orm:"column(config_value);size(255);null" description:"姓名"`
	ConfigGroup int `orm:"column(config_group);" description:"配置分组"`
}

func (t *SysConfig) TableName() string {
	return "config"
}
