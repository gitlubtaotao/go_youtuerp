//管理员工信息

package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/services"
)

type EmployeeController struct {
	Ctx     iris.Context
	Service services.IEmployeeService
	BaseController
}

//
func (e *EmployeeController) Get() iris.Map {
	return e.RenderSuccessMap(e.Ctx, make(map[string]interface{}))
}
func (e *EmployeeController) GetColumn() iris.Map {
	return iris.Map{}
}


