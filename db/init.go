package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"net/url"
	"oversea/app/entitys"
)

func Init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	tablePrefix := beego.AppConfig.String("db.prefix")

	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModelWithPrefix(tablePrefix,
		new(entitys.AdminUser),
		new(entitys.ActionLog),
	)

	orm.RegisterModelWithPrefix("oz",
		new(entitys.OzImmigrantProject),
			new(entitys.OzProjectType),
				new(entitys.OzImmigrantNation),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
