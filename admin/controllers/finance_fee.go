package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/models"
	"youtuerp/services"
)

type FinanceFee struct {
	BaseController
	service services.IFinanceFee
}

func (f *FinanceFee) Create(ctx iris.Context) {

}

//获取订单对应的费用
func (f *FinanceFee) OrderFees(ctx iris.Context) {
	var (
		id      uint
		err     error
		options map[string]interface{}
		data    map[string][]models.FinanceFee
		sy      sync.WaitGroup
		sm      sync.Mutex
	)
	sy.Add(2)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	go func(id uint) {
		sm.Lock()
		defer sm.Unlock()
		data, err = f.service.OrderFees(id, "pay", "receive")
		sy.Done()
	}(id)
	go func() {
		sm.Lock()
		defer sm.Unlock()
		options = f.service.OrderFeesOptions()
		sy.Done()
	}()
	sy.Wait()
	if err != nil {
		f.Render500(ctx, err, "")
		return
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": data, "options": options})
}

func (f *FinanceFee) Before(ctx iris.Context) {
	f.service = services.NewFinanceFee()
	ctx.Next()
}
