package main

import (
	"db"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "routers"
)

func main() {
	db.Init()
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT","GET","POST", "PATCH"},
		AllowHeaders: []string{"Origin, x-token"},
		ExposeHeaders: []string{"Content-Length, Access-Control-Allow-Origin, x-token"},
		AllowCredentials: false,    }))

	beego.Run()
}

