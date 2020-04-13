package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

type ErrorsController struct {
	BaseController
}

func (e *ErrorsController) NotFound(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{"status": http.StatusNotFound})
}
func (e *ErrorsController) InternalServerError(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{"status": http.StatusInternalServerError})
}
