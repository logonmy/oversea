package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"errors"
	"oversea/app/stdout"
)

type crmCustomerService struct {
}

// 添加客户
func (this *crmCustomerService) AddCrmCustomer(m *entitys.CrmCustomer) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 获得客户详细信息
func (this *crmCustomerService) GetCrmCustomerById(id int) (v *entitys.CrmCustomer, err error) {
	o := orm.NewOrm()
	v = &entitys.CrmCustomer{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新客户信息
func (this *crmCustomerService) UpdateCrmCustomerById(m *entitys.CrmCustomer, fileds ...string) (err error) {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}
	o := orm.NewOrm()
	v := entitys.CrmCustomer{Id: m.Id}
	if err := o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
		return err
	}
	return nil
}

// 删除客户信息
func (this *crmCustomerService) DeleteCrmCustomer(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CrmCustomer{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CrmCustomer{Id: id})
	}
	return
}

// 分页获取客户列表
func (this *crmCustomerService) GetCrmCustomerList(page, pageSize int,
	filters ...interface{}) ([]*entitys.CrmCustomer,
	int64) {
	offset := (page - 1) * pageSize

	crmCustomers := make([]*entitys.CrmCustomer, 0)

	query := orm.NewOrm().QueryTable(new(entitys.CrmCustomer))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&crmCustomers)

	return crmCustomers, total
}


func (this *crmCustomerService) GetAllCrmCustomerList() ([]*entitys.CrmCustomer) {

	crmCustomers := make([]*entitys.CrmCustomer, 0)
	query := orm.NewOrm().QueryTable(new(entitys.CrmCustomer))

	query.OrderBy("-id").All(&crmCustomers)

	return crmCustomers
}
