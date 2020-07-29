package api

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

type BaseApi struct {
	renderError
}

//错误消息处理
type renderError struct {
}

var (
	red = redis.Redis{}
)

func (b *BaseApi) RenderSuccessJson(ctx iris.Context, data interface{}) {
	m := iris.Map{
		"code": http.StatusOK,
		"data": data,
	}
	_, _ = ctx.JSON(m)
}

func (b *BaseApi) RenderModuleColumn(ctx iris.Context, model interface{}) {
	column := services.NewColumnService(ctx.GetLocale())
	data, err := column.StructColumn(model)
	if err != nil {
		b.Render500(ctx, err, err.Error())
		return
	}
	b.RenderSuccessJson(ctx, data)
}

//获取用户的的列设置
func (b *BaseApi) GetCustomerColumn(currentUser *models.Employee, model interface{}) []string {
	return []string{}
}

//根据token获取当前用户
func (b *BaseApi) CurrentUser(ctx iris.Context) (employee *models.Employee, err error) {
	tokenInfo := ctx.Values().Get("jwt").(*jwt.Token)
	foobar := tokenInfo.Claims.(jwt.MapClaims)
	email := foobar["email"].(string)
	phone := foobar["phone"].(string)
	service := services.NewEmployeeService()
	employee, err = service.FirstByPhoneAndEmail(phone, email)
	return
}

//将struct 转化成map
func (b *BaseApi) StructToMap(currentObject interface{}, ctx iris.Context) (map[string]interface{}, error) {
	service := services.NewColumnService(ctx.GetLocale())
	return service.StructToMap(currentObject)
}

func (b *BaseApi) GetPage(ctx iris.Context) int {
	return ctx.URLParamIntDefault("page", 1)
}

func (b *BaseApi) GetPer(ctx iris.Context) int {
	return ctx.URLParamIntDefault("limit", 20)
}

func (b *BaseApi) HandlerFilterDate(filters map[string]interface{}, field string) {
	timeField, ok := filters[field]
	if !ok {
		return
	}
	stringTime := timeField.(string)
	timeArray := b.StringToDateRange(stringTime)
	if len(timeArray) == 2 {
		filters[field+"-gtEq"] = timeArray[0]
		filters[field+"-ltEq"] = timeArray[1]
	}
}

func (b *BaseApi) StringToDateRange(stringDate string) []time.Time {
	timeArray := strings.Split(stringDate, ",")
	return []time.Time{b.stringToDate(timeArray[0]), b.stringToDate(timeArray[1])}
}

// 将string转化成日期格式
func (b *BaseApi) stringToDate(strTime string) time.Time {
	result, err := tools.TimeHelper{}.StringToTime(strTime)
	if err != nil {
		golog.Errorf("string to date is error %v", err)
	}
	return result
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
