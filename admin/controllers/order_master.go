package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/services"
)

type OrderMaster struct {
	BaseController
	ctx     iris.Context
	service services.IOrderMasterService
}

func (o *OrderMaster) Get(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{})
}
func (o *OrderMaster) Create(ctx iris.Context) {

}
func (o *OrderMaster) Update(ctx iris.Context) {

}

func (o *OrderMaster) Before(ctx iris.Context) {
	o.ctx = ctx
	o.service = services.NewOrderMasterService()
	ctx.Next()
}
