package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type NumberSettingController struct {
	BaseController
	ctx     iris.Context
	service services.INumberSettingService
}

func (n *NumberSettingController) GetColumn(ctx iris.Context) {
	n.RenderModuleColumn(ctx, models.ResultNumberSetting{})
}

func (n *NumberSettingController) Get(ctx iris.Context) {
	numberSettings, total, err := n.service.Find(n.GetPer(ctx), n.GetPage(ctx), n.handlerGetParams(), []string{}, []string{})
	if err != nil {
		conf.IrisApp.Logger().Errorf("number setting is err (%v)", err)
		n.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	enum := conf.Enum{Locale: ctx.GetLocale()}
	for _, v := range numberSettings {
		result, _ := n.StructToMap(v, ctx)
		result["clear_rule"] = enum.ClearRule(result["clear_rule"])
		result["application_no"] = enum.DefaultText("number_setting_application_no."+result["application_no"].(string))
		dataArray = append(dataArray, result)
	}
	n.RenderSuccessJson(ctx, iris.Map{
		"data":  dataArray,
		"total": total,
	})
}

func (n *NumberSettingController) Create(ctx iris.Context) {
	var (
		numberSetting models.NumberSetting
		err           error
	)
	if err = ctx.ReadJSON(&numberSetting); err != nil {
		conf.IrisApp.Logger().Warnf("read number setting json is err %v", err)
		n.RenderErrorJson(ctx, 0, "")
		return
	}
	fmt.Printf("%v", numberSetting)
	numberSetting, err = n.service.Create(numberSetting, ctx.GetLocale().Language())
	if err != nil {
		conf.IrisApp.Logger().Errorf("create number setting is error %v", err)
		n.RenderErrorJson(ctx, 0, err.Error())
		return
	}
	data, err := n.StructToMap(numberSetting, ctx)
	if err != nil {
		n.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	n.RenderSuccessJson(ctx, data)
}
func (n *NumberSettingController) Delete(ctx iris.Context) {
	var (
		id  int
		err error
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		n.RenderErrorJson(ctx, 0, "")
		return
	}
	if err = n.service.Delete(uint(id)); err != nil {
		n.RenderErrorJson(ctx, http.StatusInternalServerError, "")
		return
	}
	n.RenderSuccessJson(ctx, iris.Map{})
}

func (n *NumberSettingController) Before(ctx iris.Context) {
	n.ctx = ctx
	n.service = services.NewNumberSetting()
	ctx.Next()
}

func (n *NumberSettingController) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["number_settings.user_company_id-eq"] = n.ctx.URLParamDefault("user_company_id", "")
	return searchColumn
}
