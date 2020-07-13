package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/models"
	"youtuerp/services"
)

type AccountController struct {
	service services.IAccountService
	BaseController
	ctx iris.Context
	mu  sync.Mutex
}

func (a *AccountController) GetColumn(ctx iris.Context) {
	a.RenderModuleColumn(ctx, models.Account{})
}

func (a *AccountController) Get(ctx iris.Context) {
	var (
		accounts []models.Account
		total    int64
		err      error
	)
	ty := ctx.URLParamDefault("type", "oa")
	if ty == "oa" {
		accounts, total, err = a.service.FindByOa(a.GetPer(ctx), a.GetPage(ctx), a.handlerGetParams(), []string{}, []string{})
	} else {
		accounts, total, err = a.service.FindByCrm(a.GetPer(ctx), a.GetPage(ctx), a.handlerGetParams(), []string{}, []string{})
	}
	if err != nil {
		a.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, v := range accounts {
		dataArray = append(dataArray, a.handlerData(v, ty))
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total,})
}

func (a *AccountController) Create(ctx iris.Context) {
	var (
		account models.Account
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
	data, _ := a.StructToMap(account, ctx)
	a.RenderSuccessJson(ctx, data)
}

func (a *AccountController) Update(ctx iris.Context) {
	var (
		updateContent models.Account
		account       models.Account
		err           error
		id            int
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&updateContent); err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	if account, err = a.service.UpdateById(uint(id), updateContent, ctx.GetLocale().Language()); err != nil {
		a.Render500(ctx, err, "")
	}
	returnData, _ := a.StructToMap(account, ctx)
	a.RenderSuccessJson(ctx, returnData)
}

func (a *AccountController) Edit(ctx iris.Context) {

}
func (a *AccountController) Delete(ctx iris.Context) {
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

func (a *AccountController) Before(ctx iris.Context) {
	a.service = services.NewAccountService()
	a.ctx = ctx
	ctx.Next()
}

func (a *AccountController) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["name-rCount"] = a.ctx.URLParamDefault("name", "")
	searchColumn["user_name-rCount"] = a.ctx.URLParamDefault("user_name", "")
	searchColumn["bank_name-rCount"] = a.ctx.URLParamDefault("bank_name", "")
	searchColumn["bank_number-rCount"] = a.ctx.URLParamDefault("bank_number", "")
	searchColumn["category-eq"] = a.ctx.URLParamDefault("category", "")
	searchColumn["user_company_id-eq"] = a.ctx.URLParamDefault("user_company_id", "")
	return searchColumn
}
func (a *AccountController) handlerData(account models.Account, ty string) map[string]interface{} {
	data, _ := a.StructToMap(account, a.ctx)
	data["user_company_id_value"] = data["user_company_id"]
	if ty == "oa" {
		data["user_company_id"] = red.HGetCompany(data["user_company_id"], "name_nick")
	} else {
		data["user_company_id"] = red.HGetCrm(data["user_company_id"], "name_nick")
	}
	return data
}
