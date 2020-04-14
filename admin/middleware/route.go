package middleware

import (
	"github.com/kataras/iris/v12"
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
	MVCRegister()
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
	i.MVCRegister()
	i.OtherRegister()
}

func (i *Route) MVCRegister() {
	requiredAuth := i.app.Party("/", authRequired)
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
func (i *Route) OtherRegister() {
	session := controllers.SessionController{}
	i.app.Get("/login", session.Get)
	i.app.Post("/login", session.Login)
	
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
