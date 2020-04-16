package models

import (
	"draw/lib"
	"github.com/astaxie/beego/orm"
)

type DrawPrize struct {
	Id          int64  `json:"id,omitempty" orm:"auto"`
	DrawId      int64   `json:"drawId,omitempty" orm:"column(drawId)"`
	PrizeAlias  string `json:"prizeAlias" orm:"column(prizeAlias)"`
	Level       int    `json:"level,omitempty" orm:"column(level)"`
	Name        string `json:"name" orm:"column(name)"`
	Num         int	   `json:"num" orm:"size(128)" orm:"column(num)"`
	Sum         int    `json:"sum" orm:"column(sum)"`
	Type        int    `json:"type,omitempty" orm:"column(type)"`
	TypeInfo    string `json:"typeInfo" orm:"column(typeInfo)"`
	Icon        string `json:"icon" orm:"column(icon)"`
	Deleted     int    `json:"deleted" orm:"column(deleted)"`
	WinningRate int    `json:"winningRate" orm:"column(winningRate)"`
	CreateTime  int64  `json:"createTime,omitempty" orm:"column(createTime)"`
	UpdateTime  int64  `json:"-" orm:"column(updateTime)"`
}

func (d *DrawPrize) TableName() string {
	return "program_activity_draw_prize"
}

func AddDrawPrize(prize []DrawPrize, id int64) (num int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	for k, _ := range prize {
		prize[k].DrawId = id
		prize[k].CreateTime = lib.GetCurrentTimeStamp()
		prize[k].UpdateTime = lib.GetCurrentTimeStamp()
	}
	length := len(prize)
	num, err = o.InsertMulti(length, prize)
	return
}

