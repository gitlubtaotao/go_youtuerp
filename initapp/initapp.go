package initapp

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"os"
	"time"
	"youtuerp/admin/middleware"
)

/*
 * @title: 处理request info
 */
func RequestInfo(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

/*
 * @title iris 日志输出的设置
 */
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
		
		Columns: true,
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
	app := iris.New()
	app.Use(RequestInfo)
	route := middleware.NewRoute(app)
	app.Use(LogConfig())
	route.DefaultRegister()
	
	return app
}

/*
 * 创建日志文件
 * 必须在main函数中进行操作
 */
func NewLogFile() *os.File {
	filename := todayFilename()
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	f, err := os.Create("./log/"+filename)
	if err != nil {
		panic(err)
	}
	return f
}

//根据日期获取文件名，文件日志以最常用的方式工作
//但这些只是好的命名方式。
func todayFilename() string {
	today := time.Now().Format("2006-01-08")
	return today + ".log"
}
