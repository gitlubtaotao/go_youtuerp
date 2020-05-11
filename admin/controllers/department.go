//管理部门信息
package controllers

import (
	"github.com/kataras/iris/v12"
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
		d.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range departments {
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
		d.Render400(ctx, err, ctx.GetLocale().GetMessage("error.params_error"))
		return
	}
	valid := services.NewValidatorService(department)
	errString := valid.ResultError(ctx.GetLocale().Language())
	if errString != "" {
		d.Render400(ctx, nil, errString)
		return
	}
	department, err = d.Service.Create(department)
	if err != nil {
		d.Render500(ctx, err, err.Error())
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
		d.Render400(ctx, err, "")
		return
	}
	err = ctx.ReadJSON(&readData)
	if err != nil {
		d.Render400(ctx, err, err.Error())
		return
	}
	department, err := d.Service.First(uint(id))
	if err != nil {
		d.Render400(ctx, err, "")
		return
	}
	valid := services.NewValidatorService(readData)
	errString := valid.ResultError(ctx.GetLocale().Language())
	if errString != "" {
		d.Render400(ctx, nil, errString)
		return
	}
	err = d.Service.Update(department, readData)
	if err != nil {
		d.Render500(ctx, err, "")
		return
	}
	returnData, _ := d.StructToMap(department, ctx)
	d.RenderSuccessJson(ctx, returnData)
	return
}

func (d *DepartmentController) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		d.Render400(ctx, err, err.Error())
		return
	}
	err = d.Service.Delete(id)
	if err != nil {
		d.Render500(ctx, err, "")
	} else {
		d.RenderSuccessJson(ctx, iris.Map{})
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
