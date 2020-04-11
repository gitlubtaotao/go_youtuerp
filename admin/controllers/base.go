package controllers

import (
	"github.com/kataras/iris/v12"
)

type BaseController struct {
}

func (b BaseController) RenderSuccessJson(data interface{}) iris.Map {
	m := iris.Map{
		"success": true,
		"data":    data,
	}
	return m
}

func (b BaseController) RenderErrorJson(err string) iris.Map {
	return iris.Map{
		"success": false,
		"error":   err,
	}
}
