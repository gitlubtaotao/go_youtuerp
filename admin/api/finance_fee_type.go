package api

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/global"
	"youtuerp/models"
	"youtuerp/services"
)

type FinanceFeeType struct {
	BaseApi
	service     services.IFinanceBase
	codeService services.IBaseCode
	ctx         iris.Context
	sy          sync.Mutex
}

func (f *FinanceFeeType) GetColumn(ctx iris.Context) {
	f.RenderModuleColumn(ctx, models.FinanceFeeType{})
}
func (f *FinanceFeeType) Get(ctx iris.Context) {
	currencyOptions := f.codeService.FindCollect(models.CodeFinanceCurrency)
	feeTypes, total, err := f.service.FindFeeType(f.GetPer(ctx), f.GetPage(ctx), f.handleParams(), []string{}, []string{})

	dataArray := make([]map[string]interface{}, 0, len(feeTypes))
	if err != nil {
		f.Render500(ctx, err, "")
		return
	}
	f.sy.Lock()
	defer f.sy.Unlock()
	for _, i := range feeTypes {
		dataArray = append(dataArray, f.handleData(i))
	}
	_, _ = ctx.JSON(iris.Map{
		"code":            http.StatusOK,
		"data":            dataArray,
		"total":           total,
		"currencyOptions": currencyOptions,
	})
}
func (f *FinanceFeeType) Create(ctx iris.Context) {
	var (
		err     error
		feeType models.FinanceFeeType
	)
	if err = ctx.ReadJSON(&feeType); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	record, err := f.service.Create(feeType, ctx.GetLocale().Language())
	if err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	feeType = record.(models.FinanceFeeType)
	f.RenderSuccessJson(ctx, f.handleData(feeType))
}
func (f *FinanceFeeType) Update(ctx iris.Context) {
	var (
		id     uint
		record models.FinanceFeeType
		err    error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&record); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	if err = f.service.Update(id, record, ctx.GetLocale().Language()); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	record.ID = id
	f.RenderSuccessJson(ctx, f.handleData(record))
}

func (f *FinanceFeeType) Delete(ctx iris.Context) {
	var (
		id  uint
		err error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	err = f.service.Delete(id, models.FinanceFeeType{})
	if err != nil {
		f.Render500(ctx, err, "")
	} else {
		f.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (f *FinanceFeeType) Before(ctx iris.Context) {
	f.service = services.NewFinanceBase()
	f.codeService = services.NewBaseCode()
	f.ctx = ctx
	ctx.Next()
}

func (f *FinanceFeeType) handleData(feeType models.FinanceFeeType) map[string]interface{} {
	data, err := f.StructToMap(feeType, f.ctx)
	if err != nil {
		return map[string]interface{}{}
	}
	data["finance_currency_id_value"] = data["finance_currency_id"]
	data["finance_currency_id"] = global.RedSetting.HGetRecord("base_data_codes", data["finance_currency_id"], "")
	return data
}

func (f *FinanceFeeType) handleParams() map[string]interface{} {
	params := make(map[string]interface{}, 2)
	params["name-rCont"] = f.ctx.URLParamDefault("name", "")
	params["name_cn-rCount"] = f.ctx.URLParamDefault("name_cn", "")
	params["finance_currency_id-eq"] = f.ctx.URLParamIntDefault("finance_currency_id", 0)
	return params
}
