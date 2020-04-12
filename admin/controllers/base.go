package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"youtuerp/services"
)

type BaseController struct {
}

// render success to json

func (b BaseController) RenderSuccessJson(data interface{}) iris.Map {
	m := iris.Map{
		"success": true,
		"data":    data,
	}
	return m
}

//render error to json
func (b BaseController) RenderErrorJson(err string) iris.Map {
	return iris.Map{
		"success": false,
		"error":   err,
	}
}

func (b BaseController) RenderColumnJson(model interface{},loader context.Locale) iris.Map {
	column := services.NewColumnService(loader)
	
	data, err := column.DefaultColumn(model)
	if err != nil {
		return b.RenderErrorJson(err.Error())
	}
	return b.RenderSuccessJson(iris.Map{
		"column": data,
	})
}





