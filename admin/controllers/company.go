package controllers

import (
	"fmt"
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
	c.initService(ctx)
	currentUser, _ := c.CurrentUser(ctx)
	selectColumn := c.GetModelColumn(currentUser, models.UserCompany{})
	limit := ctx.URLParamIntDefault("limit", 20)
	page := ctx.URLParamIntDefault("page", 1)
	companies, total, err := c.Service.FindCompany(uint(limit), uint(page), c.handlerGetParams(), selectColumn, []string{}, true)
	if err != nil {
		fmt.Println(err)
		c.RenderErrorJson(c.Ctx, http.StatusInternalServerError, err.Error())
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	transportArray := c.Service.TransportTypeArrays(ctx.GetLocale())
	for k, v := range companies {
		temp, _ := c.StructToMap(v, ctx)
		temp["index_col"] = k + 1
		temp["company_type"] = c.Service.ShowTransportType(ctx.GetLocale(), temp["company_type"], transportArray)
		dataArray = append(dataArray, temp)
	}
	
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total,})
	
}

//
func (c *CompanyController) Create(ctx iris.Context) {

}

func (c *CompanyController) GetColumn(ctx iris.Context) {
	c.RenderModuleColumn(ctx, models.UserCompany{})
}

func (c *CompanyController) initService(ctx iris.Context) {
	c.Service = services.NewCompanyService()
	c.Ctx = ctx
}

func (c *CompanyController) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	if c.Ctx.URLParamExists("name_nick") && c.Ctx.URLParam("name_nick") != "" {
		searchColumn["name_nick-cont"] = c.Ctx.URLParam("name_nick")
	}
	if c.Ctx.URLParamExists("email") && c.Ctx.URLParam("email") != "" {
		searchColumn["email-cont"] = c.Ctx.URLParam("email")
	}
	fmt.Println(searchColumn)
	return searchColumn
}
