package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"youtuerp/database"
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
	//加载系统配置文件
	configEnv := flag.String("env", "development", "set env development or production")
	flag.Parse()
	err = initialize.InitConfig(*configEnv)
	if err != nil {
		app.Logger().Error(err)
		panic(err)
	}
	app.Logger().Info()
	//加载数据库操作
	err = new(database.DataBase).DefaultInit()
	if err != nil {
		app.Logger().Error(err)
		panic(err)
	}
	iris.RegisterOnInterrupt(func() {
		database.GetDBCon().Close()
	})
	config := iris.WithConfiguration(iris.YAML("../conf/iris.yaml"))
	
	_ = app.Run(iris.Addr(":8081"), config, iris.WithoutServerError(iris.ErrServerClosed))
	
}

/*
 * @title: 初始化数据库
 */
