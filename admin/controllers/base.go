package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"net/http"
	"youtuerp/services"
)

type BaseController struct {
}

// render success to json

func (b BaseController) RenderSuccessJson(data interface{}) iris.Map {
	m := iris.Map{
		"code": http.StatusOK,
		"data": data,
	}
	return m
}

//render error to json
func (b BaseController) RenderErrorJson(err string, code int) iris.Map {
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
		return b.RenderErrorJson(err.Error(),0)
	}
	return b.RenderSuccessJson(iris.Map{
		"column": data,
	})
}
