package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"youtuerp/initialize"
)

func main() {
	app := initialize.NewApp()
	//将文件记录到log日志中
	f, err := initialize.NewLogFile("iris")
	if err != nil {
		app.Logger().Error(err)
	}
	defer f.Close()
	app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
	err = new(initialize.DataBase).DefaultInit()
	if err != nil {
		app.Logger().Error(err)
	}
	iris.RegisterOnInterrupt(func() {
		initialize.DataEngine.Close()
	})
	config := iris.WithConfiguration(iris.YAML("../conf/iris.yml"))
	_ = app.Run(iris.Addr(":8081"), config, iris.WithoutServerError(iris.ErrServerClosed))
	
}

/*
 * @title: 初始化数据库
 */
