package controllers

import (
	"draw/lib"
	"draw/models"
	"encoding/json"
	"fmt"
)

// DrawController operations for Draw
type DrawController struct {
	base
}

const CRYPT_KEY = "ac8d51aj"

// URLMapping ...
func (c *DrawController) URLMapping() {
	c.Mapping("SaveConfig", c.SaveConfig)
	c.Mapping("GetInfo", c.GetInfo)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Login", c.Login)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Draw
// @Param	id						formData 	int64	false		"body for Draw content"
// @Param	title					formData 	string	true		"body for Draw content"
// @Param	intro					formData 	string	true		"body for Draw content"
// @Param	type					formData 	int		true		"body for Draw content"
// @Param	times					formData 	int		true		"body for Draw content"
// @Param	showResult				formData 	int		true		"body for Draw content"
// @Param	showRate				formData 	int		true		"body for Draw content"
// @Param	showType				formData 	int		true		"body for Draw content"
// @Param	condition				formData 	int		true		"body for Draw content"
// @Param	joinNum					formData 	int		true		"body for Draw content"
// @Param	prizeConfigs			formData 	string	false		"body for Draw content"
// @Param	playConfigs				formData 	string	false		"body for Draw content"
// @Success 201 {object} models.Draw
// @Failure 403 body is empty
// @router /SaveConfig [post]
func (c *DrawController) SaveConfig() {
	var draw models.Draw
	id, _ := c.GetInt64("id",0)
	draw.Type, _ = c.GetInt("type",1)
	draw.ShowResult, _ = c.GetInt("showResult",0)
	draw.ShowRate, _ = c.GetInt("showRate",0)
	draw.ShowType, _ = c.GetInt("showType",0)
	draw.Condition, _ = c.GetInt("condition",0)
	draw.JoinNum, _ = c.GetInt("joinNum",0)
	prizeConfigs := c.GetString("prizeConfigs")
	playConfigs := c.GetString("playConfigs")
	title := c.GetString("title")
	intro := c.GetString("intro")
	if title == "" || intro == "" || playConfigs == "" || prizeConfigs == ""  {
		c.FormatJson("",2,"参数缺失")
	}
	draw.Title = title
	draw.Intro = intro
	times , _ := c.GetInt("times",1)
	if draw.Type == 2 { //type为2 定时抽奖抽奖次数为1
		times = 1
	}
	draw.Times = times

	var PlayConfigs []*models.DrawPlay
	var PrizeConfigs []*models.DrawPrize

	json.Unmarshal([]byte(prizeConfigs), &PrizeConfigs)
	json.Unmarshal([]byte(playConfigs), &PlayConfigs)
	draw.PlayConfigs = PlayConfigs
	draw.PrizeConfigs = PrizeConfigs

	//获取开始时间
	draw.StartTime = draw.PlayConfigs[0].StartTime
	playLen := len(draw.PlayConfigs)
	draw.EndTime = draw.PlayConfigs[playLen -1].EndTime
	localTime := lib.GetCurrentTimeStamp()
	status := 1
	if draw.StartTime > localTime{
		status = 2
	}else if draw.EndTime < localTime {
		status = 0
	}
	draw.Status = status
	if id != 0 {
		draw.Id = id
		draw.Update()
	}else{
		id, _ = draw.Insert()
	}
	c.FormatJson(id,0,"")
}

// GetOne ...
// @Title GetOne
// @Description get Draw by id
// @Param	token	header 	string	true		"The key for staticblock"
// @Param	id		query 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Draw
// @Failure 403 :id is empty
// @router /getInfo [get]
func (c *DrawController) GetInfo() {
	c.GetUserInfo()
	fmt.Println(c.Data["userInfo"])
	id, _ := c.GetInt64("id")
	draw := models.Draw{Id: id}
	if err :=draw.Read(); err != nil{
		c.FormatJson("",9,"该抽奖不存在")
	}
	//获取场次
	var playConfigs []*models.DrawPlay
	var drawPlay models.DrawPlay
	drawPlay.Query().Filter("drawId", id).Filter("deleted",0).All(&playConfigs)
	draw.PlayConfigs = playConfigs

	//获取奖项列表
	var prizeConfigs []*models.DrawPrize
	var drawPrize models.DrawPrize
	drawPrize.Query().Filter("drawId", id).Filter("deleted",0).All(&prizeConfigs)
	draw.PrizeConfigs = prizeConfigs
	c.FormatJson(draw, 0, "")
}

// GetAll ...
// @Title GetAll
// @Description get Draw
// @Param	num		query	int64	false	"每页显示数"
// @Param	page	query	int64	false	"页数"
// @Param	title	query	string	false	"标题"
// @Param	status	query	int64	false	"状态"
// @Param	sTime	query	string	false	"起始时间"
// @Param	eTime	query	string	false	"结束时间"
// @Success 200 {object} models.Draw
// @Failure 403
// @router /GetAll [get]
func (c *DrawController) GetAll() {

	page, _ := c.GetInt64("page", 1)
	num, _ := c.GetInt64("num", 10)
	status, _ := c.GetInt64("status", -1)
	sTime:= c.GetString("sTime", "")
	eTime:= c.GetString("eTime", "")
	title := c.GetString("title")
	var list []*models.Draw
	var draw models.Draw

	offset := (page - 1) * num

	query := draw.Query()
	if title != "" {
		query = query.Filter("title__contains", title)
	}
	if status != -1 {
		query = query.Filter("status", status)
	}
	if sTime != "" {
		query = query.Filter("createTime__gte", lib.GetTimeStamp(sTime))
	}
	if eTime != "" {
		query = query.Filter("createTime__lte", lib.GetTimeStamp(eTime))
	}
	query.OrderBy("-id").Limit(num, offset).All(&list)
	if list != nil{
		for k, v := range list{
			list[k].PreviewUrl = fmt.Sprintf("http://web.guangdianyun.tv/live/%d/?uin=0", v.Id)
		}
	}

	count, _ := query.Count()
	listResult := ListResult{List:list, Count:count}
	c.FormatJson(listResult, 0, "")
}

// Put ...
// @Title Put
// @Description update the Draw
// @Param	id						formData 	int64	true		"The id you want to update"
// @Param	title					formData 	string	true		"body for Draw content"
// @Param	intro					formData 	string	true		"body for Draw content"
// @Param	type					formData 	int		true		"body for Draw content"
// @Param	times					formData 	int		true		"body for Draw content"
// @Param	showResult				formData 	int		true		"body for Draw content"
// @Param	showRate				formData 	int		true		"body for Draw content"
// @Param	showType				formData 	int		true		"body for Draw content"
// @Param	condition				formData 	int		true		"body for Draw content"
// @Param	joinNum					formData 	int		true		"body for Draw content"
// @Success 200 {object} models.Draw
// @Failure 403 :id is not int
// @router /Edit [put]
func (c *DrawController) Edit() {
	id, _ := c.GetInt64("id")
	draw := models.Draw{Id: id}
	if err := draw.Read(); err != nil {
		c.FormatJson("该抽奖不存在",9,"")
	}
	draw.Type, _ = c.GetInt("type",1)
	draw.ShowResult, _ = c.GetInt("showResult",0)
	draw.ShowRate, _ = c.GetInt("showRate",0)
	draw.ShowType, _ = c.GetInt("showType",0)
	draw.Condition, _ = c.GetInt("condition",0)
	draw.JoinNum, _ = c.GetInt("joinNum",0)
	title := c.GetString("title")
	intro := c.GetString("intro")
	if title == "" || intro == "" {
		c.FormatJson("",2,"参数缺失")
	}
	err := draw.Update()
	if err != nil {
		c.FormatJson("", 4, "更新失败，原因为：" + err.Error())
	}
	c.FormatJson("success", 0, "")
}

// Delete ...
// @Title Delete
// @Description delete the Draw
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DrawController) Delete() {

}

// Login ...
// @Title Login
// @Description delete the Draw
// @Param	phone		formData 	int64	true	"The id you want to delete"
// @Param	password	formData 	string	true	"The id you want to delete"
// @Success 200 {string}
// @Failure 403 id is empty
// @router /Login [POST]
func (c *DrawController) Login(){
	phone, _ := c.GetInt64("phone")
	password := c.GetString("password")
	if phone == 0 || password == "" {
		c.FormatJson("",2,"用户名或密码错误")
	}
	//var user models.User
	var user  models.User
	password = lib.Md5V(CRYPT_KEY + password)
	err := user.Login(phone, password)
	if err != nil {
		c.FormatJson("", 1, err.Error())
	}

	if user.Id == 0 {
		c.FormatJson("", 2, "账号或密码错误")
	}
	var claim lib.Claims
	claim.Id = user.Id
	claim.Uin = user.Uin
	claim.Name = user.Name
	token, _ := lib.CreateToken(claim)
	c.FormatJson(token, 0, "")
}



