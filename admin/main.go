package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"io"
	"log"
	"os"
	"youtuerp/initapp"
)

var DataEngine *gorm.DB

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
		DataEngine.Close()
	})
	config := iris.WithConfiguration(iris.YAML("../conf/iris.yml"))
	_ = app.Run(iris.Addr(":8081"), config, iris.WithoutServerError(iris.ErrServerClosed))
}

/*
 * @title: 初始化数据库
 */
func initDataBase() error {
	var err error
	DataEngine, err = gorm.Open("mysql", "root:qweqwe123@tcp(127.0.0.1:3306)/go_youtuerp?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true")
	if err != nil {
		return err
	}
	defer DataEngine.Close()
	DataEngine.DB().SetMaxOpenConns(1200)
	DataEngine.LogMode(true)
	DataEngine.SetLogger(log.New(os.Stdout, "\r\n", 0))
	return nil
}
