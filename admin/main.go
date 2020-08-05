package main

import (
	"encoding/json"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"os"
	"runtime/trace"
	"time"
	"youtuerp/conf"
	"youtuerp/global"
)

func main() {
	var err error
	////加载系统配置文件
	if err = NewAppInfo(); err != nil {
		panic(err)
	}
	if err = setupIrisLogger(); err != nil {
		panic(err)
	}
	config := iris.WithConfiguration(iris.YAML("../config/iris.yaml"))
	global.IrisAppEngine.Run(iris.Addr(":8082"), config, iris.WithoutServerError(iris.ErrServerClosed))
}

//初始化app
func NewAppInfo() error {
	global.NewIrisAppEngine()
	// loading router
	//routers.DefaultIrisRoute(global.IrisAppEngine)
	//加载数据库操作
	//国际化翻译
	if err := setupI18nEngine(); err != nil {
		return err
	}
	return nil
}

func init() {
	if err := setupSetting(); err != nil {
		panic(err)
	}
	if err := setupDBEngine(); err != nil {
		panic(err)
	}
	if err := global.NewRedisEngine(); err != nil {
		panic(err)
	}
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

func setupSetting() error {
	if err := global.SetupCommonSetting(); err != nil {
		return err
	}
	configEnv := flag.String("env", "development", "set env development or production")
	flag.Parse()
	return conf.NewSysConfig(*configEnv)
}

// set database engine
func setupDBEngine() error {
	if err := global.NewDBEngine(); err != nil {
		return err
	}
	if err := global.NewDBMigrate(); err != nil {
		return err
	}
	return nil
}

// set i18n engine
func setupI18nEngine() error {
	err := global.IrisAppEngine.I18n.Load("../locales/*/*", "en", "zh-CN")
	if err != nil {
		return err
	}
	global.IrisAppEngine.I18n.SetDefault("zh-CN")
	global.IrisAppEngine.I18n.Subdomain = true
	global.IrisAppEngine.I18n.URLParameter = "lang"
	return nil
}

// setup iris log
func setupIrisLogger() error {
	if conf.Configuration.Env == "dev" {
		global.IrisAppEngine.Logger().SetLevel("debug")
	} else {
		global.IrisAppEngine.Logger().SetLevel("error")
	}
	today := time.Now().Format("2006-01-02")
	name := today + "-" + "iris" + ".log"
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	f, err := os.Create("./log/" + name)
	if err != nil {
		return nil
	}
	global.IrisAppEngine.Logger().AddOutput(f)
	return nil
}
