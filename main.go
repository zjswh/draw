package main

import (
	_ "draw/routers"
	"draw/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	logs.SetLogger("file",`{"filename":"logs/test.log"}`)
	beego.Run("127.0.0.1:9001")
}
