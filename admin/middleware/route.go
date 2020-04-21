package middleware

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
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
	allowedOrigins := strings.Split(conf.Configuration.AllowedOrigins, ",")
	crs := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	r.SessionRegister(crs)
	r.OaRegister(crs)
}

func (r *Route) SessionRegister(crs context.Handler) {
	j := r.jwtAccess()
	session := controllers.SessionController{}
	users := r.app.Party("user/", crs).AllowMethods(
		iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodDelete, iris.MethodOptions)
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

func (r *Route) OaRegister(crs context.Handler) {
	j := r.jwtAccess()
	company := controllers.CompanyController{}
	companyApi := r.app.Party("/companies", crs).AllowMethods(
		iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodDelete, iris.MethodOptions)
	{
		companyApi.Post("/data", j.Serve, company.Get)
		companyApi.Get("/column", j.Serve, company.GetColumn)
		companyApi.Post("/", j.Serve, company.Create)
	}
}

//验证jwt token
func (r *Route) jwtAccess() *jwt.Middleware {
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

func (r *Route) versionNotFound(ctx iris.Context) {
	ctx.StatusCode(404)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusNotFound})
}
