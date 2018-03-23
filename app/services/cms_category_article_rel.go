package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
)

type cmsCategoryArticleRelService struct {

}

func (this *cmsCategoryArticleRelService) AddCmsCategoryArticleRel(m *entitys.CmsCategoryArticleRel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}


func (this *cmsCategoryArticleRelService) GetCmsCategoryArticleRelById(id int) (v *entitys.CmsCategoryArticleRel, err error) {
	o := orm.NewOrm()
	v = &entitys.CmsCategoryArticleRel{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


func UpdateCmsCategoryArticleRelById(m *entitys.CmsCategoryArticleRel, fileds ...string) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsCategoryArticleRel{Id: m.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
	}
	return
}


func DeleteCmsCategoryArticleRel(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsCategoryArticleRel{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CmsCategoryArticleRel{Id: id})
	}
	return
}


// 分页获取文章列表
func (this *cmsArticleService) GetCmsCategoryArticleRelList(page, pageSize int,
	filters ...interface{}) ([]*entitys.CmsCategoryArticleRel,
	int64) {
	offset := (page - 1) * pageSize

	cmsCategoryArticleRels := make([]*entitys.CmsArticle, 0)

	query := orm.NewOrm().QueryTable(new(entitys.CmsArticle))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&cmsCategoryArticleRels)

	return cmsCategoryArticleRels, total
}