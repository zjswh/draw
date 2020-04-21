package models

import (
	"draw/lib"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

type Draw struct {
	Id         int64 `json:"id,omitempty" orm:"auto"`
	Aid        int `json:"aid,omitempty"`
	Uin        int `json:"uin"`
	Title      string `json:"title,omitempty" orm:"size(128)"`
	Type       int `json:"type,omitempty"`
	Times      int `json:"times,omitempty"`
	Intro      string `json:"intro,omitempty" orm:"size(128)"`
	Status     int `json:"status"`
	ShowResult int `json:"showResult,omitempty" orm:"column(showResult)"`
	ShowType   int `json:"showType,omitempty" orm:"column(showType)"`
	NowPlay    int `json:"nowPlay,omitempty" orm:"column(nowPlay)"`
	TotalPlay  int `json:"totalPlay,omitempty" orm:"column(totalPlay)"`
	StartTime  int64 `json:"startTime" orm:"column(startTime)"`
	EndTime    int64 `json:"endTime" orm:"column(endTime)"`
	CreateTime int64 `json:"createTime" orm:"column(createTime)"`
	UpdateTime int64 `json:"-" orm:"column(updateTime)"`
	STaskId    int `json:"-" orm:"column(sTaskId)"`
	ETaskId    int `json:"-" orm:"column(eTaskId)"`
	Condition  int `json:"condition,omitempty" orm:"column(condition)"`
	CountDown  int `json:"countDown,omitempty" orm:"column(countDown)"`
	JoinNum    int `json:"joinNum,omitempty" orm:"column(joinNum)"`
	IsDeleted  int `json:"-" orm:"column(isDeleted)"`
	ShowRate   int `json:"showRate,omitempty" orm:"column(showRate)"`
	PrizeConfigs  []*DrawPrize `json:"prizeConfigs,omitempty" orm:"-"`
	PlayConfigs   []*DrawPlay `json:"playConfigs,omitempty" orm:"-"`
	PreviewUrl string `json:"previewUrl" orm:"-"`
}

type DrawInfo struct {
	*Draw
	PlayConfigs   []*DrawPlay `json:"playConfigs"`
}

func (d *Draw) TableName() string {
	return "program_activity_draw"
}

// AddDraw insert a new Draw into database and returns
// last inserted Id on success.
func AddDraw(m *Draw, prize []DrawPrize, play []DrawPlay) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	m.CreateTime = lib.GetCurrentTimeStamp()
	id, err = o.Insert(m)
	if err != nil {
		o.Rollback()
	}
	for k, _ := range prize {
		prize[k].DrawId = id
		prize[k].CreateTime = lib.GetCurrentTimeStamp()
		prize[k].UpdateTime = lib.GetCurrentTimeStamp()
	}

	length := len(prize)
	_, err = o.InsertMulti(length, prize)
	if err != nil {
		o.Rollback()
	}
	for k, _ := range play {
		play[k].DrawId = id
		play[k].Play = k + 1
	}
	length = len(play)
	_, err = o.InsertMulti(length, play)
	if err != nil {
		o.Rollback()
	}
	o.Commit()
	return
}

func UpdateDraw(m *Draw, prize []DrawPrize, play []DrawPlay) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	//o.Begin()
	v := Draw{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(&m); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	//o.QueryTable(new(Draw)).Filter("id",m.Id).Update("")
	//if err != nil {
	//	o.Rollback()
	//	return
	//}
	//currentTime := lib.GetCurrentTimeStamp()
	//for _, v := range prize {
	//	if v.Id == 0 { //新增
	//		v.DrawId = m.Id
	//		v.CreateTime = currentTime
	//		v.UpdateTime = currentTime
	//		o.Insert(&v)
	//	}else{  //更新
	//		o.ReadForUpdate(&v)
	//	}
	//}
	//
	//for k, v := range play {
	//	fmt.Println("v---",v.Id)
	//	if v.Id == 0 { //新增
	//		v.DrawId = m.Id
	//		v.Play = k + 1
	//		o.Insert(&v)
	//	}else{  //更新
	//		o.ReadForUpdate(&v)
	//	}
	//}
	//o.Commit()
	return
}

// GetDrawById retrieves Draw by Id. Returns error if
// Id doesn't exist
func GetDrawById(id int64) (v Draw, err error) {
	o := orm.NewOrm()
	v = Draw{Id: id}
	if err = o.QueryTable(new(Draw)).Filter("Id", id).RelatedSel().One(&v); err == nil {
		return
	}
	return
}

func GetListDraw(title string, status int64, createTime []int64, page int64, num int64) (ml []Draw, err error) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(Draw))
	if title != "" {
		qs = qs.Filter("title__contains",title)
	}

	if status != -1 {
		qs = qs.Filter("status",status)
	}

	if len(createTime) > 0 {
		qs = qs.Filter("createTime__gte",createTime[0]).Filter("createTime__lte",createTime[1])
	}

	offset := (page -1) * num
	qs.OrderBy("-id").Limit(num).Offset(offset).All(&ml)
	return
}


// GetAllDraw retrieves all Draw matches certain condition. Returns empty list if
// no records exist
func GetAllDraw(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Draw))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Draw
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

func countDraw()  {
	
}


// UpdateDraw updates Draw by Id and returns error if
// the record to be updated doesn't exist
func UpdateDrawById(m *Draw) ( err error){
	o := orm.NewOrm()
	v := Draw{Id: m.Id}
	//ascertain id exists in the database
	if err := o.ReadForUpdate(&v,"title"); err != nil {
		return err
	}

	//o.QueryTable(new(Draw)).Filter("id", m.Id).Update()

	return nil
}

func (m *Draw) Update() (err error) {
	localTime := lib.GetCurrentTimeStamp()
	o := orm.NewOrm()
	o.Using("default")
	o.Begin()
	err = o.Read(&Draw{Id:m.Id})
	if err != nil {
		return
	}
	m.UpdateTime = localTime
	_, err = o.Update(m)
	if err != nil {
		o.Rollback()
		return
	}

	prize := m.PrizeConfigs
	for _, v := range prize {
		v.UpdateTime = localTime
		if v.Id == 0 { //新增
			v.DrawId = m.Id
			v.CreateTime = localTime
			_, err = o.Insert(v)
		}else {  //修改
			_, err = o.Update(v,"prizeAlias","level","name","sum","typeInfo","icon","winningRate")
		}

		if err != nil {
			o.Rollback()
			return
		}
	}

	play := m.PlayConfigs
	for k, v := range play {
		if v.Id == 0 { //新增
			v.DrawId = m.Id
			v.Play = k + 1
			_, err = o.Insert(v)
		}else {  //修改
			_, err = o.Update(v,"startTime","endTime")
		}
		if err != nil {
			o.Rollback()
			return
		}

	}
	o.Commit()
	return
}

func (m *Draw) Read(fields ...string) (err error) {
	err = orm.NewOrm().Read(m, fields...)
	return
}

func(m *Draw) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func(m *Draw) Insert() (id int64, err error) {
	m.CreateTime = lib.GetCurrentTimeStamp()
	o := orm.NewOrm()
	o.Using("default")
	o.Begin()
	id, err = o.Insert(m)
	if err != nil {
		o.Rollback()
		return
	}

	prize := m.PrizeConfigs
	for k, _ := range prize {
		prize[k].DrawId = id
		prize[k].CreateTime = lib.GetCurrentTimeStamp()
		prize[k].UpdateTime = lib.GetCurrentTimeStamp()
	}
	length := len(prize)
	_, err = o.InsertMulti(length, prize)
	if err != nil {
		o.Rollback()
	}
	play := m.PlayConfigs
	for k, _ := range play {
		play[k].DrawId = id
		play[k].Play = k + 1
	}
	length = len(play)
	_, err = o.InsertMulti(length, play)
	if err != nil {
		o.Rollback()
	}
	o.Commit()
	return
}

// DeleteDraw deletes Draw by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDraw(id int64) (err error) {
	o := orm.NewOrm()
	v := Draw{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Draw{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
