package entitys

import "time"

type CmsArticle struct {
	Id             int       `orm:"column(id);auto"`
	Uid            int       `orm:"column(uid)" description:"管理员ID"`
	Title          string    `orm:"column(title);size(80)" description:"标题"`
	SubTitle       string    `orm:"column(sub_title);size(80)" description:"副标题"`
	Color          string    `orm:"column(color);size(24)" description:"标题颜色"`
	Font           string    `orm:"column(font);size(24)" description:"标题加粗"`
	Thumb          string    `orm:"column(thumb);size(255)" description:"图片地址"`
	Content        string    `orm:"column(content)" description:"内容"`
	CopyFrom       string    `orm:"column(copy_from);size(100)" description:"来源"`
	Keywords       string    `orm:"column(keywords);size(100)" description:"关键字"`
	Description    string    `orm:"column(description);size(250)" description:"描述"`
	Relation       string    `orm:"column(relation);size(255)" description:"相关文章"`
	PageType       int8      `orm:"column(page_type)" description:"分页方式"`
	MaxCharPerPage int32     `orm:"column(max_char_per_page)" description:"分页字符数"`
	Status         int8      `orm:"column(status)" description:"是否生效 0:失效 1:生效"`
	Hits           int8      `orm:"column(hits)" description:"点击数"`
	IsComment      int8      `orm:"column(is_comment)" description:"是否允许评论"`
	CreateTime     time.Time `orm:"column(create_time);type(timestamp);auto_now_add" description:"创建时间"`
	UpdateTime     time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新时间"`
}

func (t *CmsArticle) TableName() string {
	return "article"
}
