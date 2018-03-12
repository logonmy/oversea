package main

import (
	_ "oversea/routers"
	"github.com/astaxie/beego"
	"oversea/db"
)

func main() {
    db.Init()
	beego.Run()
}

