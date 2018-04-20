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
		new(entitys.SysConfig),
	)

	orm.RegisterModelWithPrefix("oz_",
		new(entitys.OzImmigrantProject),
		new(entitys.OzProjectType),
		new(entitys.OzImmigrantNation),
	)

	orm.RegisterModelWithPrefix("crm_",
		new(entitys.CrmCustomer),
		new(entitys.CrmLinkman),
		new(entitys.CrmCustomerSource),
		new(entitys.CrmCallType),
		new(entitys.CrmCallLog),
		new(entitys.CrmChance),
	)

	orm.RegisterModelWithPrefix("cms_",
		new(entitys.CmsArticle),
		new(entitys.CmsCategory),
		new(entitys.CmsCategoryArticleRel),
		new(entitys.CmsUsers),
		new(entitys.CmsGuestbook),
	)

}

