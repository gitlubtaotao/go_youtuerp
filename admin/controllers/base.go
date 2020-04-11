package controllers

import (
	"github.com/kataras/iris/v12"
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

func (b BaseController) RenderColumnJson(model interface{}) iris.Map {
	column := services.NewColumnService()
	data, err := column.DefaultColumn(model)
	if err != nil {
		return b.RenderErrorJson(err.Error())
	}
	return b.RenderSuccessJson(iris.Map{
		"column": data,
	})
}





