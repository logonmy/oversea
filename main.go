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
	"flag"
	"fmt"
)

const appEnv = EnvDev

const (
	EnvDev = "dev"
	EnvProd = "prod"
	EnvStg = "stg"
)

func main() {
	var env string
	flag.StringVar(&env, "env", EnvDev, "env is stg")
	flag.Parse()
	fmt.Println("env ======> " , env)

	if env != EnvDev && env != EnvStg && env != EnvProd {
		fmt.Println("参数错误，env参数目前只支持 dev, prod, stg 三种配置。")
		return
	}

	beego.LoadAppConfig("ini", "conf/" + env + ".conf")

	if env == EnvDev {
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