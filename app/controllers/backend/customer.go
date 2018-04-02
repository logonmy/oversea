package backend

import (
	"oversea/app/services"
	"oversea/app/form/backend"
	"encoding/json"
)

type CustomerController struct {
   AdminBaseController
}
func (this *CustomerController) List() {


	var customerForm backend.CustomerForm
	json.Unmarshal(this.Ctx.Input.RequestBody,&customerForm)

	if customerForm.Page < 1 {
		customerForm.Page = 1
	}

	if customerForm.PageSize < 1 {
		customerForm.PageSize = this.pageSize
	}

	filters := make([]interface{}, 0)

	id := customerForm.Id
	if id > 0 {
		filters = append(filters, "id", id)
	}

	if customerForm.AssignTo != `` {
		filters = append(filters, "assign_to", customerForm.AssignTo)
	}

	if customerForm.Name != `` {
		filters = append(filters, "name", customerForm.Name)
	}


	userList, count := services.CrmCustomerService.GetCrmCustomerList(customerForm.Page, customerForm.PageSize, filters...)

	this.StdoutQuerySuccess(customerForm.Page, customerForm.PageSize, count, userList)

}