package backend

type ContactForm struct {
	PageSize    int
	Page        int
	Id          int    // 联系人ID
	Content     string // 沟通内容
	ContactDate string // 联系时间
	NextDate    string // 下次联系时间
	CustId      int    // 客户ID
}
