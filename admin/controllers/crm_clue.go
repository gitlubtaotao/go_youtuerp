package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/services"
)

type CrmClue struct {
	BaseController
	ctx     iris.Context
	service services.ICrmClueService
}

func (c *CrmClue) Before(ctx iris.Context) {
	c.service = services.NewCrmClueService()
	c.ctx = ctx
	ctx.Next()
}

