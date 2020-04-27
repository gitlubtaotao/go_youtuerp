package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/services"
)

type ContactController struct {
	Ctx iris.Context
	BaseController
	Service services.IContactService
}

//func (c *ContactController) Get() iris.Map {
//	//return c.RenderSuccessMap(c.Ctx, map[string]interface{}{})
//}

func (c *ContactController) GetColumn() iris.Map {
	return iris.Map{}
	//return c.RenderColumnMap(c.Ctx, models.Contact{})
}
