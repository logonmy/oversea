package backend

import (
	"oversea/app/services"
	"encoding/json"
	"oversea/app/entitys"
	"oversea/utils"
	"oversea/app/stdout"
)

type ContactController struct {
	AdminBaseController
}

// 添加沟通记录
func (this *ContactController) AddContactNote () {
	v := entitys.ActionLog{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		if utils.IsEmpty(v.Extra) {
			this.StdoutError(stdout.ParamsError, stdout.ParamsErrorMsg, nil)
		}
		services.SysActionLogService.AddContactLog(v.Customer, v.Contact, v.Extra)
	}
	this.StdoutSuccess(nil)
}