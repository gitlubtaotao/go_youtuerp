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
func (c *CompanyController) Get(ctx iris.Context) {
	//companies, err := c.Service.AllCompany(map[string]interface{}{
	//	"name_nick": "优途互联",
	//}, []string{"user_companies.name_nick", "user_companies.id"}, []string{})
	//if err == nil {
	//	return c.RenderSuccessJson(c.Ctx, companies)
	//}
	//return c.RenderErrorMap(c.Ctx, http.StatusInternalServerError, err.Error())
}

//
func (c *CompanyController) Create(ctx iris.Context) {

}

func (c *CompanyController) GetColumn() iris.Map {
	return c.RenderColumnMap(c.Ctx, models.UserCompany{})
}
