package db

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)


func NewConn() *redis.Client{

	client := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis.host") + ":" + beego.AppConfig.String("redis.port"),
		Password:  beego.AppConfig.String("redis.password"), // no password set
		DB:        0,  // use default DB
		MaxRetries: 1,
	})

	return client
}