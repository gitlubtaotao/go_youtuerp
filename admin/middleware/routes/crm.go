package routes

import "youtuerp/admin/controllers"

type Crm struct {
	Route *Route
}

func (c *Crm) Index() {
	c.company()
	c.user()
	c.address()
	c.clue()
	c.invoice()
	c.track()
}

func (c *Crm) track() {
	track := controllers.CrmTrack{}
	j := c.Route.jwtAccess()
	trackApi := c.Route.app.Party("/crm/tracks")
	{
		trackApi.Use(track.Before)
		trackApi.Post("/data", j.Serve, track.Get)
		trackApi.Post("/create", j.Serve, track.Create)
	}
}
func (c *Crm) user() {
	j := c.Route.jwtAccess()
	crmUserApi := c.Route.app.Party("/crm/users")
	{
		crmUser := controllers.CrmUser{}
		crmUserApi.Use(crmUser.Before)
		crmUserApi.Get("/column", j.Serve, crmUser.GetColumn)
		crmUserApi.Post("/data", j.Serve, crmUser.Get)
		crmUserApi.Post("/create", j.Serve, crmUser.Create)
		crmUserApi.Patch("/{id:uint}/update", j.Serve, crmUser.Update)
		crmUserApi.Delete("/{id:uint}/delete", j.Serve, crmUser.Delete)
	}
}
func (c *Crm) company() {
	j := c.Route.jwtAccess()
	crmCompanyApi := c.Route.app.Party("/crm/companies")
	{
		record := controllers.CrmCompany{}
		crmCompanyApi.Use(record.Before)
		crmCompanyApi.Get("/column", j.Serve, record.GetColumn)
		crmCompanyApi.Post("/create", j.Serve, record.Create)
		crmCompanyApi.Post("/data", j.Serve, record.Get)
		crmCompanyApi.Post("/create", j.Serve, record.Create)
		crmCompanyApi.Get("/{id:uint}/show", j.Serve, record.Show)
		crmCompanyApi.Get("/{id:uint}/edit", j.Serve, record.Edit)
		crmCompanyApi.Patch("/{id:uint}/update", j.Serve, record.Update)
		crmCompanyApi.Delete("/{id:uint}/delete", j.Serve, record.Delete)
		crmCompanyApi.Patch("/{id:uint}/changeStatus", j.Serve, record.ChangeStatus)
		crmCompanyApi.Patch("/{id:uint}/changeType", j.Serve, record.ChangeType)
		crmCompanyApi.Get("/{id:uint}/operationInfo", j.Serve, record.GetOperationInfo)
	}
}

func (c *Crm) invoice() {
	j := c.Route.jwtAccess()
	InvoiceApi := c.Route.app.Party("/invoices")
	{
		record := controllers.Invoice{}
		InvoiceApi.Use(record.Before)
		InvoiceApi.Get("/column", j.Serve, record.GetColumn)
		InvoiceApi.Post("/data", j.Serve, record.Get)
		InvoiceApi.Post("/create", j.Serve, record.Create)
		InvoiceApi.Patch("/{id:uint}/update", j.Serve, record.Update)
		InvoiceApi.Delete("/{id:uint}/delete", j.Serve, record.Delete)
	}
}
func (c *Crm) address() {
	j := c.Route.jwtAccess()
	AddressApi := c.Route.app.Party("/address")
	{
		record := controllers.Address{}
		AddressApi.Use(record.Before)
		AddressApi.Get("/column", j.Serve, record.GetColumn)
		AddressApi.Post("/data", j.Serve, record.Get)
		AddressApi.Post("/create", j.Serve, record.Create)
		AddressApi.Patch("/{id:uint}/update", j.Serve, record.Update)
		AddressApi.Delete("/{id:uint}/delete", j.Serve, record.Delete)
	}
}

func (c *Crm) clue() {
	r := c.Route
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
}

func NewRouteCrm(r *Route) *Crm {
	return &Crm{r}
}
