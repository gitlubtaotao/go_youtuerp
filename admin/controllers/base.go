package controllers

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/models"
	"youtuerp/services"
)

type BaseController struct {
}

func (b BaseController) RenderSuccessJson(ctx iris.Context, data interface{}) {
	m := iris.Map{
		"code": http.StatusOK,
		"data": data,
	}
	_, _ = ctx.JSON(m)
}

//render error to json
func (b BaseController) RenderErrorJson(ctx iris.Context, code int, err string) {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	ctx.StatusCode(code)
	_, _ = ctx.JSON(iris.Map{"code": code, "message": err})
}

func (b BaseController) RenderModuleColumn(ctx iris.Context, model interface{}) {
	column := services.NewColumnService(ctx.GetLocale())
	data, err := column.DefaultColumn(model)
	if err != nil {
		b.RenderErrorJson(ctx, http.StatusBadRequest, err.Error())
	}
	b.RenderSuccessJson(ctx, data)
}

//获取用户的的列设置
func (b BaseController) GetModelColumn(currentUser *models.Employee, model interface{}) []string {
	return []string{}
}

//根据token获取当前用户
func (b BaseController) CurrentUser(ctx iris.Context) (employee *models.Employee, err error) {
	tokenInfo := ctx.Values().Get("jwt").(*jwt.Token)
	foobar := tokenInfo.Claims.(jwt.MapClaims)
	email := foobar["email"].(string)
	phone := foobar["phone"].(string)
	service := services.NewEmployeeService()
	employee, err = service.FirstByPhoneAndEmail(phone, email)
	return
}

//将struct 转化成map
func (b *BaseController) StructToMap(currentObject interface{}, ctx iris.Context) (map[string]interface{}, error) {
	service := services.NewColumnService(ctx.GetLocale())
	return service.StructToMap(currentObject)
}

func (b *BaseController) GetPage(ctx iris.Context) uint {
	return uint(ctx.URLParamIntDefault("page", 1))
}
func (b *BaseController) GetPer(ctx iris.Context) uint {
	return uint(ctx.URLParamIntDefault("per", 20))
}
