package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/conf"
)

type SessionController struct {
	BaseController
}

func (s *SessionController) Get(ctx iris.Context) {
	err :=  ctx.View("session/new.html")
	conf.IrisApp.Logger().Error(err)
}
func (s *SessionController) Login(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{"message": "pong"})
}
