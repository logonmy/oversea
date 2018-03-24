package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
)

type cmsUsersService struct {

}

// 添加新用户
func (this *cmsUsersService) AddCmsUsers(m *entitys.CmsUsers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据id获得用户信息
func (this *cmsUsersService) GetCmsUsersById(id int) (v *entitys.CmsUsers, err error) {
	o := orm.NewOrm()
	v = &entitys.CmsUsers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新用户信息
func (this *cmsUsersService) UpdateCmsUsersById(m *entitys.CmsUsers) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsUsers{Id: m.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return
}

// 删除用户
func (this *cmsUsersService) DeleteCmsUsers(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsUsers{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CmsUsers{Id: id})
	}
	return
}

// 分页获取客户列表
func (this *cmsUsersService) GetAllCmsUsers(page, pageSize int,
	filters ...interface{}) ([]*entitys.CmsUsers,
	int64) {
	offset := (page - 1) * pageSize

	cmsUsers := make([]*entitys.CmsUsers, 0)

	query := orm.NewOrm().QueryTable(new(entitys.CmsUsers))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&cmsUsers)

	return cmsUsers, total
}
