package backend

import (
	"oversea/app/form/backend"
	"encoding/json"
	"oversea/app/services"
)

type CrmChanceController struct {
	AdminBaseController
}


// 分页获取客户联系人列表
func (this *CrmChanceController) GetCustomerChanceList() {

	crmChanceForm := backend.New()
	json.Unmarshal(this.Ctx.Input.RequestBody, &crmChanceForm)


	filters := make([]interface{}, 0)


	if crmChanceForm.CustId > 0 {
		filters = append(filters, "customer", crmChanceForm.CustId)
	}

	notes, count := services.CrmChanceService.GetAllCrmChance(crmChanceForm.Page, crmChanceForm.PageSize, filters...)

	this.StdoutQuerySuccess(crmChanceForm.Page, crmChanceForm.PageSize, count, notes)

}