package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//dbConfig, err := config.NewConfig("ini", "conf/db.conf")
	user := beego.AppConfig.String("db::user")
	password := beego.AppConfig.String("db::password")
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	database := beego.AppConfig.String("db::database")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database)

	orm.RegisterDataBase("default", "mysql", dataSource)

	orm.RegisterModel(new(Draw), new(DrawPrize), new(DrawPlay), new(User))
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}

