//管理供应商信息
package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/services"
)

type SupplierController struct {
	BaseController
	Ctx     iris.Context
	Service services.ICrmCompanyService
}

func (s *SupplierController) GetColumn() iris.Map {
	return iris.Map{}
}
