package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type CrmClue struct {
	BaseController
	ctx     iris.Context
	service services.ICrmClueService
}

func (c *CrmClue) GetColumn(ctx iris.Context) {
	c.RenderModuleColumn(ctx, models.CrmClue{})
}

func (c *CrmClue) Get(ctx iris.Context) {
	clues, total, err := c.service.Find(c.GetPer(ctx), c.GetPage(ctx), c.handlerGetParams(), []string{}, []string{})
	if err != nil {
		c.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	enum := conf.Enum{Locale: ctx.GetLocale()}
	for _, v := range clues {
		result, _ := c.StructToMap(v, ctx)
		result["company_type"] = enum.TransportTypeText(result["company_type"])
		dataArray = append(dataArray, result)
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total,})
}

func (c *CrmClue) Create(ctx iris.Context) {
	var (
		clue        models.CrmClue
		err         error
		currentUser *models.Employee
	)
	if err = ctx.ReadJSON(&clue); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if currentUser, err = c.CurrentUser(ctx); err != nil {
		c.Render401(ctx, err, "")
		return
	}
	clue.CreateId = currentUser.ID
	clue.UserCompanyId = uint(currentUser.UserCompanyId)
	if clue, err = c.service.Create(clue, ctx.GetLocale().Language()); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	data, _ := c.StructToMap(clue, ctx)
	c.RenderSuccessJson(ctx, data)
}

func (c *CrmClue) Delete(ctx iris.Context) {

}
func (c *CrmClue) Show(ctx iris.Context) {

}
func (c *CrmClue) Update(ctx iris.Context) {

}
func (c *CrmClue) Before(ctx iris.Context) {
	c.service = services.NewCrmClueService()
	c.ctx = ctx
	ctx.Next()
}

func (c *CrmClue) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["crm_clues.name_nick-rCount"] = c.ctx.URLParamDefault("name_nick", "")
	searchColumn["crm_clues.tel-rCount"] = c.ctx.URLParamDefault("tel", "")
	searchColumn["crm_clues.user_name-rCount"] = c.ctx.URLParamDefault("user_name", "")
	searchColumn["crm_clues.user_email-rCount"] = c.ctx.URLParamDefault("user_email", "")
	searchColumn["crm_clues.create_id-eq"] = c.ctx.URLParamDefault("create_id", "")
	searchColumn["crm_clues.company_type-eq"] = c.ctx.URLParamDefault("company_type","")
	return searchColumn
}
