package controllers
import (
	"github.com/kataras/iris/v12"
	"youtuerp/models"
	"youtuerp/services"
)

type ContactController struct {
	Ctx iris.Context
	BaseController
	Service services.IContactService
}


func (c *ContactController) Get() iris.Map {
	return c.RenderSuccessJson(map[string]interface{}{})
}

func (c *ContactController) GetColumn() iris.Map {
	return c.RenderColumnJson(models.Contact{}, c.Ctx.GetLocale())
}
