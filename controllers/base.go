package controllers

import (
	"draw/lib"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
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

func (c *base) GetUserInfo() {
	tokens := c.Ctx.Request.Header["Token"]
	if tokens == nil {
		c.FormatJson("",1,"未登录")
	}
	token := tokens[0]
	userInfo,ok := lib.ValidateToken(token)
	if !ok {
		c.FormatJson("",1,"鉴权失败")
	}
	c.Data["userInfo"] = userInfo
}

func (c *base) FormatJson(data interface{}, errorCode int, errorMessage string) {
	result := jsonResult{Code:200, ErrorCode:errorCode, Data:data, ErrorMessage:errorMessage}
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
	return
}

func CreateToken()(tokenss string,err error){
	//自定义claim
	claim := jwt.MapClaims{
		"id":       "aaa",
		"username": "sss",
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenss,err  = token.SignedString([]byte("GDY"))
	return
}

func secret()jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("GDY"),nil
	}
}

