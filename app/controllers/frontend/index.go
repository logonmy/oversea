package frontend

type IndexController struct {
	FrontendBaseController
}

// 网站首页
func (this *IndexController) Home() {
	this.display()
}