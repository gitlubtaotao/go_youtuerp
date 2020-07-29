//管理部门信息
package api

import (
	"github.com/kataras/iris/v12"
	"youtuerp/global"
	"youtuerp/models"
	"youtuerp/services"
)

type Department struct {
	ctx iris.Context
	BaseApi
	service services.IDepartmentService
}

func (d *Department) GetColumn(ctx iris.Context) {
	d.RenderModuleColumn(ctx, models.ResponseDepartment{})
}

func (d *Department) Get(ctx iris.Context) {
	departments, total, err := d.service.Find(d.GetPer(ctx), d.GetPage(ctx), d.handlerGetParams(), []string{}, []string{}, true)
	if err != nil {
		d.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range departments {
		result, _ := d.StructToMap(v, ctx)
		dataArray = append(dataArray, result)
	}
	d.RenderSuccessJson(ctx, iris.Map{"data": dataArray, "total": total})
}

func (d *Department) Create(ctx iris.Context) {
	var (
		department models.Department
		err        error
	)
	err = ctx.ReadJSON(&department)
	if err != nil {
		d.Render400(ctx, err, "")
		return
	}
	valid := services.NewValidatorService(department)
	errString := valid.ResultError(ctx.GetLocale().Language())
	if errString != "" {
		d.Render400(ctx, nil, errString)
		return
	}
	department, err = d.service.Create(department)
	if err != nil {
		d.Render500(ctx, err, err.Error())
		return
	}
	d.RenderSuccessJson(ctx, d.handlerDataShow(department))
}

func (d *Department) Update(ctx iris.Context) {
	var (
		id       uint
		err      error
		readData models.Department
	)
	id, err = ctx.Params().GetUint("id")
	if err != nil {
		d.Render400(ctx, err, "")
		return
	}
	err = ctx.ReadJSON(&readData)
	if err != nil {
		d.Render400(ctx, err, err.Error())
		return
	}
	valid := services.NewValidatorService(readData)
	errString := valid.ResultError(ctx.GetLocale().Language())
	if errString != "" {
		d.Render400(ctx, nil, errString)
		return
	}
	err = d.service.Update(id, readData)
	if err != nil {
		d.Render500(ctx, err, "")
		return
	}
	d.RenderSuccessJson(ctx, d.handlerDataShow(readData))
	return
}

func (d *Department) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		d.Render400(ctx, err, err.Error())
		return
	}
	err = d.service.Delete(id)
	if err != nil {
		d.Render500(ctx, err, "")
	} else {
		d.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (d *Department) handlerDataShow(department interface{}) map[string]interface{} {
	data, _ := d.StructToMap(department, d.ctx)
	data["user_companies_name_nick"] = global.RedSetting.HGetCompany(data["user_company_id"], "name_nick")
	return data
}
func (d *Department) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["departments.name_cn-rCount"] = d.ctx.URLParamDefault("name_cn", "")
	searchColumn["departments.name_en-rCount"] = d.ctx.URLParamDefault("name_en", "")
	searchColumn["user_company_id-eq"] = d.ctx.URLParamDefault("user_company_id", "")
	return searchColumn
}

func (d *Department) Before(ctx iris.Context) {
	d.service = services.NewDepartmentService()
	d.ctx = ctx
	ctx.Next()
}
