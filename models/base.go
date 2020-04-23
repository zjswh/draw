package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//dbConfig, err := config.NewConfig("ini", "conf/db.conf")
	runMode := beego.BConfig.RunMode
	user := beego.AppConfig.String(runMode + "::mysql_user")
	password := beego.AppConfig.String(runMode + "::mysql_password")
	host := beego.AppConfig.String(runMode + "::mysql_host")
	port := beego.AppConfig.String(runMode + "::mysql_port")
	database := beego.AppConfig.String(runMode + "::mysql_database")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database)

	orm.RegisterDataBase("default", "mysql", dataSource)

	orm.RegisterModel(new(Draw), new(DrawPrize), new(DrawPlay), new(User), new(DrawResult))
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}

