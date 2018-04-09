package backend

import (
	"oversea/app/services"
	"oversea/app/form/backend"
	"encoding/json"
	"oversea/app/stdout"
	"oversea/app/entitys"
	"strconv"
)

type CustomerController struct {
   AdminBaseController
}

// 获取客户列表
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

// 获取客户信息
func (this *CustomerController) GetInfo() {

	custId, _ := this.GetInt("id", 0)
	if custId <= 0 {
		this.StdoutError(stdout.ParamsError, stdout.ParamsErrorMsg, nil)
	}

	customer, err := services.CrmCustomerService.GetCrmCustomerById(custId)
	if err != nil {
		this.StdoutError(stdout.DBError, err.Error(), nil)
	}

	this.StdoutSuccess(customer)
}

// 添加客户资料
func (c *CustomerController) AddCrmCustomer() {
	var v entitys.CrmCustomer
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if len(v.Birthday) >= 10 {
			v.Birthday = v.Birthday[:10]
		}

		if custId, err := services.CrmCustomerService.AddCrmCustomer(&v); err == nil {
			services.CrmLinkmanService.AddCrmLinkman(&entitys.CrmLinkman{
				Phone:v.Mobile,
				Wechat:v.Wechat,
				Qq:v.Qq,
				Name:v.Name,
				Address:v.Address,
				Sex:v.Sex,
                Email:v.Email,
                CustId:custId,
			})
			c.Ctx.Output.SetStatus(201)
			c.StdoutSuccess(nil)
		} else {
			c.StdoutError(stdout.DBError, err.Error(), nil)
		}
	} else {
		c.StdoutError(stdout.DBError, err.Error(), nil)
	}
}


// 更新客户信息
func (c *CustomerController) UpdateCrmCustomerById() {
	v := entitys.CrmCustomer{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fields := c.checkFileds(v, []string{"CreateTime", "UpdateTime"})
		if len(v.Birthday) >= 10 {
			v.Birthday = v.Birthday[:10]
		}

		if err := services.CrmCustomerService.UpdateCrmCustomerById(&v, fields...); err == nil {
			c.StdoutSuccess(nil)
		} else {
			c.StdoutError(stdout.DBError, err.Error(), nil)
		}
	} else {
		c.StdoutError(stdout.ParamsError, err.Error(), nil)
	}
}

// 删除客户
func (c *CustomerController) DeleteCrmCustomer() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := services.CrmCustomerService.DeleteCrmCustomer(id); err == nil {
		c.StdoutSuccess(nil)
	} else {
		c.StdoutError(stdout.DBError, err.Error(), nil)
	}
}