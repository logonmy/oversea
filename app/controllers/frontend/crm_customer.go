package frontend

import (
	"github.com/astaxie/beego/validation"
	"oversea/app/entitys"
	"oversea/app/stdout"
	"oversea/app/services"
)

type CrmCustomerController struct {
	FrontendBaseController
}

// 客户前端填写
func (this *CrmCustomerController) AddCustomer() {

	name := this.GetString("name")
	mobile := this.GetString("mobile")

	customer := entitys.CrmCustomer{Name:name, Mobile:mobile}
	valid := validation.Validation{}
	valid.Required(customer.Name, "name").Message(stdout.UsernameEmptyError)
	valid.Mobile(customer.Mobile, "mobile").Message(stdout.MobilePhoneError)

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			this.StdoutError(stdout.ParamsError, err.Message)
		}
	}

	_, err := services.CrmCustomerService.AddCrmCustomer(&customer)
	if err != nil {
		this.StdoutError(stdout.DBError, err.Error())
	}

	this.StdoutSuccess(nil)
}
