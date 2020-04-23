package models

import (
	"github.com/astaxie/beego/orm"
)

type DrawResult struct {
	Id          int64  `json:"id,omitempty" orm:"auto"`
	Uin         int64  `json:"uin,omitempty" orm:"column(uin)"`
	DrawId      int64  `json:"drawId,omitempty" orm:"column(drawId)"`
	UserId  	int64  `json:"userId,omitempty" orm:"column(userId)"`
	Phone 		string `json:"phone,omitempty" orm:"column(phone)"`
	UserNick 	string `json:"userNick,omitempty" orm:"column(userNick)"`
	UserIp 		string `json:"userIp,omitempty" orm:"column(userIp)"`
	Avatar 		string `json:"avatar,omitempty" orm:"column(avatar)"`
	Status     	int    `json:"status" orm:"column(status)"`
	DrawPlay    int64   `json:"drawPlay,omitempty" orm:"column(drawPlay)"`
	PrizeId    	int    `json:"prizeId,omitempty" orm:"column(prizeId)"`
	PrizeLevel  int    `json:"prizeLevel,omitempty" orm:"column(prizeLevel)"`
	PrizeName   string `json:"prizeName,omitempty" orm:"column(prizeName)"`
	DrawTime    int    `json:"drawTime,omitempty" orm:"column(drawTime)"`
	Source      string `json:"source,omitempty" orm:"column(source)"`
	SourceId    int64  `json:"sourceId,omitempty" orm:"column(sourceId)"`
}

type DrawResultList struct {
	List []*DrawResult `json:"list"`
	Count int64 `json:"count"`
}

func (d *DrawResult) TableName() string {
	return "program_activity_draw_result"
}

func (m *DrawResult) Update(fields ...string) (err error) {
	_, err = orm.NewOrm().Update(m, fields...)
	return
}

func (m *DrawResult) Read(fields ...string) (err error) {
	err = orm.NewOrm().Read(m, fields...)
	return
}

func(m *DrawResult) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}