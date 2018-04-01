package backend

import "oversea/app/services"

type CustomerController struct {
   AdminBaseController
}
func (this *CustomerController) List() {
	page, _ := this.GetInt("page")
	pageSize, _:= this.GetInt("pageSize")
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = this.pageSize
	}

	filters := make([]interface{}, 0)

	id := this.GetString("id")
	if id != "" {
		filters = append(filters, "id", id)
	}

	userList, count := services.CrmCustomerService.GetCrmCustomerList(page, pageSize, filters...)

	this.StdoutQuerySuccess(page, page, count, userList)

}