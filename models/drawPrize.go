package models

import "github.com/astaxie/beego/orm"

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


func (m *DrawPrize) Update(fields ...string) (err error) {
	_, err = orm.NewOrm().Update(m, fields...)
	return
}

func (m *DrawPrize) Read(fields ...string) (err error) {
	err = orm.NewOrm().Read(m, fields...)
	return
}

func(m *DrawPrize) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
