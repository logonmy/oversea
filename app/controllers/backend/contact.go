package backend

import (
	"oversea/app/services"
	"encoding/json"
	"oversea/utils"
	"oversea/app/stdout"
	"oversea/app/form/backend"
)

type ContactController struct {
	AdminBaseController
}

// 添加沟通记录
func (this *ContactController) AddContactNote () {
	v := backend.ContactForm{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		if utils.IsEmpty(v.Content) {
			this.StdoutError(stdout.ParamsError, stdout.ParamsErrorMsg, nil)
		}
		services.SysActionLogService.AddContactLog(v.CustId, v.Id, v.Content)
	} else {
		this.StdoutError(stdout.ParamsError, err.Error(), nil)
	}
	this.StdoutSuccess(nil)
}

// 分页获取客户联系人列表
func (this *ContactController) GetAllContactNoteList() {

	contactForm := backend.ContactForm{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &contactForm)

	if contactForm.Page < 1 {
		contactForm.Page = 1
	}

	if contactForm.PageSize < 1 {
		contactForm.PageSize = this.pageSize
	}

	filters := make([]interface{}, 0)


	if contactForm.CustId > 0 {
		filters = append(filters, "customer", contactForm.CustId)
	}

	filters = append(filters, "action", "addcontact")
	notes, count := services.SysActionLogService.GetList(contactForm.Page, contactForm.PageSize, filters...)

	this.StdoutQuerySuccess(contactForm.Page, contactForm.PageSize, count, notes)

}