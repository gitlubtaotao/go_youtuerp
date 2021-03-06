package controllers

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"time"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/services"
	"youtuerp/tools"
)

type BaseController struct {
	renderError
}

var (
	red       = redis.Redis{}
)

func (b BaseController) RenderSuccessJson(ctx iris.Context, data interface{}) {
	m := iris.Map{
		"code": http.StatusOK,
		"data": data,
	}
	_, _ = ctx.JSON(m)
}

func (b BaseController) RenderModuleColumn(ctx iris.Context, model interface{}) {
	column := services.NewColumnService(ctx.GetLocale())
	data, err := column.StructColumn(model)
	if err != nil {
		b.Render500(ctx, err, err.Error())
		return
	}
	b.RenderSuccessJson(ctx, data)
}

//获取用户的的列设置
func (b BaseController) GetCustomerColumn(currentUser *models.Employee, model interface{}) []string {
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
	return uint(ctx.URLParamIntDefault("limit", 20))
}

func (b *BaseController) HandlerFilterDate(filters map[string]interface{}, field string) {
	timeField, ok := filters[field]
	if !ok {
		return
	}
	stringTime := timeField.(string)
	timeArray := strings.Split(stringTime, ",")
	if len(timeArray) == 2 {
		filters[field+"-gtEq"] = b.stringToDate(timeArray[0])
		filters[field+"-ltEq"] = b.stringToDate(timeArray[1])
	}
}

// 将string转化成日期格式
func (b *BaseController) stringToDate(strTime string) time.Time {
	result, err := tools.TimeHelper{}.StringToTime(strTime)
	if err != nil {
		golog.Errorf("string to date is error %v", err)
	}
	return result
}

//错误消息处理
type renderError struct {
}

//render 500 error
func (r renderError) Render500(ctx iris.Context, err error, message string) {
	if message == "" {
		message = ctx.GetLocale().GetMessage("error.inter_error")
	}
	if err != nil {
		golog.Errorf("render 500 error %v", err)
	}
	//ctx.StatusCode(http.StatusInternalServerError)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusInternalServerError, "message": message})
}

//render 400 error
func (r renderError) Render400(ctx iris.Context, err error, message string) {
	if message == "" {
		message = ctx.GetLocale().GetMessage("error.params_error")
	}
	if err != nil {
		golog.Warnf("render 400 warn %v", err)
	}
	//ctx.StatusCode(http.StatusBadRequest)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusBadRequest, "message": message})
}

func (r renderError) Render401(ctx iris.Context, err error, message string) {
	if message == "" {
		message = ctx.GetLocale().GetMessage("error.invalid_error")
	}
	if err != nil {
		golog.Warnf("render 401 warn %v", err)
	}
	//ctx.StatusCode(http.StatusBadRequest)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusUnauthorized, "message": message})
}
