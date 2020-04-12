package controllers

import "github.com/kataras/iris/v12"

type ContactController struct {
 Ctx iris.Context
 BaseController
 Service
}
