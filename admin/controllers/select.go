package controllers

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"net/http"
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

func (s *SelectController) GetCommon(ctx iris.Context) {
	readData, err := s.base(ctx)
	if err != nil {
		s.renderError(ctx, err)
		return
	}
	if len(readData.SelectKeys) == 0 {
		readData.SelectKeys = []string{"name", "id"}
	}
	name := ctx.URLParamDefault("name", "")
	data, _ := s.service.FindTable(readData.TableName, name, readData.Scope, readData.SelectKeys)
	_, _ = s.ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}

//获取公司select_api
func (s *SelectController) GetCompany(ctx iris.Context) {
	readData, err := s.base(ctx)
	if err != nil {
		s.renderError(ctx, err)
		return
	}
	
	name := ctx.URLParamDefault("name", "")
	data, _ := s.service.GetCompanySelect(name,readData.Scope,readData.SelectKeys)
	_, _ = s.ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}


func (s *SelectController) Employee(ctx iris.Context) {
	service := services.NewEmployeeService()
	data := service.FindRedis()
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}

func (s *SelectController) OwnerCompany(ctx iris.Context) {
	service := services.NewCompanyService()
	data := service.AllCompanyRedis()
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}

func (s *SelectController) base(ctx iris.Context) (readData ReadData, err error) {
	s.service = services.NewSelectService(ctx)
	s.ctx = ctx
	err = ctx.ReadJSON(&readData)
	if err != nil {
		golog.Errorf("select api read data error %v", err)
		return
	}
	return
}

func (s *SelectController) renderError(ctx iris.Context, err error) {
	s.Render400(ctx, err, ctx.GetLocale().GetMessage("error.params_error"))
	return
}
