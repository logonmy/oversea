package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
)

// 添加分类
func AddCmsCategory(m *entitys.CmsCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据id 获得分类
func GetCmsCategoryById(id int) (v *entitys.CmsCategory, err error) {
	o := orm.NewOrm()
	v = &entitys.CmsCategory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新分类
func UpdateCmsCategoryById(m *entitys.CmsCategory) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsCategory{Id: m.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return
}

// 删除分类
func DeleteCmsCategory(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsCategory{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CmsCategory{Id: id})
	}
	return
}


// 分页获取文章分类列表
func (this *crmCustomerService) GetCmsCategoryList(page, pageSize int,
	filters ...interface{}) ([]*entitys.CmsCategory,
	int64) {
	offset := (page - 1) * pageSize

	cmsCategorys := make([]*entitys.CmsCategory, 0)

	query := orm.NewOrm().QueryTable(new(entitys.CmsCategory))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&cmsCategorys)

	return cmsCategorys, total
}