package api

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/internal/models"
	"youtuerp/internal/services"
)

type Invoice struct {
	BaseApi
	service services.IInvoiceService
	ctx     iris.Context
	mu      sync.Mutex
}

func (a *Invoice) GetColumn(ctx iris.Context) {
	a.RenderModuleColumn(ctx, models.Invoice{})
}

func (a *Invoice) Get(ctx iris.Context) {
	var (
		invoices []models.Invoice
		total    int64
		err      error
	)
	ty := ctx.URLParamDefault("type", "oa")
	if ty == "oa" {
		invoices, total, err = a.service.FindByOa(a.GetPer(ctx), a.GetPage(ctx), a.handlerGetParams(), []string{}, []string{})
	} else {
		invoices, total, err = a.service.FindByCrm(a.GetPer(ctx), a.GetPage(ctx), a.handlerGetParams(), []string{}, []string{})
	}
	if err != nil {
		a.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, v := range invoices {
		dataArray = append(dataArray, a.handlerData(v))
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total})
}

func (a *Invoice) Create(ctx iris.Context) {
	var (
		account models.Invoice
		err     error
	)
	if err = ctx.ReadJSON(&account); err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	account, err = a.service.Create(account, ctx.GetLocale().Language())
	if err != nil {
		a.Render500(ctx, err, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	a.RenderSuccessJson(ctx, a.handlerData(account))
}

func (a *Invoice) Update(ctx iris.Context) {
	var (
		account models.Invoice
		err     error
		id      int
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&account); err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	if account, err = a.service.UpdateById(uint(id), account, ctx.GetLocale().Language()); err != nil {
		a.Render500(ctx, err, "")
	}
	a.RenderSuccessJson(ctx, a.handlerData(account))
}

func (a *Invoice) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	if err = a.service.Delete(uint(id)); err != nil {
		a.Render500(ctx, err, "")
	} else {
		a.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (a *Invoice) Before(ctx iris.Context) {
	a.service = services.NewInvoiceService()
	a.ctx = ctx
	ctx.Next()
}

func (a *Invoice) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["name-rCount"] = a.ctx.URLParamDefault("name", "")
	searchColumn["user_company_id-eq"] = a.ctx.URLParamDefault("user_company_id", "")
	return searchColumn
}

func (a *Invoice) handlerData(account models.Invoice) map[string]interface{} {
	data, _ := a.StructToMap(account, a.ctx)
	data["user_company_id_value"] = data["user_company_id"]
	data["user_company_id"] = RedSetting.HGetCrm(data["user_company_id"], "name_nick")
	return data
}
