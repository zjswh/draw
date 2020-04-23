package controllers

import (
	"draw/lib"
	"draw/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type base struct {
	beego.Controller
}

type jsonResult struct {
	Data interface{} `json:"data"`
	ErrorCode int `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Code int `json:"code"`
}

type ListResult struct {
	List interface{} `json:"list"`
	Count int64 `json:"count"`
}

func (c *base) CheckLogin() (user models.User) {
	tokens := c.Ctx.Request.Header["Token"]
	if tokens == nil {
		c.FormatJson("",1,"未登录")
	}
	token := tokens[0]
	claim,ok := lib.ValidateToken(token)
	if !ok {
		c.FormatJson("",1,"登录已过期，请重新登陆")
	}
	userInfo, err := lib.RedisGetString(claim.Token)
	if err != nil{
		c.FormatJson("",3,err.Error())
	}
	json.Unmarshal([]byte(userInfo), &user)
	return
}

func (c *base) FormatJson(data interface{}, errorCode int, errorMessage string) {
	result := jsonResult{Code:200, ErrorCode:errorCode, Data:data, ErrorMessage:errorMessage}
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
	return
}
