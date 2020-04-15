package controllers

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"net/http"
	"youtuerp/models"
	"youtuerp/services"
)

type BaseController struct {
}

// render success to json

func (b BaseController) RenderJson(ctx iris.Context, data iris.Map) {
	_, _ = ctx.JSON(data)
}
func (b BaseController) RenderSuccessJson(data interface{}) iris.Map {
	m := iris.Map{
		"code": http.StatusOK,
		"data": data,
	}
	return m
}

//render error to json
func (b BaseController) RenderErrorJson(code int, err string) iris.Map {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	return iris.Map{
		"code":    code,
		"message": err,
	}
}

func (b BaseController) RenderColumnJson(model interface{}, loader context.Locale) iris.Map {
	column := services.NewColumnService(loader)
	data, err := column.DefaultColumn(model)
	if err != nil {
		return b.RenderErrorJson(0, err.Error())
	}
	return b.RenderSuccessJson(iris.Map{
		"column": data,
	})
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
