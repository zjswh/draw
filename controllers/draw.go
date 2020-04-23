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
	c.Mapping("GetDrawContent", c.GetDrawContent)
	c.Mapping("GetDrawInfo", c.GetDrawInfo)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Login", c.Login)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Draw
// @Param	token					header 		string	true		"登录凭证"
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
	userInfo := c.CheckLogin()
	var draw models.Draw
	id, _ := c.GetInt64("id",0)
	draw.Uin = userInfo.Uin
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
// @router /GetDrawContent [get]
func (c *DrawController) GetDrawContent() {
	userInfo := c.CheckLogin()
	id, _ := c.GetInt64("id")
	draw := models.Draw{Id: id, Uin:userInfo.Uin}
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
// @Param	token		header	string	true	"登录凭证"
// @Param	num			query	int64	false	"每页显示数"
// @Param	page		query	int64	false	"页数"
// @Param	title		query	string	false	"标题"
// @Param	status		query	int64	false	"状态"
// @Param	sTime		query	string	false	"起始时间"
// @Param	eTime		query	string	false	"结束时间"
// @Success 200 {object} models.Draw
// @Failure 403
// @router /GetAll [get]
func (c *DrawController) GetAll() {
	userInfo := c.CheckLogin()
	page, _ := c.GetInt64("page", 1)
	num, _ := c.GetInt64("num", 10)
	status, _ := c.GetInt64("status", -1)
	sTime:= c.GetString("sTime", "")
	eTime:= c.GetString("eTime", "")
	title := c.GetString("title")
	var list []*models.Draw
	var draw models.Draw

	offset := (page - 1) * num

	query := draw.Query().Filter("uin", userInfo.Uin)
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

	//生成tokenkey   $token = strtoupper(md5(uniqid($time . $userInfo['uin'])));
	tokenKey := lib.UniqueId()
	expiredTime := 86400
	info, _ := json.Marshal(user)
	_, err = lib.RedisSetString(tokenKey,string(info), expiredTime)
	if err != nil {
		c.FormatJson("",3,err.Error())
	}
	var claim lib.Claims
	claim.Token = tokenKey
	claim.ExpiresAt = int64(expiredTime)
	token, _ := lib.CreateToken(claim)
	c.FormatJson(token, 0, "")
}

// getDrawInfo ...
// @Title getDrawInfo
// @Description get Draw by id
// @Param	id		query 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Draw
// @Failure 403 id is empty
// @router /GetDrawInfo [get]
func (c *DrawController) GetDrawInfo(){
	id, _ := c.GetInt64("id")
	key := "draw_info:" + string(id)
	info, err := lib.RedisGetString(key)
	if err != nil {
		c.FormatJson("",3,err.Error())
	}
	var draw models.Draw
	if info != "" {
		json.Unmarshal([]byte(info), &draw)
		c.FormatJson(draw, 0, "")
	}

	draw.Id = id
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
	infoString, _ := json.Marshal(draw)

	lib.RedisSetString(key,string(infoString),3600)
	c.FormatJson(draw, 0, "")
}



