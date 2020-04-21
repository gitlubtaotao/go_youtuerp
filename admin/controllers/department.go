//管理部门信息
package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/services"
)

type DepartmentController struct {
	Ctx iris.Context
	BaseController
	Service services.DepartmentService
}

func (d *DepartmentController) GetColumn() iris.Map {
	return iris.Map{}
}
