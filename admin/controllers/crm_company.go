package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type CrmCompany struct {
	ctx iris.Context
	BaseController
	service services.ICrmCompanyService
}

func (c *CrmCompany) GetColumn(ctx iris.Context) {
	column := services.NewColumnService(ctx.GetLocale())
	data, err := column.StructColumn(models.CrmCompany{})
	if err != nil {
		c.Render500(ctx, err, err.Error())
		return
	}
	//type := ctx.URLParamDefault("type","customer")
	c.RenderSuccessJson(ctx, data)
}
func (c *CrmCompany) Get(ctx iris.Context) {
	CType := ctx.URLParamDefault("type", "customer")
	companies, total, err := c.service.Find(c.GetPer(ctx), c.GetPage(ctx), c.handlerGetParams(CType), []string{}, []string{})
	if err != nil {
		c.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range companies {
		result, _ := c.handleCompany(v)
		dataArray = append(dataArray, result)
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total,})
}

func (c *CrmCompany) Create(ctx iris.Context) {
	var (
		company     models.CrmCompany
		err         error
		currentUser *models.Employee
	)
	if err = ctx.ReadJSON(&company); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	currentUser, _ = c.CurrentUser(ctx)
	company.ParentId = currentUser.UserCompanyId
	company.Roles = append(company.Roles, models.Role{Name: models.RoleNameCreate, UserId: currentUser.ID})
	company, err = c.service.Create(company, ctx.GetLocale().Language())
	if err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	data, _ := c.handleCompany(company)
	c.RenderSuccessJson(ctx, data)
}

func (c *CrmCompany) Edit(ctx iris.Context) {
	var (
		id      uint
		err     error
		company models.CrmCompany
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if company, err = c.service.First(id, "Roles"); err != nil {
		c.Render500(ctx, err, err.Error())
		return
	}
	data, _ := c.StructToMap(company, ctx)
	delete(data, "created_at")
	delete(data, "updated_at")
	businessTypeName := data["business_type_name"].(string)
	data["business_type_name"] = strings.Split(businessTypeName, ",")
	c.RenderSuccessJson(ctx, data)
}

func (c *CrmCompany) Update(ctx iris.Context) {
	var (
		id      uint
		company models.CrmCompany
		err     error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&company); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if company, err = c.service.Update(id, company, ctx.GetLocale().Language()); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	data, _ := c.handleCompany(company)
	c.RenderSuccessJson(ctx, data)
}

func (c *CrmCompany) Delete(ctx iris.Context) {
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

func (c *CrmCompany) ChangeStatus(ctx iris.Context) {
	var (
		id     uint
		err    error
		status string
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	status = ctx.URLParam("status")
	if err = c.service.UpdateByMap(id, map[string]interface{}{"status": status}); err != nil {
		c.Render400(ctx, err, err.Error())
	} else {
		c.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (c *CrmCompany) ChangeType(ctx iris.Context) {
	var (
		id           uint
		err          error
		companyType  int
		userSalesman int
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	companyType, _ = ctx.URLParamInt("company_type")
	userSalesman, _ = ctx.URLParamInt("user_salesman_id")
	_ = c.service.UpdateByMap(id, map[string]interface{}{
		"company_type": companyType, "user_salesman_id": uint(userSalesman)})
}

func (c *CrmCompany) Show(ctx iris.Context) {
	var (
		id      uint
		err     error
		company models.CrmCompany
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	company, err = c.service.First(id, "Roles", "CrmContacts", "Accounts", "Invoices", "Address", "CrmTracks")
	if err != nil {
		c.Render500(ctx, err, "")
		return
	}
	data, _ := c.handleCompany(company)
	c.RenderSuccessJson(ctx, data)
}

// 获取公司的信息
func (c *CrmCompany) GetOperationInfo(ctx iris.Context) {
	var (
		id  uint
		err error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	c.RenderSuccessJson(ctx, c.service.GetOperationInfo(id))
}


func (c *CrmCompany) Before(ctx iris.Context) {
	c.ctx = ctx
	c.service = services.NewCrmCompanyService()
	ctx.Next()
}

func (c *CrmCompany) handleCompany(company models.CrmCompany) (data map[string]interface{}, err error) {
	enum := conf.Enum{Locale: c.ctx.GetLocale()}
	data, err = c.StructToMap(company, c.ctx)
	if err != nil {
		return
	}
	data["company_type_value"] = data["company_type"]
	data["company_type"] = enum.CompanyTypeText(data["company_type"])
	status := data["status"].(string)
	data["status"] = enum.DefaultText("user_companies_status.", status)
	data["status_value"] = status
	accountPeriod := data["account_period"]
	roles := data["roles"].([]models.Role)
	for i, v := range roles {
		roles[i].UserName = red.HGetRecord("users", v.UserId, "name")
		roles[i].Name = enum.DefaultText("roles_name.", v.Name)
	}
	data["roles"] = roles
	data["parent_id"] = red.HGetCompany(data["parent_id"], "name_nick")
	data["account_period"] = enum.DefaultText("user_companies_account_period.", accountPeriod)
	data["account_period_value"] = accountPeriod
	return data, err
}

func (c *CrmCompany) handlerGetParams(companyType string) map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["user_companies.name_nick-rCount"] = c.ctx.URLParamDefault("name_nick", "")
	searchColumn["user_companies.telephone-rCount"] = c.ctx.URLParamDefault("tel", "")
	searchColumn["user_companies.email-rCount"] = c.ctx.URLParamDefault("user_name", "")
	searchColumn["user_companies.account_period-rCount"] = c.ctx.URLParamDefault("user_email", "")
	searchColumn["user_companies.user_salesman_id-eq"] = c.ctx.URLParamDefault("create_id", "")
	searchColumn["user_companies.parent_id-eq"] = c.ctx.URLParamDefault("company_type", "")
	status := c.ctx.URLParamDefault("status", "")
	if status == "" {
		searchColumn["user_companies.status-in"] = []string{"rejected", "approved", "approving"}
	} else {
		searchColumn["user_companies.status-eq"] = status
	}
	if companyType == "customer" {
		searchColumn["user_companies.company_type-in"] = []int{1, 3}
	} else {
		searchColumn["user_companies.company_type-in"] = []int{2, 3}
	}
	return searchColumn
}
