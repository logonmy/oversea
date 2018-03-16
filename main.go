package main

import (
	_ "oversea/routers"
	"github.com/astaxie/beego"
	"oversea/db"
	"oversea/app/backend/services"
	_"github.com/go-sql-driver/mysql"
	_"github.com/astaxie/beego/session/mysql"
)

func main() {
    db.Init()
	services.InitServices()
	beego.Run()
}

