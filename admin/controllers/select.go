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

//获取合作单位对应的数据
func (s *SelectController) GetCompany(ctx iris.Context) {
	readData, err := s.base(ctx)
	if err != nil {
		s.renderError(ctx, err)
		return
	}
	
	name := ctx.URLParamDefault("name", "")
	data, _ := s.service.GetCompanySelect(name, readData.Scope, readData.SelectKeys)
	_, _ = s.ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}

//获取公司员工对应的数据
func (s *SelectController) Employee(ctx iris.Context) {
	service := services.NewEmployeeService()
	data := service.FindRedis()
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}

//获取所属公司的数据
func (s *SelectController) OwnerCompany(ctx iris.Context) {
	service := services.NewCompanyService()
	data := service.AllCompanyRedis()
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}

//获取仓库地址对应的数据
func (s *SelectController) WarehouseAddress(ctx iris.Context) {
	service := services.NewBaseWarehouse()
	data, err := service.FindAllBySelect()
	if err != nil {
		s.Render500(ctx, err, "")
	} else {
		_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
	}
}

//获取订单的数据信息
func (s *SelectController) GetOrderMaster(ctx iris.Context) {
	serialNumber := ctx.URLParamDefault("serial_number", "")
	per := 100
	if serialNumber != "" {
		per = 0
	}
	service := services.NewOrderMasterService()
	filters := map[string]interface{}{
		"serial_number-rCount": serialNumber,
	}
	data, _, err := service.FindMasterNoTotal(per, 0, filters, []string{"order_masters.id", "serial_number"}, []string{"id desc"})
	if err != nil {
		s.Render500(ctx, err, "")
	} else {
		s.RenderSuccessJson(ctx, data)
	}
}

// 获取基础代码的数据
func (s *SelectController) GetBaseCode(ctx iris.Context) {
	key := ctx.URLParamDefault("key", "")
	service := services.NewBaseCode()
	s.RenderSuccessJson(ctx, service.FindCollect(key))
}

//获取承运方的数据
func (s *SelectController) GetCarrier(ctx iris.Context) {

}

//获取港口的数据
func (s *SelectController) GetPort(ctx iris.Context) {

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
