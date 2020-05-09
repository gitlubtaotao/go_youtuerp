package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type SettingController struct {
	BaseController
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
	fmt.Printf("setting data is %v",systemSetting)
	if err = service.UpdateSystem(key, systemSetting);err != nil{
		conf.IrisApp.Logger().Error(err)
		s.RenderErrorJson(ctx, 0, "")
		return
	}
	s.RenderSuccessJson(ctx, iris.Map{})
}

func (s *SettingController) UpdateUser(ctx iris.Context) {

}
