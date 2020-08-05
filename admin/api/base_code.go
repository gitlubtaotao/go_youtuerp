package api

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/internal/models"
	"youtuerp/internal/services"
)

type BaseCodeController struct {
	service services.IBaseCode
	BaseApi
	ctx iris.Context
	mu  sync.Mutex
}

func (b *BaseCodeController) GetColumn(ctx iris.Context) {
	b.RenderModuleColumn(ctx, models.BaseDataCode{})
}

func (b *BaseCodeController) Get(ctx iris.Context) {
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
	for _, user := range codes {
		data := b.handleData(user)
		dataArray = append(dataArray, data)
	}
	codeLevel, _ := b.service.FindAllLevel()
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total, "code_level": codeLevel})
}

func (b *BaseCodeController) Create(ctx iris.Context) {
	var (
		err  error
		code models.BaseDataCode
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
func (b *BaseCodeController) Update(ctx iris.Context) {
	var (
		id   uint
		err  error
		code models.BaseDataCode
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
func (b *BaseCodeController) Delete(ctx iris.Context) {
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

func (b *BaseCodeController) Before(ctx iris.Context) {
	b.service = services.NewBaseCode()
	b.ctx = ctx
	ctx.Next()
}

func (b *BaseCodeController) handleParams() map[string]interface{} {
	data := make(map[string]interface{}, 0)
	data["name-rCount"] = b.ctx.URLParamDefault("name", "")
	data["code_name-eq"] = b.ctx.URLParamDefault("code_name", "")
	return data
}

func (b *BaseCodeController) handleData(code models.BaseDataCode) map[string]interface{} {
	data, _ := b.StructToMap(code, b.ctx)
	data["code_name_value"] = data["code_name"]
	data["code_name"] = RedSetting.HGetValue("base_data_levels", data["code_name"], "name")
	return data
}
