package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/services"
)

type FinanceFeeType struct {
	BaseController
	service services.IFinanceBase
	ctx     iris.Context
}

func (f *FinanceFeeType) GetColumn(ctx iris.Context) {
	f.RenderModuleColumn(ctx,models.FinanceFeeType{})
}
func (f *FinanceFeeType) Get(ctx iris.Context) {

}
func (f *FinanceFeeType) Create(ctx iris.Context) {
	var (
		err     error
		feeType models.FinanceFeeType
	)
	if err = ctx.ReadJSON(feeType); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	record, err := f.service.Create(feeType, ctx.GetLocale().Language())
	if err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	feeType = record.(models.FinanceFeeType)
	f.RenderSuccessJson(ctx, f.handleData(redis.NewRedis(), feeType))
}
func (f *FinanceFeeType) Update(ctx iris.Context) {

}
func (f *FinanceFeeType) Delete(ctx iris.Context) {

}

func (f *FinanceFeeType) Before(ctx iris.Context) {
	f.service = services.NewFinanceBase()
	f.ctx = ctx
	ctx.Next()
}

func (f *FinanceFeeType) handleData(red redis.Redis, feeType models.FinanceFeeType) map[string]interface{} {
	data, err := f.StructToMap(feeType, f.ctx)
	if err != nil {
		return map[string]interface{}{}
	}
	data["finance_currency_id_value"] = data["finance_currency_id"]
	data["finance_currency_id"] = red.HGetRecord("base_data_codes", data["finance_currency_id"], "")
	return data
}

func (f *FinanceFeeType) HandleParams() map[string]interface{} {
	return map[string]interface{}{}
}
