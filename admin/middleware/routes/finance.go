package routes

import (
	"github.com/kataras/iris/v12"
	"youtuerp/admin/controllers"
)

type Finance struct {
	Route *Route
}

func (f *Finance) Index() {
	f.fee()
	f.feeType()
	f.rate()
}

func (f *Finance) fee() {
	r := f.Route
	j := r.jwtAccess()
	r.app.PartyFunc("/finance/fees", func(c iris.Party) {
		record := controllers.FinanceFee{}
		c.Use(record.Before)
		c.Post("/create", j.Serve, record.Create)
		c.Get("/{id:uint}/OrderFees", j.Serve, record.OrderFees)
		c.Post("/DeleteFee", j.Serve, record.DeleteFee)
		c.Post("/ChangeStatus", j.Serve, record.ChangeStatus)
	})
}

func (f *Finance) feeType() {
	r := f.Route
	j := r.jwtAccess()
	r.app.PartyFunc("/finance/fee_types", func(c iris.Party) {
		record := controllers.FinanceFeeType{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Post("/data", j.Serve, record.Get)
		c.Patch("/{id:uint}/update", j.Serve, record.Update)
		c.Delete("/{id:uint}/destroy", j.Serve, record.Delete)
	})
}

func (f *Finance) rate() {
	r := f.Route
	j := r.jwtAccess()
	r.app.PartyFunc("/finance/rates", func(c iris.Party) {
		record := controllers.FinanceRate{}
		c.Use(record.Before)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/create", j.Serve, record.Create)
		c.Post("/data", j.Serve, record.Get)
		c.Delete("/{id:uint}/delete", j.Serve, record.Delete)
	})
}
func NewRouteFinance(r *Route) *Finance {
	return &Finance{r}
}
