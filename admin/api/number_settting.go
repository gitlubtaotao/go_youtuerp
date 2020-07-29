package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"youtuerp/conf"
	"youtuerp/internal/models"
	"youtuerp/internal/services"
)

type NumberSetting struct {
	BaseApi
	ctx     iris.Context
	service services.INumberSettingService
}

func (n *NumberSetting) GetColumn(ctx iris.Context) {
	n.RenderModuleColumn(ctx, models.ResponseNumberSetting{})
}

func (n *NumberSetting) Get(ctx iris.Context) {
	numberSettings, total, err := n.service.Find(n.GetPer(ctx), n.GetPage(ctx), n.handlerGetParams(), []string{}, []string{})
	if err != nil {
		n.Render500(ctx, err, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	enum := conf.Enum{Locale: ctx.GetLocale()}
	fmt.Print(ctx.GetLocale().GetMessage("clear_rule"))
	for _, v := range numberSettings {
		result, _ := n.StructToMap(v, ctx)
		result["clear_rule"] = enum.ClearRuleText(result["clear_rule"])
		result["application_no"] = enum.DefaultText("number_setting_application_no.", result["application_no"])
		dataArray = append(dataArray, result)
	}
	n.RenderSuccessJson(ctx, iris.Map{
		"data":  dataArray,
		"total": total,
	})
}

func (n *NumberSetting) Create(ctx iris.Context) {
	var (
		numberSetting models.NumberSetting
		err           error
	)
	if err = ctx.ReadJSON(&numberSetting); err != nil {
		n.Render400(ctx, err, "")
		return
	}
	numberSetting, err = n.service.Create(numberSetting, ctx.GetLocale().Language())
	if err != nil {
		n.Render500(ctx, err, err.Error())
		return
	}
	data, err := n.StructToMap(numberSetting, ctx)
	if err != nil {
		n.Render500(ctx, err, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	n.RenderSuccessJson(ctx, data)
}
func (n *NumberSetting) Delete(ctx iris.Context) {
	var (
		id  int
		err error
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		n.Render400(ctx, err, "")
		return
	}
	if err = n.service.Delete(uint(id)); err != nil {
		n.Render500(ctx, err, "")
		return
	}
	n.RenderSuccessJson(ctx, iris.Map{})
}

func (n *NumberSetting) Before(ctx iris.Context) {
	n.ctx = ctx
	n.service = services.NewNumberSetting()
	ctx.Next()
}

func (n *NumberSetting) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["number_settings.user_company_id-eq"] = n.ctx.URLParamDefault("user_company_id", "")
	return searchColumn
}
