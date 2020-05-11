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
	currentUser, _ := c.CurrentUser(ctx)
	selectColumn := c.GetCustomerColumn(currentUser, models.UserCompany{})
	limit := ctx.URLParamIntDefault("limit", 20)
	page := ctx.URLParamIntDefault("page", 1)
	companies, total, err := c.Service.FindCompany(uint(limit), uint(page), c.handlerGetParams(), selectColumn, []string{}, true)
	if err != nil {
		c.Render500(ctx, err, err.Error())
		return
	}
	
	dataArray := make([]map[string]interface{}, 0)
	transportArray := c.Service.TransportTypeArrays(ctx.GetLocale())
	for _, v := range companies {
		dataArray = append(dataArray, c.itemChange(v, transportArray))
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total,})
}

//
func (c *CompanyController) Create(ctx iris.Context) {
	var company models.UserCompany
	err := ctx.ReadJSON(&company)
	if err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	company.CompanyType = 4
	company.Status = "approved"
	company.IsHeadOffice = false
	validator := services.NewValidatorService(&company)
	errM:= validator.ResultError(ctx.GetLocale().Language())
	if errM != "" {
		c.Render400(ctx,nil,errM)
		return
	}
	company, err = c.Service.Create(company)
	if err != nil {
		c.Render500(ctx, err, "")
		return
	}
	data, _ := c.StructToMap(company, ctx)
	c.RenderSuccessJson(ctx, data)
}

func (c *CompanyController) Edit(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		c.Render400(ctx,err,err.Error())
		return
	}
	company, err := c.Service.FirstCompany(id)
	if err != nil {
		c.Render500(ctx,err,err.Error())
		return
	}
	c.RenderSuccessJson(ctx, company)
}

//更新公司信息
func (c *CompanyController) Update(ctx iris.Context) {
	var readData models.UserCompany
	err := ctx.ReadJSON(&readData)
	if err != nil {
		c.Render400(ctx,err,err.Error())
		return
	}
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		c.Render400(ctx, err, ctx.GetLocale().GetMessage("error.error"))
		return
	}
	company, _ := c.Service.FirstCompany(id)
	err = c.Service.Update(company, readData)
	if err != nil {
		c.Render500(ctx,err ,"")
		return
	}
	c.RenderSuccessJson(ctx, c.itemChange(company, c.Service.TransportTypeArrays(ctx.GetLocale())))
}

func (c *CompanyController) GetColumn(ctx iris.Context) {
	c.RenderModuleColumn(ctx, models.UserCompany{})
}

func (c *CompanyController) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		c.Render400(ctx, err, ctx.GetLocale().GetMessage("error.error"))
		return
	}
	if err = c.Service.Delete(id); err != nil {
		c.Render500(ctx, err, ctx.GetLocale().GetMessage("error.error"))
	} else {
		c.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (c *CompanyController) Show(ctx iris.Context) {
	var (
		id      int
		err     error
		company models.UserCompany
	)
	id, err = ctx.Params().GetInt("id")
	if err != nil {
		c.Render400(ctx, err, "")
		return
	}
	company, err = c.Service.FirstCompanyByRelated(uint(id), "Employees", "Accounts", "Departments")
	if err != nil {
		c.Render500(ctx, err, "")
		return
	}
	data, _ := c.StructToMap(company, ctx)
	c.RenderSuccessJson(ctx, data)
}

func (c *CompanyController) Before(ctx iris.Context) {
	c.Service = services.NewCompanyService()
	c.Ctx = ctx
	ctx.Next()
}

func (c *CompanyController) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	if c.Ctx.URLParamExists("name_nick") && c.Ctx.URLParam("name_nick") != "" {
		searchColumn["name_nick-cont"] = c.Ctx.URLParam("name_nick")
	}
	if c.Ctx.URLParamExists("email") && c.Ctx.URLParam("email") != "" {
		searchColumn["email-cont"] = c.Ctx.URLParam("email")
	}
	return searchColumn
}

func (c *CompanyController) itemChange(v *models.UserCompany, transportArray []map[string]interface{}) map[string]interface{} {
	temp, _ := c.StructToMap(v, c.Ctx)
	temp["company_type"] = c.Service.ShowTransportType(c.Ctx.GetLocale(), temp["company_type"], transportArray)
	return temp
}
