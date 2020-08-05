package api

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/internal/models"
	"youtuerp/internal/services"
	"youtuerp/pkg/enumerize"
)

type BasePort struct {
	BaseApi
	service services.IBasePort
	ctx     iris.Context
	mu      sync.Mutex
	enum    enumerize.Enumerize
}

func (b *BasePort) GetColumn(ctx iris.Context) {
	b.RenderModuleColumn(ctx, models.BaseDataPort{})
}

func (b *BasePort) Get(ctx iris.Context) {
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
	for _, port := range codes {
		dataArray = append(dataArray, b.handleData(port))
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total})
}

func (b *BasePort) Create(ctx iris.Context) {
	var (
		err  error
		code models.BaseDataPort
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
func (b *BasePort) Update(ctx iris.Context) {
	var (
		id   uint
		err  error
		code models.BaseDataPort
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		b.Render400(ctx, err, err.Error())
		return
	}
	_ = ctx.ReadJSON(&code)
	if err = b.service.Update(id, code, ctx.GetLocale().Language()); err != nil {
		b.Render400(ctx, err, err.Error())
		return
	}
	code.ID = id
	b.RenderSuccessJson(ctx, b.handleData(code))
}

func (b *BasePort) Delete(ctx iris.Context) {
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
func (b *BasePort) Before(ctx iris.Context) {
	b.service = services.NewBasePort()
	b.ctx = ctx
	b.enum = enumerize.Enumerize{Locale: ctx.GetLocale()}
	ctx.Next()
}

func (b *BasePort) handleParams() map[string]interface{} {
	data := make(map[string]interface{}, 0)
	data["name-rCount"] = b.ctx.URLParamDefault("name", "")
	data["type-eq"] = b.ctx.URLParamIntDefault("type", 0)
	return data
}

func (b *BasePort) handleData(port models.BaseDataPort) map[string]interface{} {
	data, err := b.StructToMap(port, b.ctx)
	if err != nil {
		return map[string]interface{}{}
	}
	data["type_value"] = data["type"]
	data["type"] = b.enum.DefaultText("base_data_ports_type.", data["type"])
	return data
}
