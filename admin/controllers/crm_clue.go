package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"strconv"
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
	for _, v := range clues {
		result, _ := c.handleClue(v)
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
	data, _ := c.handleClue(clue)
	c.RenderSuccessJson(ctx, data)
}

func (c *CrmClue) Edit(ctx iris.Context) {
	var (
		id   uint
		err  error
		clue models.CrmClue
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if clue, err = c.service.First(id, false); err != nil {
		c.Render500(ctx, err, "")
		return
	}
	c.RenderSuccessJson(ctx, clue)
}

func (c *CrmClue) Delete(ctx iris.Context) {
	var (
		id  uint
		err error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if err = c.service.Delete(id); err != nil {
		c.Render500(ctx, err, "")
	} else {
		c.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (c *CrmClue) Show(ctx iris.Context) {
	var (
		id   uint
		err  error
		clue models.CrmClue
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if clue, err = c.service.First(id, true); err != nil {
		c.Render500(ctx, err, "")
	} else {
		data, _ := c.handleClue(clue)
		c.RenderSuccessJson(ctx, data)
	}
}

func (c *CrmClue) Update(ctx iris.Context) {
	var (
		id   uint
		err  error
		clue models.CrmClue
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&clue); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if err = c.service.Update(id, clue); err != nil {
		c.Render500(ctx, err, "")
		return
	}
	data, _ := c.handleClue(clue)
	c.RenderSuccessJson(ctx, data)
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
	searchColumn["crm_clues.company_type-eq"] = c.ctx.URLParamDefault("company_type", "")
	searchColumn["crm_clues.status-eq"] = c.ctx.URLParamDefault("status", "")
	return searchColumn
}

func (c *CrmClue) handleClue(clue models.CrmClue) (data map[string]interface{}, err error) {
	enum := conf.Enum{Locale: c.ctx.GetLocale()}
	data, err = c.StructToMap(clue, c.ctx)
	if err != nil {
		return
	}
	data["company_type_value"] = data["company_type"]
	data["company_type"] = enum.CompanyTypeText(data["company_type"])
	status := data["status"].(uint)
	statusStr := strconv.Itoa(int(status))
	data["status"] = enum.DefaultText("crm_clues_status.", statusStr)
	data["status_value"] = status
	return
}
