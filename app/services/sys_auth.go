package services

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"oversea/app/entitys"
	"errors"
	"oversea/utils"
	"oversea/app/stdout"
	"strings"
	"github.com/astaxie/beego"
	"strconv"
)

// 登录验证服务
type AuthService struct {
	loginUser *entitys.AdminUser    // 当前登录用户
}

// 用户登录
func (this *AuthService) Login(userName, password string) (string, error) {

	user, err := BackUserService.GetUserByName(userName)
	if err != nil {
		if err == orm.ErrNoRows {
			return "", errors.New(stdout.UserOrderPasswordError)
		} else {
			return "", errors.New(stdout.SystemError)
		}
	}

	if user.Password != utils.MD5(password+user.Salt) {
		return "", errors.New(stdout.UserOrderPasswordError)
	}
	if user.Status == -1 {
		return "", errors.New(stdout.UserIsDisenbled)
	}

	user.LastIp = utils.GetIpAddress()
	user.LastLogin = time.Now()
	BackUserService.UpdateAdminUser(user, "LastLogin", "LastIp")
	this.loginUser = user

	token := fmt.Sprintf("%d|%s", user.Id, utils.MD5(user.Password+user.Salt))
	return token, nil
}

// 检查是否登录
func (this *AuthService) IsLogined() bool {
	return this.loginUser != nil && this.loginUser.Id > 0
}

// 获取当前登录的用户对象
func (this *AuthService) GetUser() *entitys.AdminUser {
	return this.loginUser
}

// 获取当前登录的用户id
func (this *AuthService) GetUserId() int {
	if this.IsLogined() {
		return this.loginUser.Id
	}
	return 0
}

// 获取当前登录的用户名
func (this *AuthService) GetUserName() string {
	if this.IsLogined() {
		return this.loginUser.UserName
	}
	return ""
}

// 获取当前ip
func (this *AuthService) GetLastIp() string {
	if this.IsLogined() {
		return this.loginUser.LastIp
	}
	return ""
}

// 初始化
func (this *AuthService) Init(token string) {
	arr := strings.Split(token, "|")
	beego.Trace("登录验证, token: ", token)
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := BackUserService.GetUser(userId)
			if err == nil && password == utils.MD5(user.Password + user.Salt) {
				this.loginUser = user
				beego.Trace("验证成功，用户信息: ", user)
			}
		}
	}
}

// 退出登录
func (this *AuthService) Logout() error {
	this.loginUser.Id = 0
	return nil
}
