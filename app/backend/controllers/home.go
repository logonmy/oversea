package controllers


type HomeController struct {
	AdminBaseController
}

// 列表
func (this *HomeController) Index() {
	this.Data["pageTitle"] = "跳板机列表"
	this.display()
}