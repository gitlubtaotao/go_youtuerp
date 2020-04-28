//管理员工信息

package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type EmployeeController struct {
	ctx     iris.Context
	service services.IEmployeeService
	BaseController
}

//
//func (e *EmployeeController) Get() iris.Map {
//	return e.RenderSuccessMap(e.Ctx, make(map[string]interface{}))
//}
func (e *EmployeeController) GetColumn(ctx iris.Context) {
	e.RenderModuleColumn(ctx, models.ResultEmployee{})
}

func (e *EmployeeController) Get(ctx iris.Context) {
	employees, total, err := e.service.Find(e.GetPer(ctx), e.GetPage(ctx), e.handlerGetParams(), []string{}, []string{}, true)
	if err != nil {
		conf.IrisApp.Logger().Errorf("employee is err (%v)", err)
		e.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range employees {
		fmt.Printf("%v", v)
		result, _ := e.StructToMap(v, ctx)
		dataArray = append(dataArray, result)
	}
	e.RenderSuccessJson(ctx, iris.Map{"data": dataArray, "total": total,})
}

func (e *EmployeeController) Create(ctx iris.Context) {

}

func (e *EmployeeController) Update(ctx iris.Context) {

}

func (e *EmployeeController) Delete(ctx iris.Context) {

}

func (e *EmployeeController) Before(ctx iris.Context) {
	e.service = services.NewEmployeeService()
	e.ctx = ctx
	ctx.Next()
}

func (e *EmployeeController) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["name-rCount"] = e.ctx.URLParamDefault("name", "")
	searchColumn["phone-rCount"] = e.ctx.URLParamDefault("phone", "")
	searchColumn["email-rCount"] = e.ctx.URLParamDefault("email", "")
	searchColumn["users.user_company_id-eq"] = e.ctx.URLParamDefault("user_company_id", "")
	searchColumn["users.department_id-eq"] = e.ctx.URLParamDefault("department_id", "")
	return searchColumn
}
