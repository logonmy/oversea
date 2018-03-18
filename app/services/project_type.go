package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
)

func AddOzProjectType(m *entitys.OzProjectType) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetOzProjectTypeById(id int) (v *entitys.OzProjectType, err error) {
	o := orm.NewOrm()
	v = &entitys.OzProjectType{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}