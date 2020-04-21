package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/models"
	"youtuerp/services"
)

type CompanyController struct {
	Ctx     iris.Context
	Service services.ICompanyService
	BaseController
}

//
func (c *CompanyController) Get(ctx iris.Context) {
	c.initService()
	currentUser, _ := c.CurrentUser(ctx)
	selectColumn := c.GetModelColumn(currentUser, models.UserCompany{})
	companies, total, err := c.Service.FindCompany(20, 1, map[string]interface{}{}, selectColumn, []string{}, true)
	dataArray := make([]map[string]interface{}, len(companies)-1)
	transportArray := c.Service.TransportTypeArrays(ctx.GetLocale())
	for k, v := range companies {
		temp, _ := c.StructToMap(v, ctx)
		temp["index_col"] = k + 1
		temp["company_type"] = c.Service.ShowTransportType(ctx.GetLocale(),temp["company_type"],transportArray)
		dataArray = append(dataArray, temp)
	}
	if err == nil {
		_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total,})
	} else {
		c.RenderErrorJson(c.Ctx, http.StatusInternalServerError, err.Error())
	}
}

//
func (c *CompanyController) Create(ctx iris.Context) {

}

func (c *CompanyController) GetColumn(ctx iris.Context) {
	c.RenderModuleColumn(ctx, models.UserCompany{})
}


func (c *CompanyController) initService() {
	c.Service = services.NewCompanyService()
}

