package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
)

type sysConfigService struct {

}

// 添加配置
func (this *sysConfigService) AddSysConfig(m *entitys.SysConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 批量添加配置
func (this *sysConfigService) AddSysMultiConfig(m []*entitys.SysConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.InsertMulti(len(m), m)
	return
}

// 根据分组获得配置
func (this *sysConfigService) GetSysConfigByConfigGroup(configGroup int) []*entitys.SysConfig {
	v := make([]*entitys.SysConfig, 0)
	query := orm.NewOrm().QueryTable(new(entitys.SysConfig))
	query.OrderBy("-id").All(&v)
	return v
}


// 获得配置
func (this *sysConfigService) GetSysConfigById(id int) (v *entitys.SysConfig, err error) {
	o := orm.NewOrm()
	v = &entitys.SysConfig{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// 根据id更新系统配置
func (this *sysConfigService) UpdateSysConfigById(m *entitys.SysConfig, fileds ...string) (err error) {
	o := orm.NewOrm()
	v := entitys.SysConfig{Id: m.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
	}
	return
}

func (this *sysConfigService) UpdateSysConfigByKey(m *entitys.SysConfig, fileds ...string) (err error) {
	o := orm.NewOrm()
	v := entitys.SysConfig{ConfigKey: m.ConfigKey}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
	}
	return
}

// 根据id 删除系统配置
func (this *sysConfigService) DeleteSysConfig(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.SysConfig{Id: id}
	if err = o.Read(&v); err == nil {
		_,err = o.Delete(&entitys.SysConfig{Id: id})
	}
	return
}
