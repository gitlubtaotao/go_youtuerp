package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type BaseCarrier struct {
	BaseController
	service services.IBaseCarrier
	ctx     iris.Context
	mu      sync.Mutex
	enum conf.Enum
}

func (b *BaseCarrier) GetColumn(ctx iris.Context) {
	b.RenderModuleColumn(ctx, models.BaseDataCarrier{})
}
func (b *BaseCarrier) Get(ctx iris.Context) {
	filter := b.handleParams()
	codes, total, err := b.service.Find(b.GetPer(ctx), b.GetPage(ctx),
		filter, []string{}, []string{})
	if err != nil {
		b.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, carrier := range codes {
		data := b.handleData(carrier)
		dataArray = append(dataArray, data)
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total})
}

func (b *BaseCarrier) Create(ctx iris.Context) {
	var (
		err  error
		code models.BaseDataCarrier
	)
	if err = ctx.ReadJSON(&code); err != nil {
		b.Render400(ctx, err, err.Error())
		return
	}
	if code, err = b.service.Create(code, ctx.GetLocale().Language()); err != nil {
		b.Render400(ctx, err, err.Error())
		return
	}
	b.RenderSuccessJson(ctx, b.handleData(code))
}
func (b *BaseCarrier) Update(ctx iris.Context) {
	var (
		id   uint
		err  error
		code models.BaseDataCarrier
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		b.Render400(ctx, err, err.Error())
		return
	}
	if err = b.service.Update(id, code, ctx.GetLocale().Language()); err != nil {
		b.Render400(ctx, err, err.Error())
		return
	}
	code.ID = id
	b.RenderSuccessJson(ctx, b.handleData(code))
}

func (b *BaseCarrier) Delete(ctx iris.Context) {
	var (
		id  uint
		err error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		b.Render400(ctx, err, err.Error())
		return
	}
	if err = b.service.Delete(id); err != nil {
		b.Render500(ctx, err, "")
	} else {
		b.RenderSuccessJson(ctx, iris.Map{})
	}
}
func (b *BaseCarrier) Before(ctx iris.Context) {
	b.service = services.NewBaseCarrier()
	b.ctx = ctx
	b.enum = conf.Enum{Locale: ctx.GetLocale()}
	ctx.Next()
}

func (b *BaseCarrier) handleParams() map[string]interface{} {
	data := make(map[string]interface{}, 0)
	data["name-rCount"] = b.ctx.URLParamDefault("name", "")
	data["type-eq"] = b.ctx.URLParamIntDefault("type", 0)
	return data
}

func (b *BaseCarrier) handleData(carrier models.BaseDataCarrier) map[string]interface{} {
	data, err := b.StructToMap(carrier, b.ctx)
	if err != nil {
		return map[string]interface{}{}
	}
	data["type_value"] = data["type"]
	data["type"] = b.enum.DefaultText("base_data_carriers_type.", data["type"])
	return data
}
