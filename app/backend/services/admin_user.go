package services

import (
	"oversea/app/backend/entitys"
	"github.com/astaxie/beego/orm"
	"errors"
	"oversea/app/backend/stdout"
	"oversea/utils"
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

func (m *AdminUserService) UpdateAdminUser (user *entitys.AdminUser, fileds ...string) error {
	if len(fileds) < 1 {
		return errors.New(stdout.FieldsLengthMustMoreThanOne)
	}
	o := orm.NewOrm()
	v := entitys.AdminUser{Id:user.Id}
	if err := o.Read(&v); err == nil {
		_,err = o.Update(user, fileds...)
		return err
	}
	return nil
}

func (m *AdminUserService) AddUser(userName, email, password string, sex int) (*entitys.AdminUser, error) {
	if exists, _ := m.GetUserByName(userName); exists.Id > 0 {
		return nil, errors.New(stdout.UserIsExists)
	}
	o := orm.NewOrm()
	user := &entitys.AdminUser{}
	user.UserName = userName
	user.Sex = sex
	user.Email = email
	user.Salt = utils.NewNoDashUUID()
	user.Password = utils.MD5(password + user.Salt)
	// user.LastLogin = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	_, err := o.Insert(user)
	return user, err
}

// 根据用户id获取一个用户信息
func (m *AdminUserService) GetUser(userId int) (*entitys.AdminUser, error) {
	o := orm.NewOrm()
	user := &entitys.AdminUser{}
	user.Id = userId

	err := o.Read(user)

	return user, err
}

func (m *AdminUserService) GetAdminUsersList(page, pageSize int, filters ...interface{}) ([]*entitys.AdminUser, int64) {
	offset := (page - 1) * pageSize

	users := make([]*entitys.AdminUser, 0)

	u := &entitys.AdminUser{}
	query := orm.NewOrm().QueryTable(u.TableName())
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&users)

	return users, total
}