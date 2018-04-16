package backend

import (
	"oversea/app/services"
	"oversea/app/form/backend"
	"encoding/json"
	"oversea/app/stdout"
	"oversea/app/entitys"
	"strconv"
	"os"
	"github.com/astaxie/beego"
	"oversea/utils"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type CustomerController struct {
	AdminBaseController
}

// 获取客户列表
func (this *CustomerController) List() {

	var customerForm backend.CustomerForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &customerForm)

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

	if  !utils.IsEmpty(customerForm.AssignStatus) {
		filters = append(filters, "assign_status", customerForm.AssignStatus)
	}

	if  !utils.IsEmpty(customerForm.FollowStatus) {
		filters = append(filters, "follow_status", customerForm.FollowStatus)
	}


	if customerForm.Name != `` {
		filters = append(filters, "name", customerForm.Name)
	}

	if customerForm.Director != `` {
		filters = append(filters, "Director", customerForm.Director)
	}

	userList, count := services.CrmCustomerService.GetCrmCustomerList(customerForm.Page, customerForm.PageSize, customerForm.Source, filters...)

	this.StdoutQuerySuccess(customerForm.Page, customerForm.PageSize, count, userList)

}


// 获取我的客户列表
func (this *CustomerController) MyCustomerList() {

	var customerForm backend.CustomerForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &customerForm)

	if customerForm.Page < 1 {
		customerForm.Page = 1
	}

	if customerForm.PageSize < 1 {
		customerForm.PageSize = this.pageSize
	}

	filters := make([]interface{}, 0)

	filters = append(filters, "assign_to", this.userId)
	id := customerForm.Id
	if id > 0 {
		filters = append(filters, "id", id)
	}


	if  !utils.IsEmpty(customerForm.AssignStatus) {
		filters = append(filters, "assign_status", customerForm.AssignStatus)
	}

	if  !utils.IsEmpty(customerForm.FollowStatus) {
		filters = append(filters, "follow_status", customerForm.FollowStatus)
	}


	if customerForm.Name != `` {
		filters = append(filters, "name", customerForm.Name)
	}

	if customerForm.Director != `` {
		filters = append(filters, "Director", customerForm.Director)
	}

	userList, count := services.CrmCustomerService.GetCrmCustomerList(customerForm.Page, customerForm.PageSize, customerForm.Source, filters...)

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

		v.CreateBy = c.userId

		if custId, err := services.CrmCustomerService.AddCrmCustomer(&v); err == nil {
			services.CrmLinkmanService.AddCrmLinkman(&entitys.CrmLinkman{
				Phone:   v.Mobile,
				Wechat:  v.Wechat,
				Qq:      v.Qq,
				Name:    v.Name,
				Address: v.Address,
				Sex:     v.Sex,
				Email:   v.Email,
				CustId:  custId,
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

// 导出文件
func (c *CustomerController) Export() {

	xlsx := excelize.NewFile()
	// Create a new sheet.
	sheetName := "客户列表"
	index := xlsx.NewSheet(sheetName)
	xlsx.NewSheet("Sheet2")

	xlsx.SetCellValue(sheetName, "A1", "姓名")
	xlsx.SetCellValue(sheetName, "B1", "客户来源")
	xlsx.SetCellValue(sheetName, "C1", "联系电话")
	xlsx.SetCellValue(sheetName, "D1", "邮箱")
	xlsx.SetCellValue(sheetName, "E1", "QQ")
	xlsx.SetCellValue(sheetName, "F1", "客户资产")
	xlsx.SetCellValue(sheetName, "G1", "客户意向")
	xlsx.SetCellValue(sheetName, "H1", "创建时间")
	xlsx.SetCellValue(sheetName, "I1", "跟进时间")
	xlsx.SetCellValue(sheetName, "J1", "下次跟进时间")
	xlsx.SetCellValue(sheetName, "K1", "跟进状态")
	xlsx.SetCellValue(sheetName, "L1", "分配状态")

	contactStateMap := map[string]string {
		"0":"未分配",
		"1":"已分配",
		"2":"无需分配",
	}

	followStateMap := map[string]string {
		"0":"未跟进",
		"1":"已跟进",
		"2":"无需跟进",
	}

	resultList := services.CrmCustomerService.GetAllCrmCustomerList()
	for k, v := range resultList {
		keyIndex := strconv.Itoa(k + 2)
		xlsx.SetCellValue(sheetName, "A" + keyIndex, v.Name)
		xlsx.SetCellValue(sheetName, "B" + keyIndex, v.SourceName )
		xlsx.SetCellValue(sheetName, "C" + keyIndex, v.Mobile)
		xlsx.SetCellValue(sheetName, "D" + keyIndex, v.Email)
		xlsx.SetCellValue(sheetName, "E" + keyIndex, v.Qq)
		xlsx.SetCellValue(sheetName, "F" + keyIndex, v.Capital)
		xlsx.SetCellValue(sheetName, "G" + keyIndex, v.Intension)
		xlsx.SetCellValue(sheetName, "H" + keyIndex, v.CreateTime)
		xlsx.SetCellValue(sheetName, "I" + keyIndex, v.ContactedDate)
		xlsx.SetCellValue(sheetName, "J" + keyIndex, v.NextDate)
		xlsx.SetCellValue(sheetName, "K" + keyIndex, followStateMap[strconv.Itoa(v.FollowStatus)])
		xlsx.SetCellValue(sheetName, "L" + keyIndex, contactStateMap[strconv.Itoa(v.AssignStatus)])

	}

	xlsx.SetActiveSheet(index)

	path, _ := os.Getwd()
	path += "/static/frontend/excel/"
	filename := utils.NewUUID() + ".xlsx"
	err := xlsx.SaveAs(path + "/" + filename)
	if err != nil {
		c.StdoutError(stdout.HttpStatusFail, stdout.HttpErrorMsg, nil)
	}
	website := beego.AppConfig.String("frontend_website")
	c.StdoutSuccess(website + "excel/" + filename)
}


// 导出文件
func (c *CustomerController) MyCustomerExport() {

	xlsx := excelize.NewFile()
	// Create a new sheet.
	sheetName := "客户列表"
	index := xlsx.NewSheet(sheetName)
	xlsx.NewSheet("Sheet2")

	xlsx.SetCellValue(sheetName, "A1", "姓名")
	xlsx.SetCellValue(sheetName, "B1", "客户来源")
	xlsx.SetCellValue(sheetName, "C1", "联系电话")
	xlsx.SetCellValue(sheetName, "D1", "邮箱")
	xlsx.SetCellValue(sheetName, "E1", "QQ")
	xlsx.SetCellValue(sheetName, "F1", "客户资产")
	xlsx.SetCellValue(sheetName, "G1", "客户意向")
	xlsx.SetCellValue(sheetName, "H1", "创建时间")
	xlsx.SetCellValue(sheetName, "I1", "跟进时间")
	xlsx.SetCellValue(sheetName, "J1", "下次跟进时间")
	xlsx.SetCellValue(sheetName, "K1", "跟进状态")
	xlsx.SetCellValue(sheetName, "L1", "分配状态")

	contactStateMap := map[string]string {
		"0":"未分配",
		"1":"已分配",
		"2":"无需分配",
	}

	followStateMap := map[string]string {
		"0":"未跟进",
		"1":"已跟进",
		"2":"无需跟进",
	}

	resultList := services.CrmCustomerService.GetAllMyCrmCustomerList(c.userId)
	for k, v := range resultList {
		keyIndex := strconv.Itoa(k + 2)
		xlsx.SetCellValue(sheetName, "A" + keyIndex, v.Name)
		xlsx.SetCellValue(sheetName, "B" + keyIndex, v.SourceName )
		xlsx.SetCellValue(sheetName, "C" + keyIndex, v.Mobile)
		xlsx.SetCellValue(sheetName, "D" + keyIndex, v.Email)
		xlsx.SetCellValue(sheetName, "E" + keyIndex, v.Qq)
		xlsx.SetCellValue(sheetName, "F" + keyIndex, v.Capital)
		xlsx.SetCellValue(sheetName, "G" + keyIndex, v.Intension)
		xlsx.SetCellValue(sheetName, "H" + keyIndex, v.CreateTime)
		xlsx.SetCellValue(sheetName, "I" + keyIndex, v.ContactedDate)
		xlsx.SetCellValue(sheetName, "J" + keyIndex, v.NextDate)
		xlsx.SetCellValue(sheetName, "K" + keyIndex, followStateMap[strconv.Itoa(v.FollowStatus)])
		xlsx.SetCellValue(sheetName, "L" + keyIndex, contactStateMap[strconv.Itoa(v.AssignStatus)])

	}

	xlsx.SetActiveSheet(index)

	path, _ := os.Getwd()
	path += "/static/frontend/excel/"
	filename := utils.NewUUID() + ".xlsx"
	err := xlsx.SaveAs(path + "/" + filename)
	if err != nil {
		c.StdoutError(stdout.HttpStatusFail, stdout.HttpErrorMsg, nil)
	}
	website := beego.AppConfig.String("frontend_website")
	c.StdoutSuccess(website + "excel/" + filename)
}


func (c *CustomerController) AssignTo()  {
	var v entitys.CrmCustomer
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Id <= 0 {
			c.StdoutError(stdout.ParamsError, stdout.ParamsErrorMsg, nil)
		}
		if v.AssignTo <= 0 {
			c.StdoutError(stdout.ParamsError, stdout.ParamsErrorMsg, nil)
		}
		v.AssignStatus = 1
		err = services.CrmCustomerService.UpdateCrmCustomerById(&v, "AssignTo", "AssignStatus")
		c.checkError(err)
		c.StdoutSuccess(nil)
	} else {
		c.StdoutError(stdout.ParamsError, err.Error(), nil)
	}

}