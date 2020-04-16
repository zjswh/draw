package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	beego.Controller
}

type ErrorInfo struct {
	Code int
	Message string
	Data string
}

func (err *ErrorController) Error404() {
	logs.Error("api not found")
	err.Data["json"] = ErrorInfo{404,"api not found",""}
	err.ServeJSON()
}

func (err *ErrorController) Error501() {
	err.Data["json"] = ErrorInfo{501,"server error",""}
	err.ServeJSON()
}

func (err *ErrorController) ErrorDb() {
	err.Data["json"] = ErrorInfo{500,"database is now down",""}
	err.ServeJSON()
}
