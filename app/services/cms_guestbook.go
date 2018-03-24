package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
)

type cmsGuestbookService struct {

}

//添加留言
func (this *cmsGuestbookService) AddCmsGuestbook(m *entitys.CmsGuestbook) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据id获得留言信息
func GetCmsGuestbookById(id int) (v *entitys.CmsGuestbook, err error) {
	o := orm.NewOrm()
	v = &entitys.CmsGuestbook{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新回复留言
func (this *cmsGuestbookService) UpdateCmsGuestbookById(m *entitys.CmsGuestbook) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsGuestbook{Id: m.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return
}

// 删除留言
func (this *cmsGuestbookService) DeleteCmsGuestbook(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsGuestbook{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CmsGuestbook{Id: id})
	}
	return
}


// 分页获取客户留言列表
func (this *cmsGuestbookService) GetAllCmsGuestbook(page, pageSize int,
	filters ...interface{}) ([]*entitys.CmsGuestbook,
	int64) {
	offset := (page - 1) * pageSize

	cmsGuestbooks := make([]*entitys.CmsGuestbook, 0)

	query := orm.NewOrm().QueryTable(new(entitys.CmsGuestbook))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&cmsGuestbooks)

	return cmsGuestbooks, total
}
