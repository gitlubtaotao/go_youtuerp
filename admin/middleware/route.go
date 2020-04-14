package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"net/http"
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
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	i.MVCRegister(crs)
	i.OtherRegister(crs)
}

func myAuthenticatedHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")
	foobar := user.Claims.(jwt.MapClaims)
	for key, value := range foobar {
		ctx.Writef("%s = %s", key, value)
	}
}

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

func (i *Route) jwtAccess() *jwt.Middleware {
	j := jwt.New(jwt.Config{
		// 通过 "token" URL参数提取。
		Extractor: jwt.FromParameter("token"),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Configuration.TokenSecret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return j
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
	}
}

//验证必须进行登录
func authRequired(ctx iris.Context) {
	seesionName := conf.Configuration.SessionName
	data := sessions.Get(ctx).Get(seesionName)
	if data != nil {
		conf.IrisApp.Logger().Infof("当前用户对于的session", data)
		ctx.Next()
	} else {
		ctx.Redirect("/login", http.StatusMovedPermanently)
	}
}
