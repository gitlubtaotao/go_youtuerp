package global

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/versioning"
	"strings"
	"youtuerp/conf"
)

var allowMethods = []string{iris.MethodGet, iris.MethodPost, iris.MethodPatch,
	iris.MethodDelete, iris.MethodOptions}

func NewIrisAppEngine() {
	IrisAppEngine = iris.New()
	IrisAppEngine.Use(RequestInfo)
	IrisAppEngine.Use(defaultVersion)
	IrisAppEngine.Use(LogConfig())
	IrisAppEngine.Use(setAllowedMethod())
	IrisAppEngine.AllowMethods(iris.MethodOptions)
}

func setAllowedMethod() context.Handler {
	allowedOrigins := strings.Split(conf.Configuration.AllowedOrigins, ",")
	crs := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   allowMethods,
		AllowCredentials: true,
		Debug:            true,
	})
	return crs
}

//@title: 处理request info
func RequestInfo(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func defaultVersion(ctx iris.Context) {
	ctx.Values().Set(versioning.Key, ctx.URLParamDefault("version", "1.0"))
	ctx.Next()
}

// @title iris 日志输出的设置
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
