package initialize

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/sessions"
	"os"
	"time"
	"youtuerp/admin/controllers"
	"youtuerp/admin/middleware"
	"youtuerp/conf"
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
		Query:   true,
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
	//加载web端口对应的web secure cookie
	WebSecureCookie(app)
	route.DefaultRegister()
	conf.IrisApp = app
	RegisterView(app)
	app.OnErrorCode(iris.StatusNotFound, new(controllers.ErrorsController).NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, new(controllers.ErrorsController).InternalServerError)
	return app
}

/*
 * 创建日志文件
 * 必须在main函数中进行操作
 */
func NewLogFile(fileName string) (f *os.File, err error) {
	name := todayFilename(fileName)
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	return os.Create("./log/" + name)
}

func InitConfig(env string) error {
	return conf.NewSysConfig(env)
}

//i18n 国际化
func I18nInit(app *iris.Application) (err error) {
	err = app.I18n.Load("../locales/*/*", "en", "zh-CN")
	if err != nil {
		return
	}
	app.I18n.SetDefault("zh-CN")
	app.I18n.Subdomain = true
	app.I18n.URLParameter = "lang"
	return
}

//登录验证
//设置session对应的规则
func WebSecureCookie(app *iris.Application) {
	sessionName := conf.Configuration.SessionName
	hashKey := []byte("HDESS***SSS5EP0*")
	blockKey := []byte("AES-128")
	secureCookie := securecookie.New(hashKey, blockKey)
	mySessions := sessions.New(sessions.Config{
		Cookie:       sessionName,
		Encode:       secureCookie.Encode,
		Decode:       secureCookie.Decode,
		AllowReclaim: true,
		Expires:      -1,
	})
	app.Use(mySessions.Handler())
}

//进行前端页面的注册
func RegisterView(app *iris.Application) {
	tmpl := iris.HTML("./view", ".html")
	tmpl.Reload(true)
	tmpl.Layout("layouts/application.html")
	app.RegisterView(tmpl)
}

//根据日期获取文件名，文件日志以最常用的方式工作
//但这些只是好的命名方式。
func todayFilename(fileName string) string {
	today := time.Now().Format("2006-01-02")
	return today + "-" + fileName + ".log"
}
