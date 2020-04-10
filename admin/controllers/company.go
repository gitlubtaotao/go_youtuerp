package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/services"
)

//import (
//	"github.com/kataras/iris/v12"
//	_ "youtuerp/service"
//)
//
type CompanyController struct {
	Ctx iris.Context
	services.ICompanyService
}
//
//func (c *CompanyController) Get() iris.Map {
//
//	format := c.Ctx.URLParam("format")
//	return iris.Map{
//		"message": "Hello Iris!",
//		"format":  format,
//	}
//}
