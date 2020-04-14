package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/models"
	"youtuerp/services"
)

type CompanyController struct {
	Ctx     iris.Context
	Service services.ICompanyService
	BaseController
}

//
func (c *CompanyController) Get() iris.Map {
	companies, err := c.Service.AllCompany(map[string]interface{}{
		"name_nick": "优途互联",
	}, []string{"user_companies.name_nick", "user_companies.id"}, []string{})
	if err == nil {
		return c.RenderSuccessJson(companies)
	}
	return c.RenderErrorJson(err.Error(),0)
}



func (c *CompanyController) GetColumn() iris.Map {
	return c.RenderColumnJson(models.UserCompany{}, c.Ctx.GetLocale())
}
