package frontend

type ContactController struct {
	FrontendBaseController
}

func (this *ContactController) Index() {
    this.display()
}