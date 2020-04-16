package models

import (
	"github.com/astaxie/beego/orm"
)

type DrawPlay struct {
	Id          int64  `json:"id,omitempty" orm:"auto"`
	DrawId      int64  `json:"drawId,omitempty" orm:"column(drawId)"`
	Play 		int  `json:"play" orm:"column(play)"`
	Deleted     int    `json:"deleted,omitempty" orm:"column(deleted)"`
	StartTime   int64  `json:"startTime" orm:"column(startTime)"`
	EndTime     int64  `json:"endTime" orm:"column(endTime)"`
}

func (d *DrawPlay) TableName() string {
	return "program_activity_draw_play"
}

func AddDrawPlay(play []DrawPlay, id int64) (num int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	for k, _ := range play {
		play[k].DrawId = id
		play[k].Play = k + 1
	}
	length := len(play)
	num, _ = o.InsertMulti(length, play)
	return
}

