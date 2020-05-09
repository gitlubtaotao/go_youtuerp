package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type SettingController struct {
	BaseController
}

func (s *SettingController) Get(ctx iris.Context) {
	key := ctx.URLParamDefault("key", "system")
	service := services.NewSettingService()
	settings, err := service.Find(key)
	if err != nil {
		s.RenderErrorJson(ctx, http.StatusInternalServerError, "")
		return
	}
	s.RenderSuccessJson(ctx, settings)
}
func (s *SettingController) UpdateSystem(ctx iris.Context) {
	var (
		systemSetting []models.ResultSetting
		err           error
	)
	if err = ctx.ReadJSON(&systemSetting); err != nil {
		conf.IrisApp.Logger().Error(err)
		s.RenderErrorJson(ctx, 0, "")
		return
	}
	key := ctx.URLParamDefault("key", "base")
	service := services.NewSettingService()
	if err = service.UpdateSystem(key, systemSetting); err != nil {
		conf.IrisApp.Logger().Error(err)
		s.RenderErrorJson(ctx, 0, "")
		return
	}
	s.RenderSuccessJson(ctx, iris.Map{})
}

func (s *SettingController) UpdateUser(ctx iris.Context) {

}
