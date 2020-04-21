package controllers

//管理客户信息
import (
	"github.com/kataras/iris/v12"
	"youtuerp/services"
)

type CustomerController struct {
	Ctx iris.Context
	BaseController
	Service services.ICrmCompanyService
}

func (c *CustomerController) GetColumn() iris.Map {
	return iris.Map{}
}
