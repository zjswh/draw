package controllers

import "github.com/astaxie/beego"

type base struct {
	beego.Controller
}

type jsonResult struct {
	Data interface{} `json:"data"`
	ErrorCode int `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Code int `json:"code"`
}

func (c *base) FormatJson(data interface{}, errorCode int, errorMessage string) {
	result := jsonResult{Code:200, ErrorCode:errorCode, Data:data, ErrorMessage:errorMessage}
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
	return
}
