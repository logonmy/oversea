package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"errors"
	"oversea/app/stdout"
)

type crmCustomerSourceService struct {

}

// 添加来源
func (this *crmCustomerSourceService) AddCrmCustomerSource(m *entitys.CrmCustomerSource) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 获得来源
func (this *crmCustomerSourceService) GetCrmCustomerSourceById(id int) (v *entitys.CrmCustomerSource, err error) {
	o := orm.NewOrm()
	v = &entitys.CrmCustomerSource{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


//更新来源
func (this *crmCustomerSourceService) UpdateCrmCustomerSourceById(m *entitys.CrmCustomerSource, fileds ...string) (err error) {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}

	o := orm.NewOrm()
	v := entitys.CrmCustomerSource{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
	}
	return
}

// 删除来源
func (this *crmCustomerSourceService) DeleteCrmCustomerSource(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CrmCustomerSource{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CrmCustomerSource{Id: id})
	}
	return
}


// 分页获取客户来源列表
func (this *crmCustomerSourceService) GetCrmCustomerSourceList(page, pageSize int,
	filters ...interface{}) ([]*entitys.CrmCustomerSource,
	int64) {
	offset := (page - 1) * pageSize

	rmCustomerSourcecs := make([]*entitys.CrmCustomerSource, 0)

	query := orm.NewOrm().QueryTable(new(entitys.CrmCustomerSource))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&rmCustomerSourcecs)

	return rmCustomerSourcecs, total
}


// 分页获取客户来源列表
func (this *crmCustomerSourceService) GetAllCrmCustomerSourceList() ([]*entitys.CrmCustomerSource) {

	rmCustomerSourcecs := make([]*entitys.CrmCustomerSource, 0)
	query := orm.NewOrm().QueryTable(new(entitys.CrmCustomerSource))
	query.OrderBy("-id").All(&rmCustomerSourcecs)

	return rmCustomerSourcecs
}

