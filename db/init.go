package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"net/url"
	"oversea/app/backend/entitys"
	_"github.com/go-sql-driver/mysql"
)



func Init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")

	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(
		new(entitys.AdminUser),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}



