package api

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/global"
	"youtuerp/internal/models"
	"youtuerp/internal/services"
)

type Address struct {
	BaseApi
	service services.IAddressService
	ctx     iris.Context
	mu      sync.Mutex
}

func (a *Address) GetColumn(ctx iris.Context) {
	a.RenderModuleColumn(ctx, models.Address{})
}
func (a *Address) Get(ctx iris.Context) {
	var (
		accounts []models.Address
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
		dataArray = append(dataArray, a.handlerData(v))
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total})
}

func (a *Address) Create(ctx iris.Context) {
	var (
		address models.Address
		err     error
	)
	if err = ctx.ReadJSON(&address); err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	address, err = a.service.Create(address, ctx.GetLocale().Language())
	if err != nil {
		a.Render500(ctx, err, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	a.RenderSuccessJson(ctx, a.handlerData(address))
}

func (a *Address) Update(ctx iris.Context) {
	var (
		address models.Address
		err     error
		id      int
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&address); err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	golog.Infof("%v", address)
	if address, err = a.service.UpdateById(uint(id), address, ctx.GetLocale().Language()); err != nil {
		a.Render500(ctx, err, "")
	}
	data := a.handlerData(address)
	a.RenderSuccessJson(ctx, data)
}

func (a *Address) Delete(ctx iris.Context) {
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

func (a *Address) Before(ctx iris.Context) {
	a.service = services.NewAddressService()
	a.ctx = ctx
	ctx.Next()
}

func (a *Address) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["user_name-rCount"] = a.ctx.URLParamDefault("name", "")
	searchColumn["user_company_id-eq"] = a.ctx.URLParamDefault("user_company_id", "")
	return searchColumn
}

func (a *Address) handlerData(address models.Address) map[string]interface{} {
	data, _ := a.StructToMap(address, a.ctx)
	data["user_company_id_value"] = data["user_company_id"]
	data["user_company_id"] = global.RedSetting.HGetCrm(data["user_company_id"], "name_nick")
	return data
}
