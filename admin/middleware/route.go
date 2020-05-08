package middleware

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/versioning"
	"net/http"
	"strings"
	"youtuerp/admin/controllers"
	"youtuerp/conf"
)

//处理路由信息
type IRoute interface {
	DefaultRegister()
}

var allowMethods = []string{iris.MethodGet, iris.MethodPost, iris.MethodPatch,
	iris.MethodDelete, iris.MethodOptions}

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
func (r *Route) DefaultRegister() {
	r.SessionRegister()
	r.OaRegister()
	r.selectRegister()
	
}

func (r *Route) SessionRegister() {
	j := r.jwtAccess()
	session := controllers.SessionController{}
	fmt.Println(strings.Join(allowMethods, ","))
	users := r.app.Party("user/")
	{
		users.Post("/login", versioning.NewMatcher(versioning.Map{
			"1.0":               session.Login,
			versioning.NotFound: r.versionNotFound,
		}))
		users.Get("/info", j.Serve, session.Show)
		users.Post("/logout", j.Serve, session.Logout)
		users.Post("/update", j.Serve, session.Update)
		users.Post("/upload", j.Serve, session.UploadAvatar)
	}
}

func (r *Route) OaRegister() {
	j := r.jwtAccess()
	company := controllers.CompanyController{}
	companyApi := r.app.Party("/companies")
	{
		companyApi.Use(company.Before)
		companyApi.Post("/data", j.Serve, company.Get)
		companyApi.Get("/column", j.Serve, company.GetColumn)
		companyApi.Post("/create", j.Serve, company.Create)
		companyApi.Get("/{id:uint}/edit", j.Serve, company.Edit)
		companyApi.Patch("/{id:uint}/update", j.Serve, company.Update)
		companyApi.Delete("/{id:uint}/delete", j.Serve, company.Delete)
		companyApi.Get("/{id:uint}/show", j.Serve, company.Show)
	}
	department := controllers.DepartmentController{}
	r.app.PartyFunc("/departments", func(c iris.Party) {
		c.Use(department.Before)
		c.Get("/column", j.Serve, department.GetColumn)
		c.Post("/create", j.Serve, department.Create)
		c.Post("/data", j.Serve, department.Get)
		c.Patch("/{id:uint}/update", j.Serve, department.Update)
		c.Delete("/{id:uint}/delete", j.Serve, department.Delete)
	})
	employee := controllers.EmployeeController{}
	r.app.PartyFunc("/employees", func(c iris.Party) {
		c.Use(employee.Before)
		c.Get("/column", j.Serve, employee.GetColumn)
		c.Post("/create", j.Serve, employee.Create)
		c.Post("/data", j.Serve, employee.Get)
		c.Get("/{id:uint}/edit", j.Serve, employee.Edit)
		c.Patch("/{id:uint}/update", j.Serve, employee.Update)
		c.Delete("/{id:uint}/delete", j.Serve, employee.Delete)
	})
	account := controllers.AccountController{}
	r.app.PartyFunc("/accounts", func(c iris.Party) {
		c.Use(account.Before)
		c.Get("/column", j.Serve, account.GetColumn)
		c.Post("/create", j.Serve, account.Create)
		c.Post("/data", j.Serve, account.Get)
		c.Get("/{id:uint}/edit", j.Serve, account.Edit)
		c.Patch("/{id:uint}/update", j.Serve, account.Update)
		c.Delete("/{id:uint}/delete", j.Serve, account.Delete)
	})
	
}

func (r *Route) selectRegister() {
	j := r.jwtAccess()
	selectData := controllers.SelectController{}
	selectApi := r.app.Party("/select")
	{
		selectApi.Post("/companies", j.Serve, selectData.GetCompany)
	}
}

//验证jwt token
func (r *Route) jwtAccess() *jwt.Middleware {
	j := jwt.New(jwt.Config{
		// 通过 "token" URL参数提取。
		Extractor: jwt.FromAuthHeader,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Configuration.TokenSecret), nil
		},
		ErrorHandler: func(ctx context.Context, err error) {
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

func (r *Route) versionNotFound(ctx iris.Context) {
	ctx.StatusCode(404)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusNotFound})
}
