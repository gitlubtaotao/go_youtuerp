package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"youtuerp/admin/controllers"
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
	i.V1Register()
}

func (i *Route) MVCRegister() {
	mvc.New(i.app.Party("/")).Handle(&controllers.HomeController{})
	//公司信息
	mvc.New(i.app.Party("/company")).Handle(&controllers.CompanyController{Service: services.NewCompanyService()})
	//员工账户信息
	mvc.New(i.app.Party("/employee")).Handle(&controllers.EmployeeController{Service: services.NewEmployeeService()})
	//客户信息
	mvc.New(i.app.Party("/customer")).Handle(&controllers.CustomerController{Service: services.NewCrmCompanyService()})
	//供应商信息
	mvc.New(i.app.Party("/supplier")).Handle(&controllers.SupplierController{})
	
}

/*
 * @title: V1 api 路由注册方法
 * @description: api 路由注册方法
 */
func (i *Route) V1Register() {
	i.app.Get("/language", func(ctx iris.Context) {
		hi := ctx.Tr("hello_word")
		locale := ctx.GetLocale()
		ctx.Writef("From the language %s translated output: %s", locale.Language(), hi)
	})
	
}
