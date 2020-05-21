package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/versioning"
	"net/http"
	"youtuerp/admin/controllers"
	"youtuerp/conf"
)

//处理路由信息
type IRoute interface {
	DefaultRegister()
	SessionRegister()
	OaRegister()
	CrmRegister()
	BaseDataRegister()
	SelectRegister()
	OtherRegister()
}



//
type Route struct {
	app *iris.Application
}

func (r *Route) BaseDataRegister() {
	j := r.jwtAccess()
	r.app.PartyFunc("/base/codes", func(c iris.Party) {
		record := controllers.BaseCodeController{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Post("/data", j.Serve, record.Get)
		c.Patch("/{id:uint}/update",j.Serve,record.Update)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
	r.app.PartyFunc("/base/ports", func(c iris.Party) {
		record := controllers.BasePort{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Post("/data", j.Serve, record.Get)
		c.Patch("/{id:uint}/update",j.Serve,record.Update)
		c.Delete("/{id:uint}/delete", j.Serve, record.Delete)
	})
	r.app.PartyFunc("/base/carriers", func(c iris.Party) {
		record := controllers.BaseCarrier{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Patch("/{id:uint}/update",j.Serve,record.Update)
		c.Post("/data", j.Serve, record.Get)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
	r.app.PartyFunc("/base/warehouses", func(c iris.Party) {
		record := controllers.BaseWarehouse{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Patch("/{id:uint}/update",j.Serve,record.Update)
		c.Post("/data", j.Serve, record.Get)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
}

func (r *Route) DefaultRegister() {
	r.SessionRegister()
	r.OaRegister()
	r.CrmRegister()
	r.SelectRegister()
	r.BaseDataRegister()
	r.OtherRegister()
}

func (r *Route) SessionRegister() {
	j := r.jwtAccess()
	session := controllers.SessionController{}
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
	numberSetting := controllers.NumberSettingController{}
	r.app.PartyFunc("/number_settings", func(c iris.Party) {
		c.Use(numberSetting.Before)
		c.Get("/column", j.Serve, numberSetting.GetColumn)
		c.Post("/create", j.Serve, numberSetting.Create)
		c.Post("/data", j.Serve, numberSetting.Get)
		c.Delete("/{id:uint}/delete", j.Serve, numberSetting.Delete)
	})
	
}
func (r *Route) CrmRegister() {
	j := r.jwtAccess()
	clue := controllers.CrmClue{}
	clueApi := r.app.Party("/crm/clues")
	{
		clueApi.Use(clue.Before)
		clueApi.Post("/data", j.Serve, clue.Get)
		clueApi.Get("/column", j.Serve, clue.GetColumn)
		clueApi.Post("/create", j.Serve, clue.Create)
		clueApi.Get("/{id: uint}/edit", j.Serve, clue.Edit)
		clueApi.Patch("/{id:uint}/update", j.Serve, clue.Update)
		clueApi.Delete("/{id:uint}/delete", j.Serve, clue.Delete)
		clueApi.Get("/{id:uint}/show", j.Serve, clue.Show)
	}
	track := controllers.CrmTrack{}
	trackApi := r.app.Party("/crm/tracks")
	{
		trackApi.Use(track.Before)
		trackApi.Post("/data", j.Serve, track.Get)
		trackApi.Post("/create", j.Serve, track.Create)
	}
	crmCompanyApi := r.app.Party("/crm/companies")
	{
		record := controllers.CrmCompany{}
		crmCompanyApi.Use(record.Before)
		crmCompanyApi.Post("/column", j.Serve, record.GetColumn)
		crmCompanyApi.Post("/create", j.Serve, record.Create)
		crmCompanyApi.Post("/data", j.Serve, record.Get)
		crmCompanyApi.Post("/create", j.Serve, record.Create)
		crmCompanyApi.Get("/{id:uint}/show", j.Serve, record.Show)
		crmCompanyApi.Get("/{id:uint}/edit", j.Serve, record.Edit)
		crmCompanyApi.Patch("/{id:uint}/update", j.Serve, record.Update)
		crmCompanyApi.Delete("/{id:uint}/delete", j.Serve, record.Delete)
		crmCompanyApi.Patch("/{id:uint}/changeStatus", j.Serve, record.ChangeStatus)
		crmCompanyApi.Patch("/{id:uint}/changeType", j.Serve, record.ChangeType)
	}
	crmUserApi := r.app.Party("/crm/users")
	{
		crmUser := controllers.CrmUser{}
		crmUserApi.Use(crmUser.Before)
		crmUserApi.Post("/column", j.Serve, crmUser.GetColumn)
		crmUserApi.Post("/data", j.Serve, crmUser.Get)
		crmUserApi.Post("/create", j.Serve, crmUser.Create)
		crmUserApi.Patch("/{id:uint}/update", j.Serve, crmUser.Update)
		crmUserApi.Delete("/{id:uint}/delete", j.Serve, crmUser.Delete)
	}
	InvoiceApi := r.app.Party("/invoices")
	{
		record := controllers.Invoice{}
		InvoiceApi.Use(record.Before)
		InvoiceApi.Post("/column", j.Serve, record.GetColumn)
		InvoiceApi.Post("/data", j.Serve, record.Get)
		InvoiceApi.Post("/create", j.Serve, record.Create)
		InvoiceApi.Patch("/{id:uint}/update", j.Serve, record.Update)
		InvoiceApi.Delete("/{id:uint}/delete", j.Serve, record.Delete)
	}
	AddressApi := r.app.Party("/address")
	{
		record := controllers.Address{}
		AddressApi.Use(record.Before)
		AddressApi.Post("/column", j.Serve, record.GetColumn)
		AddressApi.Post("/data", j.Serve, record.Get)
		AddressApi.Post("/create", j.Serve, record.Create)
		AddressApi.Patch("/{id:uint}/update", j.Serve, record.Update)
		AddressApi.Delete("/{id:uint}/delete", j.Serve, record.Delete)
	}
}

func (r *Route) SelectRegister() {
	j := r.jwtAccess()
	selectData := controllers.SelectController{}
	selectApi := r.app.Party("/select")
	{
		selectApi.Post("/companies", j.Serve, selectData.GetCompany)
		selectApi.Post("/base", j.Serve, selectData.GetCommon)
	}
}

func (r *Route) OtherRegister() {
	j := r.jwtAccess()
	uploader := controllers.UploadController{}
	setting := controllers.SettingController{}
	r.app.Post("/upload", j.Serve, uploader.Upload)
	r.app.Post("/setting/update_system", j.Serve, setting.UpdateSystem)
	r.app.Post("/setting/update_user", j.Serve, setting.UpdateUser)
	r.app.Post("/setting/data", j.Serve, setting.Get)
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

func NewRoute(app *iris.Application) IRoute {
	return &Route{app: app}
}
