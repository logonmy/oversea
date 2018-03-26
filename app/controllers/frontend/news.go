package frontend

type NewsController struct {
	FrontendBaseController
}

func (this *NewsController) Index() {

	this.display()
}