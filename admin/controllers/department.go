//管理部门信息
package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type DepartmentController struct {
	Ctx iris.Context
	BaseController
	Service services.IDepartmentService
}

func (d *DepartmentController) GetColumn(ctx iris.Context) {
	fmt.Print("ssssss")
	d.RenderModuleColumn(ctx, models.Department{})
}

func (d *DepartmentController) Get(ctx iris.Context) () {
	
	departments, total, err := d.Service.Find(20, 1, map[string]interface{}{}, []string{}, []string{}, true)
	if err != nil {
		conf.IrisApp.Logger().Errorf("select department is err (%v)", err)
		d.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range departments {
		result, _ := d.StructToMap(v, ctx)
		dataArray = append(dataArray, result)
	}
	d.RenderSuccessJson(ctx, iris.Map{"data": dataArray, "total": total,})
}

func (d *DepartmentController) Before(ctx iris.Context) {
	d.Service = services.NewDepartmentService()
	d.Ctx = ctx
	ctx.Next()
}


