package services

import (
	"oversea/app/entitys"
	"github.com/astaxie/beego/orm"
	"errors"
	"oversea/app/stdout"
)

type ozImmigrantNationService struct {

}

// 添加国家地区
func (this *ozImmigrantNationService) AddOzImmigrantNation(m *entitys.OzImmigrantNation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据id获取国家地区
func (this *ozImmigrantNationService) GetOzImmigrantNationById(id int) (v *entitys.OzImmigrantNation, err error) {
	o := orm.NewOrm()
	v = &entitys.OzImmigrantNation{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新国家地区信息
func (this *ozImmigrantNationService) UpdateOzImmigrantNation(m *entitys.OzImmigrantNation, fileds ...string) error {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}
	o := orm.NewOrm()
	v := entitys.OzImmigrantNation{Id: m.Id}
	if err := o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
		return err
	}
	return nil
}

// 国家地区分页获取
func (this *ozImmigrantNationService) GetOzProjectTypeList(page, pageSize int,
	filters ...interface{}) ([]*entitys.OzImmigrantNation,
	int64) {
	offset := (page - 1) * pageSize

	immigrantNations := make([]*entitys.OzImmigrantNation, 0)

	query := orm.NewOrm().QueryTable(new(entitys.OzImmigrantNation))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&immigrantNations)

	return immigrantNations, total
}

// 获取所有的国家地区
func (this *ozImmigrantNationService) GetAllOzProjectTypeList() []*entitys.OzImmigrantNation {
	immigrantNations := make([]*entitys.OzImmigrantNation, 0)
	query := orm.NewOrm().QueryTable(new(entitys.OzImmigrantNation))
	query.OrderBy("-id").All(&immigrantNations)
	return immigrantNations
}