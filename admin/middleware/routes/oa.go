package routes

import (
	"github.com/kataras/iris/v12"
	"youtuerp/admin/controllers"
)

type Oa struct {
	Route *Route
}

func (o *Oa) Index() {
	o.numberSetting()
	o.companies()
	o.employees()
	o.accounts()
	o.departments()
}

func (o *Oa) companies() {
	company := controllers.CompanyController{}
	companyApi := o.Route.app.Party("/companies")
	j := o.Route.jwtAccess()
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
}

func (o *Oa) departments() {
	department := controllers.DepartmentController{}
	j := o.Route.jwtAccess()
	o.Route.app.PartyFunc("/departments", func(c iris.Party) {
		c.Use(department.Before)
		c.Get("/column", j.Serve, department.GetColumn)
		c.Post("/create", j.Serve, department.Create)
		c.Post("/data", j.Serve, department.Get)
		c.Patch("/{id:uint}/update", j.Serve, department.Update)
		c.Delete("/{id:uint}/delete", j.Serve, department.Delete)
	})
}
func (o *Oa) accounts() {
	account := controllers.AccountController{}
	j := o.Route.jwtAccess()
	o.Route.app.PartyFunc("/accounts", func(c iris.Party) {
		c.Use(account.Before)
		c.Get("/column", j.Serve, account.GetColumn)
		c.Post("/create", j.Serve, account.Create)
		c.Post("/data", j.Serve, account.Get)
		c.Get("/{id:uint}/edit", j.Serve, account.Edit)
		c.Patch("/{id:uint}/update", j.Serve, account.Update)
		c.Delete("/{id:uint}/delete", j.Serve, account.Delete)
	})
}
func (o *Oa) employees() {
	employee := controllers.EmployeeController{}
	j := o.Route.jwtAccess()
	o.Route.app.PartyFunc("/employees", func(c iris.Party) {
		c.Use(employee.Before)
		c.Get("/column", j.Serve, employee.GetColumn)
		c.Post("/create", j.Serve, employee.Create)
		c.Post("/data", j.Serve, employee.Get)
		c.Get("/{id:uint}/edit", j.Serve, employee.Edit)
		c.Patch("/{id:uint}/update", j.Serve, employee.Update)
		c.Delete("/{id:uint}/delete", j.Serve, employee.Delete)
	})
}
func (o *Oa) numberSetting() {
	numberSetting := controllers.NumberSettingController{}
	j := o.Route.jwtAccess()
	o.Route.app.PartyFunc("/number_settings", func(c iris.Party) {
		c.Use(numberSetting.Before)
		c.Get("/column", j.Serve, numberSetting.GetColumn)
		c.Post("/create", j.Serve, numberSetting.Create)
		c.Post("/data", j.Serve, numberSetting.Get)
		c.Delete("/{id:uint}/delete", j.Serve, numberSetting.Delete)
	})
}
func NewRouteOa(r *Route) *Oa {
	return &Oa{r}
}
