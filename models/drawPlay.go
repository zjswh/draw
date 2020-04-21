package models

import "github.com/astaxie/beego/orm"

type DrawPlay struct {
	Id          int64  `json:"id,omitempty" orm:"auto"`
	DrawId      int64  `json:"drawId,omitempty" orm:"column(drawId)"`
	Play 		int  `json:"play" orm:"column(play)"`
	Deleted     int    `json:"deleted,omitempty" orm:"column(deleted)"`
	StartTime   int64  `json:"startTime" orm:"column(startTime)"`
	EndTime     int64  `json:"endTime" orm:"column(endTime)"`
	//Draw		*Draw `orm:"rel(fk)"`
}

func (d *DrawPlay) TableName() string {
	return "program_activity_draw_play"
}

func (m *DrawPlay) Update(fields ...string) (err error) {
	_, err = orm.NewOrm().Update(m, fields...)
	return
}

func (m *DrawPlay) Read(fields ...string) (err error) {
	err = orm.NewOrm().Read(m, fields...)
	return
}

func(m *DrawPlay) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}