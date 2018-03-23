package entitys

type CmsCategory struct {
	Id      int    `orm:"column(id);auto" description:"栏目id"`
	Pid     uint16 `orm:"column(pid)" description:"父id"`
	Type    int8   `orm:"column(type)" description:"类别 0:栏目 1:单网页"`
	Name    string `orm:"column(name);size(30)" description:"栏目名称"`
	Enname  string `orm:"column(enname);size(30)" description:"栏目英文名称"`
	Desc    string `orm:"column(desc)" description:"描述"`
	Url     string `orm:"column(url);size(100)" description:"链接地址"`
	Hits    uint   `orm:"column(hits)" description:"点击数量"`
	Setting string `orm:"column(setting)" description:"栏目配置"`
	Order   uint16 `orm:"column(order)" description:"排序"`
	IsMenu  uint8  `orm:"column(is_menu)" description:"是否显示，1 显示"`
}

func (t *CmsCategory) TableName() string {
	return "category"
}