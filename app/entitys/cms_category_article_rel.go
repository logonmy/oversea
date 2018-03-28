package entitys

type CmsCategoryArticleRel struct {
	Id          int          `orm:"column(id);auto" description:"栏目id"`
	Cid         uint         `orm:"column(cid)" description:"栏目id category表相对应id"`
	Aid         uint         `orm:"column(aid)" description:"文章id"`
	Tid         uint         `orm:"column(tid)" description:"模板id"`
	Status      int8         `orm:"column(status)" description:"是否发布 0:不发布 1:发布"`
	IsTop       int8         `orm:"column(is_top)" description:"是否置顶 0:不置顶 1:置顶"`
	//CmsArticleData  *CmsArticle  `orm:"reverse(one)"` // 设置一对一反向关系(可选)
	//CmsCategoryData *CmsCategory `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

func (t *CmsCategoryArticleRel) TableName() string {
	return "category_article_rel"
}
