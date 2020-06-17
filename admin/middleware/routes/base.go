package routes

import (
	"github.com/kataras/iris/v12"
	"youtuerp/admin/controllers"
)

type Base struct {
	Route *Route
}

func (b *Base) Index() {
	b.Carrier()
	b.Code()
	b.Port()
	b.Warehouse()
}
func (b *Base) Code() {
	r := b.Route
	j := r.jwtAccess()
	r.app.PartyFunc("/base/codes", func(c iris.Party) {
		record := controllers.BaseCodeController{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Post("/data", j.Serve, record.Get)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
	
}

func (b *Base) Port() {
	r := b.Route
	j := r.jwtAccess()
	r.app.PartyFunc("/base/ports", func(c iris.Party) {
		record := controllers.BasePort{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Post("/data", j.Serve, record.Get)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Delete("/{id:uint}/delete", j.Serve, record.Delete)
	})
}

func (b *Base) Carrier() {
	r := b.Route
	j := r.jwtAccess()
	r.app.PartyFunc("/base/carriers", func(c iris.Party) {
		record := controllers.BaseCarrier{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Post("/data", j.Serve, record.Get)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
}
func (b *Base) Warehouse() {
	r := b.Route
	j := r.jwtAccess()
	r.app.PartyFunc("/base/warehouses", func(c iris.Party) {
		record := controllers.BaseWarehouse{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Post("/data", j.Serve, record.Get)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
}

func NewRouteBase(r *Route) *Base {
	return &Base{r}
}
