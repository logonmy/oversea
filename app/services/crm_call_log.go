package services

import (
	"github.com/astaxie/beego/orm"
	"oversea/app/entitys"
	"oversea/app/stdout"
	"errors"
)

type crmCallLogService struct {

}

// 添加来电记录
func (this *crmCallLogService) AddCrmCallLog(m *entitys.CrmCallLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 获取来电详情
func (this *crmCallLogService) GetCrmCallLogById(id int) (v *entitys.CrmCallLog, err error) {
	o := orm.NewOrm()
	v = &entitys.CrmCallLog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// 更新来电记录
func (this *crmCallLogService) UpdateCrmCallLogById(m *entitys.CrmCallLog, fileds ...string) (err error) {

	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}

	o := orm.NewOrm()
	v := entitys.CrmCallLog{Id: m.Id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m, fileds...)
	}
	return
}

// 删除来电记录
func (this *crmCallLogService) DeleteCrmCallLog(id int) (err error) {
	o := orm.NewOrm()
	v := entitys.CrmCallLog{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&entitys.CrmCallLog{Id: id})
	}
	return
}

func (this *crmCallLogService) GetCrmCallLogList()  {

}