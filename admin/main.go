package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"youtuerp/admin/middleware"
)

var Engine *xorm.Engine

func main() {
	app := middleware.NewApp()
	config := iris.WithConfiguration(iris.YAML("../conf/iris.yml"))
	err := initDataBase()
	if err != nil {
		app.Logger().Error(err)
	}
	_ = app.Run(iris.Addr(":8081"), config)
}

/*
 * @title: 初始化数据库
 */
func initDataBase() error {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:qweqwe123@tcp(127.0.0.1:3306)/youtuerp_development_1?charset=utf8&parseTime=true")
	if err != nil {
		return err
	}
	defer Engine.Close()
	Engine.DB().SetMaxOpenConns(1200)
	f, err := os.Create("./log/sql.log")
	if err != nil {
		return err
	}
	defer f.Close()
	Engine.SetLogger(log.NewSimpleLogger(f))
	return nil
}
