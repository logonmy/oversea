package frontend

type AboutController struct {
	FrontendBaseController
}

func (this *AboutController) Index()  {

	this.display()
}