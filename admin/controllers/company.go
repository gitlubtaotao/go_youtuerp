package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type CompanyController struct {
	ctx     iris.Context
	service services.ICompanyService
	enum    conf.Enum
	BaseController
}

//
func (c *CompanyController) Get(ctx iris.Context) {
	companies, total, err := c.service.FindCompany(c.GetPer(ctx), c.GetPage(ctx), c.handlerGetParams(), []string{}, []string{}, true)
	if err != nil {
		c.Render500(ctx, err, err.Error())
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range companies {
		dataArray = append(dataArray, c.itemChange(v))
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
	validator := services.NewValidatorService(&company)
	errM := validator.ResultError(ctx.GetLocale().Language())
	if errM != "" {
		c.Render400(ctx, nil, errM)
		return
	}
	company, err = c.service.Create(company)
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
		c.Render400(ctx, err, err.Error())
		return
	}
	company, err := c.service.FirstCompany(id)
	if err != nil {
		c.Render500(ctx, err, err.Error())
		return
	}
	c.RenderSuccessJson(ctx, company)
}

//更新公司信息
func (c *CompanyController) Update(ctx iris.Context) {
	var readData models.UserCompany
	err := ctx.ReadJSON(&readData)
	if err != nil {
		c.Render400(ctx, err, "")
		return
	}
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		c.Render400(ctx, err, "")
		return
	}
	err = c.service.Update(id, readData)
	if err != nil {
		c.Render500(ctx, err, "")
		return
	}
	c.RenderSuccessJson(ctx, c.itemChange(readData))
}

func (c *CompanyController) GetColumn(ctx iris.Context) {
	c.RenderModuleColumn(ctx, models.UserCompany{})
}

func (c *CompanyController) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		c.Render400(ctx, err, "")
		return
	}
	if err = c.service.Delete(id); err != nil {
		c.Render500(ctx, err, "")
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
	company, err = c.service.FirstCompanyByRelated(uint(id), "Employees", "Accounts", "Departments")
	if err != nil {
		c.Render500(ctx, err, "")
		return
	}
	data, _ := c.StructToMap(company, ctx)
	c.RenderSuccessJson(ctx, data)
}

func (c *CompanyController) Before(ctx iris.Context) {
	c.service = services.NewCompanyService()
	c.enum = conf.NewEnum(ctx.GetLocale())
	c.ctx = ctx
	ctx.Next()
}

func (c *CompanyController) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	if c.ctx.URLParamExists("name_nick") && c.ctx.URLParam("name_nick") != "" {
		searchColumn["name_nick-rCount"] = c.ctx.URLParam("name_nick")
	}
	if c.ctx.URLParamExists("email") && c.ctx.URLParam("email") != "" {
		searchColumn["email-rCount"] = c.ctx.URLParam("email")
	}
	return searchColumn
}

func (c *CompanyController) itemChange(v models.UserCompany) map[string]interface{} {
	temp, _ := c.StructToMap(v, c.ctx)
	temp["company_type"] = c.service.ShowTransportType(c.enum, temp["company_type"])
	return temp
}
