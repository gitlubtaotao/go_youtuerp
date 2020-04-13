package controllers

import (
	"github.com/kataras/iris/v12"
)

type SessionController struct {
	BaseController
}

func (s *SessionController) Get(ctx iris.Context) {

}
func (s *SessionController) Login(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{"message": "pong"})
}
