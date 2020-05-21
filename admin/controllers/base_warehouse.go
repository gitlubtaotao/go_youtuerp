package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type BaseWarehouse struct {
	BaseController
	service services.IBaseWarehouse
	ctx     iris.Context
	mu      sync.Mutex
	enum    conf.Enum
}

func (b *BaseWarehouse) GetColumn(ctx iris.Context) {
	b.RenderModuleColumn(ctx, models.BaseWarehouse{})
}
func (b *BaseWarehouse) Get(ctx iris.Context) {
	filter := b.handleParams()
	codes, total, err := b.service.Find(b.GetPer(ctx), b.GetPage(ctx),
		filter, []string{}, []string{})
	fmt.Printf("code is  %v,total is  %v,err is  %v", codes, total, err)
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

func (b *BaseWarehouse) Create(ctx iris.Context) {
	var (
		err  error
		code models.BaseWarehouse
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
func (b *BaseWarehouse) Update(ctx iris.Context) {
	var (
		id   uint
		err  error
		code models.BaseWarehouse
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		b.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&code); err != nil {
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

func (b *BaseWarehouse) Delete(ctx iris.Context) {
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
func (b *BaseWarehouse) Before(ctx iris.Context) {
	b.service = services.NewBaseWarehouse()
	b.ctx = ctx
	b.enum = conf.Enum{Locale: ctx.GetLocale()}
	ctx.Next()
}

func (b *BaseWarehouse) handleParams() map[string]interface{} {
	data := make(map[string]interface{}, 0)
	data["name-rCount"] = b.ctx.URLParamDefault("name", "")
	data["contact_name-eq"] = b.ctx.URLParamDefault("contact_name", "")
	data["contact_tel-eq"] = b.ctx.URLParamDefault("contact_tel", "")
	return data
}

func (b *BaseWarehouse) handleData(carrier models.BaseWarehouse) map[string]interface{} {
	data, err := b.StructToMap(carrier, b.ctx)
	if err != nil {
		return map[string]interface{}{}
	}
	return data
}
