package routes

import (
	"github.com/kataras/iris/v12"
	"youtuerp/admin/controllers"
)

type OrderMaster struct {
	Route *Route
}

func (o *OrderMaster) Index() {
	r := o.Route
	j := r.jwtAccess()
	r.app.PartyFunc("/order/masters", func(p iris.Party) {
		record := controllers.OrderMaster{}
		p.Use(record.Before)
		p.Post("/create", j.Serve, record.Create)
		p.Get("/column", j.Serve, record.GetColumn)
		p.Post("/data", j.Serve, record.Get)
		p.Patch("/{id:uint}/update", j.Serve, record.Update)
		p.Get("/{id:uint}/edit", j.Serve, record.Edit)
		p.Patch("/{id:uint}/changeStatus", j.Serve, record.ChangeStatus)
		p.Delete("/{id:uint}/delete", j.Serve, record.Delete)
		p.Get("/{id:uint}/operation", j.Serve, record.Operation)
		p.Get("/{id:uint}/getFormerData", j.Serve, record.GetFormerData)
		p.Post("/{id:uint}/UpdateFormerData", j.Serve, record.UpdateFormerData)
		p.Get("/{id:uint}/GetSoNoOptions", j.Serve, record.GetSoNoOptions)
		p.Post("/{id:uint}/UpdateCargoInfo",j.Serve,record.UpdateCargoInfo)
		p.Post("/DeleteCargoInfo",j.Serve,record.DeleteCargoInfo)
	})
}

func NewRouteOrderMaster(route *Route) *OrderMaster {
	return &OrderMaster{route}
}
