package backend

import (
	"encoding/json"
	"strconv"
	"oversea/app/stdout"
	"oversea/app/services"
	"oversea/app/entitys"
	"oversea/app/form/backend"
	"github.com/astaxie/beego/logs"
)

// CrmLinkmanController operations for CrmLinkman
type CrmLinkmanController struct {
	AdminBaseController
}

// 添加联系人
func (c *CrmLinkmanController) AddCrmLinkman() {
	var v entitys.CrmLinkman
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.CustId == 0 {
			c.StdoutError(stdout.DBError, stdout.ParamsErrorMsg, nil)
		}
		if _, err := services.CrmLinkmanService.AddCrmLinkman(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.StdoutSuccess(nil)
		} else {
			c.StdoutError(stdout.DBError, err.Error(), nil)
		}
	} else {
		c.StdoutError(stdout.DBError, err.Error(), nil)
	}
}

// 获取联系人信息
func (c *CrmLinkmanController) GetCrmLinkmanById() {

	id, _ := c.GetInt("id")
	v, err := services.CrmLinkmanService.GetCrmLinkmanById(id)
	if err != nil {
		c.StdoutError(stdout.DBError, err.Error(), nil)
	} else {
		c.StdoutSuccess(v)
	}
}

// 分页获取客户联系人列表
func (this *CrmLinkmanController) GetAllCrmLinkmanList() {

	var linkmanForm backend.LinkmanForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &linkmanForm)

	if linkmanForm.Page < 1 {
		linkmanForm.Page = 1
	}

	if linkmanForm.PageSize < 1 {
		linkmanForm.PageSize = this.pageSize
	}

	filters := make([]interface{}, 0)

	id := linkmanForm.Id
	if id > 0 {
		filters = append(filters, "id", id)
	}

	if linkmanForm.CustId > 0 {
		filters = append(filters, "cust_id", linkmanForm.CustId)
	}

	if linkmanForm.Name != `` {
		filters = append(filters, "name", linkmanForm.Name)
	}

	linkmanList, count := services.CrmLinkmanService.GetCrmLinkmanList(linkmanForm.Page, linkmanForm.PageSize, filters...)

	this.StdoutQuerySuccess(linkmanForm.Page, linkmanForm.PageSize, count, linkmanList)

}

// 分页获取我的客户联系人列表
func (this *CrmLinkmanController) GetMyCrmLinkmanList() {

	var linkmanForm backend.LinkmanForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &linkmanForm)

	if linkmanForm.Page < 1 {
		linkmanForm.Page = 1
	}

	if linkmanForm.PageSize < 1 {
		linkmanForm.PageSize = this.pageSize
	}

	filters := make([]interface{}, 0)

	id := linkmanForm.Id
	if id > 0 {
		filters = append(filters, "id", id)
	}

	logs.Info(linkmanForm.CustId)
	logs.Info(string(this.Ctx.Input.RequestBody))
	logs.Info(linkmanForm)
	if linkmanForm.CustId > 0 {
		filters = append(filters, "cust_id", linkmanForm.CustId)
	} else {
		this.StdoutError(stdout.ParamsError, stdout.ParamsErrorMsg, nil)
	}

	if linkmanForm.Name != `` {
		filters = append(filters, "name", linkmanForm.Name)
	}

	linkmanList, count := services.CrmLinkmanService.GetCrmLinkmanList(linkmanForm.Page, linkmanForm.PageSize, filters...)

	this.StdoutQuerySuccess(linkmanForm.Page, linkmanForm.PageSize, count, linkmanList)

}

// 更新联系人信息
func (c *CrmLinkmanController) UpdateCrmLinkmanById() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := entitys.CrmLinkman{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fields := c.checkFileds(v, []string{"CreateTime", "UpdateTime"})
		if err := services.CrmLinkmanService.UpdateCrmLinkmanById(&v, fields...); err == nil {
			c.StdoutSuccess(nil)
		} else {
			c.StdoutError(stdout.DBError, err.Error(), nil)
		}
	} else {
		c.StdoutError(stdout.ParamsError, err.Error(), nil)
	}
}

// 删除联系人
func (c *CrmLinkmanController) DeleteCrmLinkman() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := services.CrmLinkmanService.DeleteCrmLinkman(id); err == nil {
		c.StdoutSuccess(nil)
	} else {
		c.StdoutError(stdout.DBError, err.Error(), nil)
	}
}
