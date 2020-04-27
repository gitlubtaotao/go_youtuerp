package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/models"
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
	ctx     iris.Context
}

//获取公司select_api
func (s *SelectController) GetCompany(ctx iris.Context) {
	scope := map[string]interface{}{"company_type": 4}
	readData, err := s.base(ctx)
	if err != nil {
		s.renderError(ctx, err)
		return
	}
	if len(readData.SelectKeys) == 0 {
		readData.SelectKeys = []string{"name_en", "name_nick", "name_cn"}
	}
	if value, ok := readData.Scope["value"]; ok {
		if ctx.GetLocale().Language() == "en" {
			scope["name_en-cont"] = value
		} else {
			scope["name_cn-cont"] = value
		}
	}
	s.renderModel(&models.UserCompany{},scope,readData.SelectKeys)
}



//部门下拉
func (s *SelectController) GetDepartment(ctx iris.Context) {

}

func (s *SelectController) base(ctx iris.Context) (readData ReadData, err error) {
	s.service = services.NewSelectService(ctx)
	s.ctx = ctx
	err = ctx.ReadJSON(&readData)
	if err != nil {
		conf.IrisApp.Logger().Errorf("select api read data error %v", err)
		return
	}
	return
}

func (s *SelectController) renderError(ctx iris.Context, err error) {
	s.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("error.params_error"))
	return
}

func (s *SelectController) renderModel(model interface{},scope map[string]interface{},selectKeys []string) {
	data, err := s.service.FindModel(model, scope, selectKeys)
	if err != nil {
		conf.IrisApp.Logger().Errorf("select api read data error %v", err)
		s.RenderErrorJson(s.ctx, http.StatusBadRequest, s.ctx.GetLocale().GetMessage("error.params_error"))
		return
	}
	_, _ = s.ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}


