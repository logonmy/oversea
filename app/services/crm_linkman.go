package services

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"reflect"
	"oversea/app/entitys"
	"errors"
)

type crmLinkmanService struct {
}

// 添加联系人信息
func (this *crmLinkmanService) AddCrmLinkman(m *entitys.CrmLinkman) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 获取联系人信息
func (this *crmLinkmanService) GetCrmLinkmanById(id int) (v *entitys.CrmLinkman, err error) {
	o := orm.NewOrm()
	v = &entitys.CrmLinkman{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 查询联系人列表
func (this *crmLinkmanService) GetAllCrmLinkman(query map[string]string, fields []string, sortby []string,
	order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(entitys.CrmLinkman))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []entitys.CrmLinkman
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// 更新联系人信息
func (this *crmLinkmanService) UpdateCrmLinkmanById(m *entitys.CrmLinkman) (err error) {
	o := orm.NewOrm()
	v := entitys.CrmLinkman{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 删除联系人
func (this *crmLinkmanService) DeleteCrmLinkman(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CrmLinkman{Id: id}
	if err = o.Read(&v); err == nil {
		_,err = o.Delete(&entitys.CrmLinkman{Id: id})
	}
	return
}
