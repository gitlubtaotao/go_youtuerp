package middleware

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"strings"
	"youtuerp/admin/controllers"
	"youtuerp/conf"
	"youtuerp/services"
)

//处理路由信息
type IRoute interface {
	DefaultRegister()
	MVCRegister(crs context.Handler)
	OtherRegister(crs context.Handler)
}

//
type Route struct {
	app *iris.Application
}

func NewRoute(app *iris.Application) IRoute {
	return &Route{app: app}
}

/*
 * @title 路由的注册方法
 * @description 注册系统的方法
 */
func (i *Route) DefaultRegister() {
	allowedOrigins := strings.Split(conf.Configuration.AllowedOrigins, ",")
	crs := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	i.MVCRegister(crs)
	i.OtherRegister(crs)
}

//使用iris mvc 进行路由的注册
func (i *Route) MVCRegister(crs context.Handler) {
	j := i.jwtAccess()
	requiredAuth := i.app.Party("/", crs, j.Serve).AllowMethods(
		iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodDelete, iris.MethodOptions, )
	mvc.New(requiredAuth.Party("/")).Handle(&controllers.HomeController{})
	//公司信息
	mvc.New(requiredAuth.Party("/company")).Handle(&controllers.CompanyController{Service: services.NewCompanyService()})
	//员工账户信息
	mvc.New(requiredAuth.Party("/employee")).Handle(&controllers.EmployeeController{Service: services.NewEmployeeService()})
	//客户信息
	mvc.New(requiredAuth.Party("/customer")).Handle(&controllers.CustomerController{Service: services.NewCrmCompanyService()})
	//供应商信息
	mvc.New(requiredAuth.Party("/supplier")).Handle(&controllers.SupplierController{})
}

/*
 * @title: V1 api 路由注册方法
 * @description: api 路由注册方法
 */
func (i *Route) OtherRegister(crs context.Handler) {
	j := i.jwtAccess()
	session := controllers.SessionController{}
	users := i.app.Party("user/", crs).AllowMethods(
		iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodDelete, iris.MethodOptions, )
	{
		users.Post("/login", session.Login)
		users.Get("/info", j.Serve, session.Show)
		users.Post("/logout", j.Serve, session.Logout)
		users.Post("/update", j.Serve, session.Update)
	}
}

//验证jwt token
func (i *Route) jwtAccess() *jwt.Middleware {
	j := jwt.New(jwt.Config{
		// 通过 "token" URL参数提取。
		Extractor: jwt.FromParameter("token"),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Configuration.TokenSecret), nil
		},
		ErrorHandler: func(ctx context.Context, err error) {
			fmt.Printf(ctx.URLParam("token"))
			if err == nil {
				return
			}
			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			_, _ = ctx.JSON(iris.Map{"code": iris.StatusUnauthorized, "message": err.Error()})
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return j
}
