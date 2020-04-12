package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/models"
	"youtuerp/services"
)

type EmployeeController struct {
	Ctx     iris.Context
	Service services.IEmployeeService
	BaseController
}

//
func (e *EmployeeController) Get() iris.Map {
	return e.RenderSuccessJson(make(map[string]interface{}))
}

func (e *EmployeeController) GetColumn() iris.Map {
	return e.RenderColumnJson(models.User{}, e.Ctx.GetLocale())
}
