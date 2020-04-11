package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/models"
	"youtuerp/services"
)

type CompanyController struct {
	Ctx     iris.Context
	Service services.ICompanyService
	Column  services.IColumnService
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
	return c.RenderErrorJson(err.Error())
}

func (c *CompanyController) GetColumn() iris.Map {
	column := services.NewColumnService()
	data, err := column.DefaultColumn(models.UserCompany{})
	if err != nil {
		return c.RenderErrorJson(err.Error())
	}
	return c.RenderSuccessJson(iris.Map{
		"column": data,
	})
}
