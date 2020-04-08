package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"youtuerp/initapp"
)

var Engine *xorm.Engine

func main() {
	app := initapp.NewApp()
	
	//将文件记录到log日志中
	f, err := initapp.NewLogFile("iris")
	if err != nil {
		app.Logger().Error(err)
	}
	defer f.Close()
	app.Logger().AddOutput(io.MultiWriter([]io.Writer{f, os.Stdout}...))
	err = initDataBase()
	if err != nil {
		app.Logger().Error(err)
	}
	iris.RegisterOnInterrupt(func() {
		Engine.Close()
	})
	config := iris.WithConfiguration(iris.YAML("../conf/iris.yml"))
	_ = app.Run(iris.Addr(":8081"), config, iris.WithoutServerError(iris.ErrServerClosed))
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
	f, err := initapp.NewLogFile("sql")
	if err != nil {
		return err
	}
	defer f.Close()
	Engine.SetLogger(log.NewSimpleLogger(f))
	return nil
}
