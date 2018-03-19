package services

import (
	"oversea/app/entitys"
	"github.com/astaxie/beego/orm"
	"errors"
	"oversea/app/stdout"
)

type ozImmigrantProjectService struct {

}

// 添加项目
func (this *ozImmigrantProjectService) AddOzImmigrantProject(m *entitys.OzImmigrantProject) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据id获取项目信息
func (this *ozImmigrantProjectService) GetOzImmigrantProjectById(id int) (v *entitys.OzImmigrantProject, err error) {
	o := orm.NewOrm()
	v = &entitys.OzImmigrantProject{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新项目信息
func (this *ozImmigrantProjectService) UpdateOzImmigrantProject(m *entitys.OzImmigrantProject, fileds ...string) error {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}
	o := orm.NewOrm()
	v := entitys.OzImmigrantProject{Id: m.Id}
	if err := o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
		return err
	}
	return nil
}

// 项目分页获取
func (this *ozImmigrantProjectService) GetOzImmigrantProjectList(page, pageSize int,
	filters ...interface{}) ([]*entitys.OzImmigrantProject,
	int64) {
	offset := (page - 1) * pageSize

	immigrantProject := make([]*entitys.OzImmigrantProject, 0)

	query := orm.NewOrm().QueryTable(new(entitys.OzImmigrantProject))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&immigrantProject)

	return immigrantProject, total
}
