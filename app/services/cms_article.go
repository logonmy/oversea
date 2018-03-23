package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"errors"
	"oversea/app/stdout"
)

type cmsArticleService struct {

} 

// 添加文章
func (this *cmsArticleService) AddCmsArticle(m *entitys.CmsArticle) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据id 获得文章信息
func (this *cmsArticleService) GetCmsArticleById(id int) (v *entitys.CmsArticle, err error) {
	o := orm.NewOrm()
	v = &entitys.CmsArticle{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 根据id 修改文章
func (this *cmsArticleService) UpdateCmsArticleById(m *entitys.CmsArticle, fileds ...string) (err error) {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}
	o := orm.NewOrm()
	v := entitys.CmsArticle{Id: m.Id}
	if err := o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
		return err
	}
	return nil
}

// 删除文章
func (this *cmsArticleService) DeleteCmsArticle(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CmsArticle{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CmsArticle{Id: id})
	}
	return
}


// 分页获取文章列表
func (this *cmsArticleService) GetCmsArticleList(page, pageSize int,
	filters ...interface{}) ([]*entitys.CmsArticle,
	int64) {
	offset := (page - 1) * pageSize

	cmsArticles := make([]*entitys.CmsArticle, 0)

	query := orm.NewOrm().QueryTable(new(entitys.CmsArticle))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&cmsArticles)

	return cmsArticles, total
}