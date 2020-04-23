package controllers

import (
	"draw/models"
	"github.com/astaxie/beego/orm"
)

// DrawController operations for Draw
type DrawResultController struct {
	base
}


// URLMapping ...
func (c *DrawResultController) URLMapping() {
	c.Mapping("PersonResult", c.PersonResult)
	c.Mapping("GetInfo", c.GetInfo)
	c.Mapping("WriteOff", c.WriteOff)
}

// personResult ...
// @Title personResult
// @Description get Draw by id
// @Param	userId		query 	int64	true		"The key for staticblock"
// @Param	id			query 	int64	true		"The key for staticblock"
// @Failure 403 id is empty
// @router /PersonResult [get]
func (c *DrawResultController) PersonResult() {
	id, _ := c.GetInt64("id")
	userId, _ := c.GetInt64("userId")
	//drawResult := models.DrawResult{DrawId:id}
	//err := drawResult.Read()
	var list []models.DrawResult
	var drawResult models.DrawResult
	_, err := drawResult.Query().Filter("drawId",id).Filter("userId",userId).All(&list)
	if err != nil {
		c.FormatJson("",2,err.Error())
	}
	result := make(map[int64][]models.DrawResult)
	if list != nil {
		for _, v := range list {
			result[v.DrawPlay] = append(result[v.DrawPlay], v)
		}
	}
	c.FormatJson(result,0,"")
}

// GetInfo ...
// @Title GetInfo
// @Description get Draw by id
// @Param	token		header 	string	true		"The key for staticblock"
// @Param	id			query 	int64	true		"The key for staticblock"
// @Param	userNick	query 	string	false	"The key for staticblock"
// @Param	prizeName	query 	string	false		"The key for staticblock"
// @Param	phone		query 	string	false		"The key for staticblock"
// @Param	isPrize		query 	int64	false		"The key for staticblock"
// @Param	sTime		query 	int64	false		"The key for staticblock"
// @Param	eTime		query 	int64	false		"The key for staticblock"
// @Failure 403 id is empty
// @router /GetInfo [get]
func (c *DrawResultController) GetInfo() {
	userInfo := c.CheckLogin()
	id, _ := c.GetInt64("id")
	userNick := c.GetString("userNick")
	prizeName := c.GetString("prizeName")
	phone := c.GetString("phone")
	isPrize, _ := c.GetInt("isPrize", -1)
	sTime, _ := c.GetInt64("sTime")
	eTime, _ := c.GetInt64("eTime")

	var list []*models.DrawResult
	var drawResult models.DrawResult
	query := drawResult.Query().Filter("drawId",id).Filter("uin",userInfo.Uin)

	if userNick != "" {
		query = query.Filter("userNick__contains",userNick)
	}

	if prizeName != "" {
		query = query.Filter("prizeName__contains",prizeName)
	}

	if phone != "" {
		query = query.Filter("phone__contains",phone)
	}

	if isPrize == 0 {
		query = query.Filter("prizeId",0)
	}else if isPrize == 1{
		query = query.Filter("prizeId__gt",0)
	}

	if sTime != 0 && eTime != 0 {
		query = query.Filter("drawTime__gte",sTime).Filter("drawTime__lte",eTime)
	}
	count, err := query.All(&list)
	if err != nil {
		c.FormatJson("",3,err.Error())
	}
	result := models.DrawResultList{List:list, Count:count}
	c.FormatJson(result,0,"")
}

// WriteOff ...
// @Title WriteOff
// @Description WriteOff
// @Param	token		header 		string	true		"The key for staticblock"
// @Param	recordId	formData 	int64	true		"The key for staticblock"
// @Param	userNick	formData 	string	false	"The key for staticblock"
// @Failure 403 id is empty
// @router /WriteOff [post]
func (c *DrawResultController) WriteOff(){
	userInfo := c.CheckLogin()
	recordId, _ := c.GetInt64("recordId")
	var drawResult models.DrawResult
	_, err := drawResult.Query().Filter("uin",userInfo.Uin).Filter("id",recordId).Update(orm.Params{
		"status" : 1,
	})

	if err != nil {
		c.FormatJson("",3,err.Error())
	}
	c.FormatJson("操作成功",0, "")
}


// @router /ExportDrawInfo [get]
func (c *DrawResultController) ExportDrawInfo() {
	//file := xlsx.NewFile()
	//sheet, _ := file.AddSheet("sheet1")
	//row := sheet.AddRow()
	//row.SetHeightCM(1)
	//cell := row.AddCell()
	//cell.Value = "haha"
	//cell = row.AddCell()
	//cell.Value = "xixi"
	//
	//err := file.Save("file.xlsx")
	//if err != nil {
	//	fmt.Println(err)
	//}

}
