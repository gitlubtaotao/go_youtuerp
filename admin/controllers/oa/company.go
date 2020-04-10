package oa

import (
	"github.com/kataras/iris/v12"
	"youtuerp/services/oa"
)

type CompanyController struct {
	Ctx     iris.Context
	service oa.CompanyService
}

func (c *CompanyController) Get() iris.Map {
	format := c.Ctx.URLParam("format")
	return iris.Map{
		"message": "Hello Iris!",
		"format":  format,
	}
}
