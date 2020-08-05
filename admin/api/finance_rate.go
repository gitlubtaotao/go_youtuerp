package api

import (
	"errors"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"sync"
	"time"
	"youtuerp/internal/models"
	"youtuerp/internal/services"
	"youtuerp/pkg/redisService"
	"youtuerp/pkg/util"
)

type FinanceRate struct {
	BaseApi
	service     services.IFinanceBase
	codeService services.IBaseCode
	ctx         iris.Context
	sy          sync.Mutex
}

func (f *FinanceRate) GetColumn(ctx iris.Context) {
	f.RenderModuleColumn(ctx, models.FinanceRate{})
}

func (f *FinanceRate) Get(ctx iris.Context) {
	rates, total, err := f.service.FindRate(f.GetPer(ctx), f.GetPage(ctx), f.handleParams(), []string{}, []string{})
	if err != nil {
		f.Render500(ctx, err, "")
		return
	}
	currencyOptions := f.codeService.FindCollect(models.CodeFinanceCurrency)
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range rates {
		dataArray = append(dataArray, f.handleData(v))
	}
	_, _ = ctx.JSON(iris.Map{
		"code":                   http.StatusOK,
		"data":                   dataArray,
		"total":                  total,
		"currencyOptions":        currencyOptions,
		"systemFinanceRate":      redisService.SystemRateSetting(),
		"systemStandardCurrency": redisService.SystemFinanceCurrency(),
	})
}

func (f *FinanceRate) Create(ctx iris.Context) {
	var (
		err  error
		rate models.FinanceRate
	)
	if err = ctx.ReadJSON(&rate); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	user, _ := f.CurrentUser(ctx)
	rate.UserId = user.ID
	rate.CompanyId = uint(user.UserCompanyId)
	rate.Year = uint(time.Now().Year())
	record, err := f.service.Create(rate, ctx.GetLocale().Language())
	if err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	if rate, ok := record.(models.FinanceRate); !ok {
		f.Render500(ctx, errors.New(ctx.GetLocale().GetMessage("error.inter_error")), "")
		return
	} else {
		f.RenderSuccessJson(ctx, f.handleData(rate))
	}
}
func (f *FinanceRate) Delete(ctx iris.Context) {
	var (
		id  uint
		err error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	if err = f.service.Delete(id, &models.FinanceRate{}); err != nil {
		f.Render500(ctx, err, "")
	} else {
		f.RenderSuccessJson(ctx, iris.Map{})
	}
}
func (f *FinanceRate) Before(ctx iris.Context) {
	f.service = services.NewFinanceBase()
	f.codeService = services.NewBaseCode()
	f.ctx = ctx
	ctx.Next()
}

func (f *FinanceRate) handleData(feeType models.FinanceRate) map[string]interface{} {
	data, err := f.StructToMap(feeType, f.ctx)
	if err != nil {
		return map[string]interface{}{}
	}
	data["finance_currency_id_value"] = data["finance_currency_id"]
	data["finance_currency_id"] = RedSetting.HGetValue("base_data_codesFinanceCurrency", data["finance_currency_id"], "")
	data["user_id_value"] = data["user_id"]
	data["user_id"] = RedSetting.HGetRecord(models.User{}.TableName(), data["user_id"], "name")
	data["company_id_value"] = data["company_id"]
	data["company_id"] = RedSetting.HGetCompany(data["company_id"], "")
	return data
}

func (f *FinanceRate) handleParams() map[string]interface{} {
	params := make(map[string]interface{}, 0)
	params["finance_currency_id-eq"] = f.ctx.URLParamIntDefault("finance_currency_id", 0)
	params["user_id-eq"] = f.ctx.URLParamIntDefault("user_id", 0)
	params["company_id-eq"] = f.ctx.URLParamIntDefault("company_id", 0)
	createdAt := strings.Split(f.ctx.URLParamDefault("created_at", ""), ",")
	//调用转化方法，传入上面准备好的的三个参数
	if len(createdAt) == 2 {
		tmp, _ := util.TimeHelper{}.StringToTime(createdAt[0])
		tmp2, _ := util.TimeHelper{}.StringToTime(createdAt[1])
		params["created_at-gtEq"] = tmp
		params["created_at-ltEq"] = tmp2
	}
	return params
}