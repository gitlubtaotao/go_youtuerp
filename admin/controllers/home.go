package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

type HomeController struct {
	Ctx iris.Context
}

func (i *HomeController) Get() interface{} {
	return map[string]interface{}{
		"status": http.StatusOK,
		"path":   i.Ctx.Path(),
	}
}
