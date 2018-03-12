package services

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"oversea/app/backend/entitys"
	"errors"
	"oversea/utils"
)

// 登录验证服务
type AuthService struct {
	loginUser *entitys.AdminUser    // 当前登录用户
	permMap   map[string]bool // 当前用户权限表
	openPerm  map[string]bool // 公开的权限
}

// 用户登录
func (this *AuthService) Login(userName, password string) (string, error) {
	adminUserService := &AdminUserService{}
	user, err := adminUserService.GetUserByName(userName)
	if err != nil {
		if err == orm.ErrNoRows {
			return "", errors.New("帐号或密码错误")
		} else {
			return "", errors.New("系统错误")
		}
	}

	if user.Password != utils.MD5(password+user.Salt) {
		return "", errors.New("帐号或密码错误")
	}
	if user.Status == -1 {
		return "", errors.New("该帐号已禁用")
	}

	user.LastLogin = time.Now()
	//UserService.UpdateUser(user, "LastLogin")
	//this.loginUser = user

	token := fmt.Sprintf("%d|%s", user.Id, utils.MD5(user.Password+user.Salt))
	return token, nil
}