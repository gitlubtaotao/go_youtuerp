package controllers

import (
	"github.com/kataras/golog"
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
	var (
		financeFees []models.FinanceFee
		err         error
	)
	if err = ctx.ReadJSON(&financeFees); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	financeFees, err = f.service.BulkInsert(financeFees)
	if err != nil {
		f.Render500(ctx, err, err.Error())
		return
	}
	golog.Infof("finance fee is %v", financeFees)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": financeFees})
}

//获取订单对应的费用
func (f *FinanceFee) OrderFees(ctx iris.Context) {
	var (
		id                uint
		err               error
		options           map[string]interface{}
		closingUnitOption []map[string]interface{}
		data              map[string][]models.FinanceFee
		sy                sync.WaitGroup
		sm                sync.Mutex
	)
	currentUser, _ := f.CurrentUser(ctx)
	sy.Add(3)
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
	go func(companyId uint) {
		sm.Lock()
		defer sm.Unlock()
		options = f.service.OrderFeesOptions(companyId)
		sy.Done()
	}(uint(currentUser.UserCompanyId))
	go func() {
		sm.Lock()
		defer sm.Unlock()
		scope := map[string]interface{}{"company_type": []int{1, 2, 3, 4}}
		closingUnitOption, _ = services.NewSelectService(ctx).GetCompanySelect("", scope, []string{})
		sy.Done()
	}()
	sy.Wait()
	if err != nil {
		f.Render500(ctx, err, "")
		return
	}
	_, _ = ctx.JSON(iris.Map{
		"code":                 http.StatusOK,
		"data":                 data,
		"options":              options,
		"closing_unit_options": closingUnitOption})
}

func (f *FinanceFee) DeleteFee(ctx iris.Context) {
	var (
		deleteIds map[string][]uint
		err       error
	)
	if err = ctx.ReadJSON(&deleteIds); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	if ids, ok := deleteIds["ids"]; ok {
		if err = f.service.DeleteFees(ids); err != nil {
			f.Render500(ctx, err, "")
		} else {
			f.RenderSuccessJson(ctx, iris.Map{})
		}
	}
}

func (f *FinanceFee) Before(ctx iris.Context) {
	f.service = services.NewFinanceFee()
	ctx.Next()
}
