package main

import (
	"draw/controllers"
	_ "draw/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}else{
		logs.SetLogger("file",`{"filename":"logs/test.log"}`)
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run("127.0.0.1:9001")
}
