package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"oversea/app/stdout"
	"errors"
)

type crmChanceService struct {
}

// 添加商机
func (this *crmChanceService) AddCrmChance(m *entitys.CrmChance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 获取商机详情
func (this *crmChanceService) GetCrmChanceById(id int) (v *entitys.CrmChance, err error) {
	o := orm.NewOrm()
	v = &entitys.CrmChance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新商机
func (this *crmChanceService) UpdateCrmChanceById(m *entitys.CrmChance, fileds ...string) (err error) {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}

	o := orm.NewOrm()
	v := entitys.CrmChance{Id: m.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
	}
	return
}

// 删除商机
func (this *crmChanceService) DeleteCrmChance(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CrmChance{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CrmChance{Id: id})
	}
	return
}

// 分页获取商机
func (this *crmChanceService) GetAllCrmChance(page, pageSize int,
	filters ...interface{}) ([]*entitys.CrmChance,
	int64) {
	offset := (page - 1) * pageSize

	crmChances := make([]*entitys.CrmChance, 0)

	query := orm.NewOrm().QueryTable(new(entitys.CrmChance))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&crmChances)

	return crmChances, total
}
