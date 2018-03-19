package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"errors"
	"oversea/app/stdout"
)

type ozProjectTypeService struct{}

// 添加项目类型
func (this *ozProjectTypeService) AddOzProjectType(m *entitys.OzProjectType) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据id获取项目类型信息
func (this *ozProjectTypeService) GetOzProjectTypeById(id int) (v *entitys.OzProjectType, err error) {
	o := orm.NewOrm()
	v = &entitys.OzProjectType{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新项目类型信息
func (this *ozProjectTypeService) UpdateOzProjectType(m *entitys.OzProjectType, fileds ...string) error {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}
	o := orm.NewOrm()
	v := entitys.OzProjectType{Id: m.Id}
	if err := o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
		return err
	}
	return nil
}

// 项目类型分页获取
func (this *ozProjectTypeService) GetOzProjectTypeList(page, pageSize int,
	filters ...interface{}) ([]*entitys.OzProjectType,
	int64) {
	offset := (page - 1) * pageSize

	projectType := make([]*entitys.OzProjectType, 0)

	query := orm.NewOrm().QueryTable(new(entitys.OzProjectType))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&projectType)

	return projectType, total
}

// 获取所有的项目类型
func (this *ozProjectTypeService) GetAllOzProjectTypeList() []*entitys.OzProjectType {
	projectTypes := make([]*entitys.OzProjectType, 0)
	query := orm.NewOrm().QueryTable(new(entitys.OzProjectType))
	query.OrderBy("-id").All(&projectTypes)
	return projectTypes
}