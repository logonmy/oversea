package backend

type LoginForm struct {
	UserName string `valid:"Required"`
	Password string `valid:"Required"`
	Remember string
}
