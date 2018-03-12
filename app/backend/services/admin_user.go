package services

import (
	"oversea/app/backend/entitys"
	"github.com/astaxie/beego/orm"
)

type AdminUserService struct{}

// 根据用户名获取用户信息
func (m *AdminUserService) GetUserByName(userName string) (*entitys.AdminUser, error) {
	o := orm.NewOrm()
	user := &entitys.AdminUser{}
	user.UserName = userName
	err := o.Read(user, "UserName")
	if err = o.QueryTable(new(entitys.AdminUser)).Filter("UserName", userName).One(user); err == nil {
		return user, nil
	}
	return user, err
}