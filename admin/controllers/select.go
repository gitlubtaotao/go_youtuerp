package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/services"
)

type ReadData struct {
	TableName  string                 `json:"table_name"`
	SelectKeys []string               `json:"select_keys"`
	Scope      map[string]interface{} `json:"scope"`
}

type SelectController struct {
	BaseController
	service services.ISelectService
}

func (s *SelectController) Get(ctx iris.Context) {
	s.service = services.NewSelectService(ctx)
	var (
		readData ReadData
		err      error
		data     []map[string]interface{}
	)
	err = ctx.ReadJSON(&readData)
	if err != nil {
		conf.IrisApp.Logger().Errorf("select api read data error %v", err)
		s.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("error.params_error"))
		return
	}
	data, err = s.service.Find(readData.TableName, readData.Scope, readData.SelectKeys)
	if err != nil {
		conf.IrisApp.Logger().Errorf("select api read data error %v", err)
		s.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("error.params_error"))
		return
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}
