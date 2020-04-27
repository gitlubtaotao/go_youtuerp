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
		companyApi.Post("/data", j.Serve, company.Get)
		companyApi.Get("/column", j.Serve, company.GetColumn)
		companyApi.Post("/create", j.Serve, company.Create)
		companyApi.Get("/{id:uint}/edit", j.Serve, company.Edit)
		companyApi.Patch("/{id:uint}/update", j.Serve, company.Update)
		companyApi.Delete("/{id:uint}/delete", j.Serve, company.Delete)
	}
	department := controllers.DepartmentController{}
	r.app.PartyFunc("/departments", func(c iris.Party) {
		c.Use(department.Before)
		c.Get("/column", j.Serve, department.GetColumn)
		c.Post("/data", j.Serve, department.Get)
	})
}

func (r *Route) selectRegister() {
	j := r.jwtAccess()
	r.app.Post("/select/data", j.Serve, new(controllers.SelectController).Get)
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
