package controllers

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/services"
)

type CrmUser struct {
	BaseController
	service services.ICrmUser
	ctx     iris.Context
	enum    conf.Enum
}

func (c *CrmUser) GetColumn(ctx iris.Context) {
	c.RenderModuleColumn(ctx, models.CrmUser{})
}

func (c *CrmUser) Get(ctx iris.Context) {
	filter := c.handleParams()
	users, total, err := c.service.Find(c.GetPer(ctx), c.GetPage(ctx), filter, []string{}, []string{})
	if err != nil {
		c.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	red := redis.NewRedis()
	for _, user := range users {
		data := c.handleData(red, user)
		dataArray = append(dataArray, data)
	}
	golog.Warnf("warn is %v","eeee")
	golog.Error("2222222")
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total,})
}

func (c *CrmUser) Create(ctx iris.Context) {
	var (
		user models.CrmUser
		err  error
	)
	if err = ctx.ReadJSON(&user); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if user, err = c.service.Create(user, ctx.GetLocale().Language()); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	red := redis.NewRedis()
	c.RenderSuccessJson(ctx, c.handleData(red, user))
}

func (c *CrmUser) Update(ctx iris.Context) {
	var (
		id   uint
		user models.CrmUser
		err  error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&user); err != nil {
		c.Render400(ctx, err, "")
		return
	}
	if err = c.service.Update(id, user, ctx.GetLocale().Language()); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	user.ID = id
	data := c.handleData(redis.NewRedis(), user)
	c.RenderSuccessJson(ctx, data)
}

func (c *CrmUser) Delete(ctx iris.Context) {
	var (
		id  uint
		err error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, "")
		return
	}
	if err = c.service.Delete(id); err != nil {
		c.Render500(ctx, err, "")
		return
	}
	c.RenderSuccessJson(ctx, iris.Map{})
}

func (c *CrmUser) Before(ctx iris.Context) {
	c.service = services.NewCrmUser()
	c.ctx = ctx
	c.enum = conf.Enum{Locale: c.ctx.GetLocale()}
	ctx.Next()
}

func (c *CrmUser) handleData(red redis.Redis, user models.CrmUser) map[string]interface{} {
	data, _ := c.StructToMap(user, c.ctx)
	data["user_company_id_value"] = data["user_company_id"]
	data["is_key_contact_value"] = data["is_key_contact"]
	if data["is_key_contact"].(bool) {
		data["is_key_contact"] = "是"
	} else {
		data["is_key_contact"] = "否"
	}
	data["sex_value"] = data["sex"]
	data["sex"] = c.enum.DefaultText("users_sex.", data["sex"])
	data["user_company_id_value"] = data["user_company_id"]
	data["user_company_id"] = red.GetCompany(data["user_company_id"], "name_nick")
	return data
}

func (c *CrmUser) handleParams() map[string]interface{} {
	data := make(map[string]interface{}, 0)
	data["name-rCount"] = c.ctx.URLParamDefault("name", "")
	data["phone-rCount"] = c.ctx.URLParamDefault("phone", "")
	data["email-rCount"] = c.ctx.URLParamDefault("email", "")
	data["user_company_id-eq"] = c.ctx.URLParamIntDefault("user_company_id", 0)
	return data
}
