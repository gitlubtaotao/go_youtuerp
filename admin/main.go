package main

import (
	"encoding/json"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"os"
	"runtime/trace"
	"youtuerp/conf"
	"youtuerp/database"
	"youtuerp/initialize"
	"youtuerp/redis"
)

func main() {
	//加载系统配置文件
	app := NewAppInfo()
	f, err := initialize.NewLogFile("iris")
	if err != nil {
		app.Logger().Error(err)
		panic(err)
	}
	defer f.Close()
	app.Logger().AddOutput(f)
	if conf.Configuration.Env == "dev" {
		golog.SetLevel("debug")
	} else {
		golog.SetLevel("error")
	}
	golog.SetPrefix(conf.Configuration.Env + "-")
	config := iris.WithConfiguration(iris.YAML("../conf/iris.yaml"))
	_ = app.Run(iris.Addr(":8082"), config, iris.WithoutServerError(iris.ErrServerClosed))
}

//初始化app
func NewAppInfo() *iris.Application {
	configEnv := flag.String("env", "development", "set env development or production")
	flag.Parse()
	err := initialize.InitConfig(*configEnv)
	if err != nil {
		panic(err)
	}
	app := initialize.NewApp()
	err = new(database.DataBase).DefaultInit()
	//加载数据库操作
	if err != nil {
		app.Logger().Error(err)
		panic(err)
	}
	conf.ReisCon = redis.Connect()
	iris.RegisterOnInterrupt(func() {
		conf.ReisCon.Close()
	})
	//国际化翻译
	err = initialize.I18nInit(app)
	if err != nil {
		app.Logger().Error(err)
		panic(err)
	}
	return app
}

//Golang 性能测试 (3) 协程追踪术
func traceMethod() {
	tr, _ := os.Create("trace.out")
	defer tr.Close()
	_ = trace.Start(tr)
	defer trace.Stop()
}

func jsonOutput(l *golog.Log) bool {
	enc := json.NewEncoder(l.Logger.Printer)
	enc.SetIndent("", "    ")
	err := enc.Encode(l)
	return err == nil
}
