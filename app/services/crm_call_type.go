package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"errors"
	"oversea/app/stdout"
)

type crmCallTypeService struct {

}

// 添加来电方式
func (this * crmCallTypeService) AddCrmCallType(m *entitys.CrmCallType) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 获取来电方式详情
func (this * crmCallTypeService) GetCrmCallTypeById(id int) (v *entitys.CrmCallType, err error) {
	o := orm.NewOrm()
	v = &entitys.CrmCallType{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 获取所有来电
func (this * crmCallTypeService) GetAllCallTypeList() []*entitys.CrmCallType {
	crmCallTypes := make([]*entitys.CrmCallType, 0)
	query := orm.NewOrm().QueryTable(new(entitys.CrmCallType))
	query.OrderBy("-id").All(&crmCallTypes)
	return crmCallTypes
}


// 更新来电方式
func (this * crmCallTypeService) UpdateCrmCallTypeById(m *entitys.CrmCallType,  fileds ...string) (err error) {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}

	o := orm.NewOrm()
	v := entitys.CrmCallType{Id: m.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return
}

// 删除来电方式
func (this * crmCallTypeService) DeleteCrmCallType(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CrmCallType{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CrmCallType{Id: id})
	}
	return
}
