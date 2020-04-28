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
	d.RenderModuleColumn(ctx, models.ResultDepartment{})
}

func (d *DepartmentController) Get(ctx iris.Context) () {
	departments, total, err := d.Service.Find(d.GetPer(ctx), d.GetPage(ctx), d.handlerGetParams(), []string{}, []string{}, true)
	if err != nil {
		conf.IrisApp.Logger().Errorf("select department is err (%v)", err)
		d.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range departments {
		fmt.Printf("%v", v)
		result, _ := d.StructToMap(v, ctx)
		dataArray = append(dataArray, result)
	}
	d.RenderSuccessJson(ctx, iris.Map{"data": dataArray, "total": total,})
}

func (d *DepartmentController) Create(ctx iris.Context) {
	var (
		department models.Department
		err        error
	)
	err = ctx.ReadJSON(&department)
	if err != nil {
		d.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("error.params_error"))
		return
	}
	valid := services.NewValidatorService(department)
	errString := valid.ResultError(ctx.GetLocale().Language())
	if errString != "" {
		conf.IrisApp.Logger().Errorf("create department is err %s", errString)
		d.RenderErrorJson(ctx, http.StatusBadRequest, errString)
		return
	}
	department, err = d.Service.Create(department)
	if err != nil {
		d.RenderErrorJson(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	data, _ := d.StructToMap(department, ctx)
	d.RenderSuccessJson(ctx, data)
}

func (d *DepartmentController) Update(ctx iris.Context) {
	var (
		id       int
		err      error
		readData models.Department
	)
	id, err = ctx.Params().GetInt("id")
	if err != nil {
		d.RenderErrorJson(ctx, 0, "")
		return
	}
	err = ctx.ReadJSON(&readData)
	if err != nil {
		d.RenderErrorJson(ctx, 0, "")
		return
	}
	department, err := d.Service.First(uint(id))
	if err != nil {
		d.RenderErrorJson(ctx, 0, "")
		return
	}
	valid := services.NewValidatorService(readData)
	errString := valid.ResultError(ctx.GetLocale().Language())
	if errString != "" {
		d.RenderErrorJson(ctx, http.StatusBadRequest, errString)
		return
	}
	err = d.Service.Update(department, readData)
	if err != nil {
		conf.IrisApp.Logger().Error(err)
		d.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	returnData, _ := d.StructToMap(department, ctx)
	d.RenderSuccessJson(ctx, returnData)
	return
}

func (d *DepartmentController) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		d.RenderErrorJson(ctx, 0, "")
		return
	}
	err = d.Service.Delete(id)
	if err != nil{
		conf.IrisApp.Logger().Error(err)
		d.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
	}else{
		d.RenderSuccessJson(ctx,iris.Map{})
	}
}

func (d *DepartmentController) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["name_cn-rCount"] = d.Ctx.URLParamDefault("name_cn", "")
	searchColumn["name_en-rCount"] = d.Ctx.URLParamDefault("name_en", "")
	searchColumn["user_company_id-eq"] = d.Ctx.URLParamDefault("user_company_id", "")
	return searchColumn
}

func (d *DepartmentController) Before(ctx iris.Context) {
	d.Service = services.NewDepartmentService()
	d.Ctx = ctx
	ctx.Next()
}
