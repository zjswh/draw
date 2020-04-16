package controllers

import (
	"draw/lib"
	"draw/models"
	"encoding/json"
	"strconv"
)

// DrawController operations for Draw
type DrawController struct {
	base
}

// URLMapping ...
func (c *DrawController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Draw
// @Param	title					formData 	string	true		"body for Draw content"
// @Param	intro					formData 	string	false		"body for Draw content"
// @Param	type					formData 	int		false		"body for Draw content"
// @Param	times					formData 	int		false		"body for Draw content"
// @Param	showResult				formData 	int		false		"body for Draw content"
// @Param	showRate				formData 	int		false		"body for Draw content"
// @Param	showType				formData 	int		false		"body for Draw content"
// @Param	condition				formData 	int		false		"body for Draw content"
// @Param	joinNum					formData 	int		false		"body for Draw content"
// @Param	prizeConfigs			formData 	string	false		"body for Draw content"
// @Param	playConfigs				formData 	string	false		"body for Draw content"
// @Success 201 {object} models.Draw
// @Failure 403 body is empty
// @router / [post]
func (c *DrawController) Post() {
	var draw models.Draw
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
	var playConfig []map[string]string
	var drawPlays []models.DrawPlay
	json.Unmarshal([]byte(playConfigs), &playConfig)

	//获取开始时间
	draw.StartTime = lib.GetTimeStamp(playConfig[0]["startTime"])
	playLen := len(playConfig)
	draw.EndTime = lib.GetTimeStamp(playConfig[playLen -1]["endTime"])

	var prizeConfig []models.DrawPrize

	json.Unmarshal([]byte(prizeConfigs), &prizeConfig)
	for _, v := range playConfig {
		var drawPlay models.DrawPlay
		drawPlay.Id, _ = strconv.ParseInt(v["id"], 10, 64)
		drawPlay.StartTime = lib.GetTimeStamp(v["startTime"])
		drawPlay.EndTime = lib.GetTimeStamp(v["endTime"])
		drawPlays = append(drawPlays, drawPlay)
	}
	id, _ := models.AddDraw(&draw, prizeConfig, drawPlays)
	c.FormatJson(id,0,"")
}

// GetOne ...
// @Title GetOne
// @Description get Draw by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Draw
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DrawController) GetOne() {
	id, _ := c.GetInt64(":id")
	info, _:= models.GetDrawById(id)
	c.FormatJson(info, 0, "")
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
// @router / [get]
func (c *DrawController) GetAll() {
	page, _ := c.GetInt64("page", 1)
	num, _ := c.GetInt64("num", 10)
	status, _ := c.GetInt64("status", -1)
	sTime:= c.GetString("sTime", "")
	eTime:= c.GetString("eTime", "")
	title := c.GetString("title")
	var createTime []int64
	if sTime != "" && eTime != "" {
		sTime := lib.GetTimeStamp(sTime)
		eTime := lib.GetTimeStamp(eTime)
		createTime = []int64{sTime,eTime}
	}

	list, _ := models.GetListDraw(title, status, createTime, page, num)
	c.FormatJson(list, 0, "")
}

// Put ...
// @Title Put
// @Description update the Draw
// @Param	id			path 	int64	true		"The id you want to update"
// @Param	body		body 	models.Draw	true		"body for Draw content"
// @Success 200 {object} models.Draw
// @Failure 403 :id is not int
// @router /:id [put]
func (c *DrawController) Put() {
	id, _ := c.GetInt64(":id")
	var draw models.Draw
	json.Unmarshal(c.Ctx.Input.RequestBody, &draw)
	draw.Id = id
	err := models.UpdateDrawById(&draw)
	if err != nil {
		c.FormatJson("success", 0, "")
	}
	c.FormatJson("", 2, err.Error())
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

