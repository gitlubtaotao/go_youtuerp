package routers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/admin/controllers"
)

type Finance struct {
	route *Routers
}

func (f *Finance) Index() {
	f.fee()
	f.feeType()
	f.rate()
}

func (f *Finance) fee() {
	r := f.route
	j := r.jwtAccess()
	r.app.PartyFunc("/finance/fees", func(c iris.Party) {
		record := controllers.FinanceFee{}
		c.Use(record.Before)
		c.Post("/create", j.Serve, record.Create)
		c.Get("/{id:uint}/OrderFees", j.Serve, record.OrderFees)
		c.Post("/DeleteFee", j.Serve, record.DeleteFee)
		c.Post("/ChangeStatus", j.Serve, record.ChangeStatus)
		c.Post("/CopyFee", j.Serve, record.CopyFee)
		c.Get("/GetHistoryFee", j.Serve, record.GetHistoryFee)
		c.Post("/{orderMasterId:uint}/BulkHistoryFee", j.Serve, record.BulkHistoryFee)
		c.Get("/column", j.Serve, record.GetColumn)
		c.Post("/GetConfirmBillList", j.Serve, record.GetConfirmBillList)
	})
}

func (f *Finance) feeType() {
	r := f.route
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
	r := f.route
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
func newRouteFinance(r *Routers) *Finance {
	return &Finance{route: r}
}
