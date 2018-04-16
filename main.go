package main

import (
	_ "oversea/routers"
	"github.com/astaxie/beego"
	"oversea/db"
	"oversea/app/services"
	_"github.com/go-sql-driver/mysql"
	_"github.com/astaxie/beego/session/mysql"
	"oversea/utils"
	"strconv"
	"github.com/astaxie/beego/orm"
)

const appEnv = EnvProd

const (
	EnvDev = "dev"
	EnvProd = "prod"
	EnvStg = "stg"
)

func main() {

	beego.LoadAppConfig("ini", "conf/" + appEnv + ".conf")

	if appEnv == EnvDev {
		orm.Debug = true
	}

    db.Init()
	services.InitServices()

	beego.AddFuncMap("static_url",static_url)
	beego.Run()
}

func static_url(url string) string {

	website := beego.AppConfig.String("website")
	//return website + url + "?v=" + strconv.FormatInt(time.Now().Unix(), 10)
	return website + url + "?v=" + strconv.FormatInt(utils.GetFileModTime("static/frontend/" + url), 10)
}