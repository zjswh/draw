package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int64 `json:"id,omitempty"`
	Uin int64 `json:"uin,omitempty"`
	Password string `json:"password,omitempty"`
	Phone int64 `json:"phone,omitempty"`
	Name string `json:"name,omitempty"`
}

func (m *User) TableName() string {
	return "lps_soldier"
}

func (m *User) Update(fields ...string) (err error) {
	_, err = orm.NewOrm().Update(m, fields...)
	return
}

func (m *User) Read(fields ...string) (err error) {
	err = orm.NewOrm().Read(m, fields...)
	return
}

func(m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *User) Login(phone int64, password string) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(new(User)).Filter("phone",phone).Filter("Password",password).One(m)
	return
}
