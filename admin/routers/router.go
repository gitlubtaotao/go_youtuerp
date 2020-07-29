package routers

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"net/http"
	"youtuerp/admin/api"
	"youtuerp/conf"
)

type Routers struct {
	app *iris.Application
}

func DefaultIrisRoute(app *iris.Application) {
	router := &Routers{app: app}
	router.DefaultRegister()
}

func (r *Routers) DefaultRegister() {
	r.selectRegister()
	r.otherRegister()
	newRouteSession(r).Index()
	newRouteBase(r).Index()
	newRouteOrder(r).Index()
	newRouteFinance(r).Index()
	newRouteCrm(r).Index()
	newRouteOa(r).Index()
	newAttachmentRoute(r).Index()
}

func (r *Routers) selectRegister() {
	j := r.jwtAccess()
	r.app.PartyFunc("/select", func(p iris.Party) {
		selectData := api.SelectApi{}
		p.Post("/companies", j.Serve, selectData.GetCompany)
		p.Post("/base", j.Serve, selectData.GetCommon)
		p.Get("/employee", j.Serve, selectData.Employee)
		p.Get("/owner_company", j.Serve, selectData.OwnerCompany)
		p.Get("/warehouse", j.Serve, selectData.WarehouseAddress)
		p.Get("/orderMaster", j.Serve, selectData.GetOrderMaster)
		p.Get("/baseCode", j.Serve, selectData.GetBaseCode)
		p.Get("/baseCarrier", j.Serve, selectData.GetCarrier)
		p.Get("/basePort", j.Serve, selectData.GetPort)
	})
}

func (r *Routers) otherRegister() {
	j := r.jwtAccess()
	uploader := api.UploadApi{}
	setting := api.Setting{}
	r.app.Post("/uploader", j.Serve, uploader.Upload)
	r.app.Post("/setting/update_system", j.Serve, setting.UpdateSystem)
	r.app.Post("/setting/update_user", j.Serve, setting.UpdateUser)
	r.app.Post("/setting/data", j.Serve, setting.Get)
}

//验证jwt token
func (r *Routers) jwtAccess() *jwt.Middleware {
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

func (r *Routers) versionNotFound(ctx iris.Context) {
	ctx.StatusCode(404)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusNotFound})
}
