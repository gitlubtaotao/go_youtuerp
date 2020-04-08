package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"youtuerp/admin/controllers"
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
	mvc.New(i.app.Party("/company")).Handle(&controllers.CompanyController{})
}

/*
 * @title: V1 api 路由注册方法
 * @description: api 路由注册方法
 */
func (i *Route) V1Register() {
}


