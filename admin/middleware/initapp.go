package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"os"
	"time"
)

/*
 * @title: 处理request info
 */
func RequestInfo(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func LogConfig() iris.Handler {
	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	return customLogger
}

/*
 * @title: 初始化app
 */
func NewApp() *iris.Application {
	app := iris.Default()
	app.Use(RequestInfo)
	app.Use(LogConfig())
	route := NewRoute(app)
	route.DefaultRegister()
	return app
}

func NewLogFile() *os.File {
	filename := todayFilename()
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	f, err := os.OpenFile("./log/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

//根据日期获取文件名，文件日志以最常用的方式工作
//但这些只是好的命名方式。
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".log"
}
