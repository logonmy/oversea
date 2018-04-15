package backend

import "oversea/app/services"

type CrmCustomerSourceController struct {
	AdminBaseController
}

func (this *CrmCustomerSourceController) GetAllSource()  {

	userList := services.CrmCustomerSourceService.GetAllCrmCustomerSourceList()

	this.StdoutSuccess(userList)
}