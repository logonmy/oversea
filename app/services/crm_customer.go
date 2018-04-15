package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"errors"
	"oversea/app/stdout"
	"fmt"
	"strings"
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
func (this *crmCustomerService) GetCrmCustomerList(page, pageSize int, sourceInArr []int,
	filters ...interface{}) ([]*entitys.CrmCustomerFilter,
	int64) {
	offset := (page - 1) * pageSize

	crmCustomers := make([]*entitys.CrmCustomerFilter, 0)

	qb, _ := orm.NewQueryBuilder("mysql")
	qb1, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象

	qb.Select("a.*",
		"b.source source_name", "c.user_name", "d.user_name creator").
		From("crm_customer a").
		LeftJoin("crm_customer_source b").On("a.source = b.sid").
		LeftJoin("sys_admin_user c").On("c.id = a.assign_to").
		LeftJoin("sys_admin_user d").On("d.id = a.create_by")

	qb1.Select("count(1) total").
		From("crm_customer a").
		LeftJoin("crm_customer_source b").On("a.source = b.sid").
		LeftJoin("sys_admin_user c").On("c.id = a.assign_to").
		LeftJoin("sys_admin_user d").On("d.id = a.create_by")

	filterValues := make([]interface{}, 0)

	var whereArr []string
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			if filters[k].(string) == "Director"{
				whereArr = append(whereArr, fmt.Sprintf("%s = ?", "d.user_name"))
			} else {
			   whereArr = append(whereArr, fmt.Sprintf("%s = ?", filters[k].(string)))
			}
			filterValues = append(filterValues, filters[k+1])

		}
	}


	var s []string
	for j := 0; j< len(sourceInArr); j++ {
		s = append(s, "?")
		filterValues = append(filterValues, sourceInArr[j])
	}

	if len(sourceInArr) > 0 {
		whereArr = append(whereArr, fmt.Sprintf("%s in(%s)", "a.source", strings.Join(s, ",")))
	}

	if len(whereArr) > 0 {
		qb.Where(strings.Join(whereArr, " AND "))
		qb1.Where(strings.Join(whereArr, " AND "))
	}

	countSql := qb1.String()

	var querytTable entitys.QuerytTable

	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(countSql, filterValues).QueryRow(&querytTable)

	qb.OrderBy("cust_id").Desc().
	Limit(pageSize).Offset(offset)
	sql := qb.String()
	o.Raw(sql, filterValues).QueryRows(&crmCustomers)

	return crmCustomers, querytTable.Total

}


func (this *crmCustomerService) GetAllCrmCustomerList() ([]*entitys.CrmCustomer) {

	crmCustomers := make([]*entitys.CrmCustomer, 0)
	query := orm.NewOrm().QueryTable(new(entitys.CrmCustomer))

	query.OrderBy("-id").All(&crmCustomers)

	return crmCustomers
}
