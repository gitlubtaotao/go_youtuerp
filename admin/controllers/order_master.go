package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/models"
	"youtuerp/services"
)

type OrderMaster struct {
	BaseController
	ctx             iris.Context
	service         services.IOrderMasterService
	companyService  services.ICompanyService
	employeeService services.IEmployeeService
}

func (o *OrderMaster) Get(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"code": http.StatusOK,
	})
}

func (o *OrderMaster) Create(ctx iris.Context) {
	var (
		order models.OrderMaster
		err   error
	)
	if err = ctx.ReadJSON(&order); err != nil {
		o.Render400(ctx, err, err.Error())
		return
	}
	currentUser, _ := o.CurrentUser(ctx)
	if order.OperationId == 0 {
		order.OperationId = currentUser.ID
	}
	order.CompanyId = uint(currentUser.UserCompanyId)
	order, err = o.service.Create(order, ctx.GetLocale().Language())
	if err != nil {
		o.Render400(ctx, err, err.Error())
		return
	}
	o.RenderSuccessJson(ctx, iris.Map{})
}

func (o *OrderMaster) Update(ctx iris.Context) {

}

func (o *OrderMaster) Before(ctx iris.Context) {
	o.ctx = ctx
	o.service = services.NewOrderMasterService()
	o.companyService = services.NewCompanyService()
	o.employeeService = services.NewEmployeeService()
	ctx.Next()
}
