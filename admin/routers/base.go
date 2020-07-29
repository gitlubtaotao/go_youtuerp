package routers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/admin/api"
)

type Base struct {
	route *Routers
}

func (b *Base) Index() {
	b.Carrier()
	b.Code()
	b.Port()
	b.Warehouse()
}
func (b *Base) Code() {
	r := b.route
	j := r.jwtAccess()
	r.app.PartyFunc("/base/codes", func(c iris.Party) {
		record := api.BaseCodeController{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Post("/data", j.Serve, record.Get)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})

}

func (b *Base) Port() {
	r := b.route
	j := r.jwtAccess()
	r.app.PartyFunc("/base/ports", func(c iris.Party) {
		record := api.BasePort{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Post("/data", j.Serve, record.Get)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Delete("/{id:uint}/delete", j.Serve, record.Delete)
	})
}

func (b *Base) Carrier() {
	r := b.route
	j := r.jwtAccess()
	r.app.PartyFunc("/base/carriers", func(c iris.Party) {
		record := api.BaseCarrier{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Post("/data", j.Serve, record.Get)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
}
func (b *Base) Warehouse() {
	r := b.route
	j := r.jwtAccess()
	r.app.PartyFunc("/base/warehouses", func(c iris.Party) {
		record := api.BaseWarehouse{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Post("/data", j.Serve, record.Get)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
}

func newRouteBase(r *Routers) *Base {
	return &Base{route: r}
}
