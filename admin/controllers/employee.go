//管理员工信息

package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/models"
	"youtuerp/services"
)

type EmployeeController struct {
	ctx     iris.Context
	service services.IEmployeeService
	BaseController
}

func (e *EmployeeController) GetColumn(ctx iris.Context) {
	e.RenderModuleColumn(ctx, models.ResultEmployee{})
}

func (e *EmployeeController) Get(ctx iris.Context) {
	employees, total, err := e.service.Find(e.GetPer(ctx), e.GetPage(ctx), e.handlerGetParams(), []string{}, []string{}, true)
	if err != nil {
		e.Render500(ctx, err, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range employees {
		result, _ := e.StructToMap(v, ctx)
		dataArray = append(dataArray, result)
	}
	result := e.selectDepartment()
	e.RenderSuccessJson(ctx, iris.Map{
		"data":        dataArray,
		"total":       total,
		"departments": result,
	})
	
}

func (e *EmployeeController) Create(ctx iris.Context) {
	var (
		employee     models.Employee
		readPassword models.ReadPassword
		err          error
	)
	if err = ctx.ReadJSON(&employee); err != nil {
		e.Render500(ctx, err, "")
		return
	}
	valid := services.NewValidatorService(employee)
	if message := valid.ResultError(ctx.GetLocale().Language()); message != "" {
		e.Render400(ctx, nil, message)
		return
	}
	_ = ctx.ReadJSON(&readPassword)
	//设置初始密码
	if readPassword.Password == "" {
		readPassword.Password = "qweqwe123"
	} else {
		if readPassword.Password != readPassword.ConfirmPassword {
			e.Render400(ctx, nil, ctx.GetLocale().GetMessage("error.password_error"))
			return
		}
	}
	employee.EncryptedPassword, _ = e.generatePassword(readPassword.Password)
	employee, err = e.service.Create(employee)
	if err != nil {
		e.Render500(ctx, err, ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	data, _ := e.StructToMap(employee, ctx)
	e.RenderSuccessJson(ctx, data)
}

func (e *EmployeeController) Edit(ctx iris.Context) {
	var (
		id       int
		employee *models.Employee
		err      error
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		e.Render400(ctx, err, "")
		return
	}
	if employee, err = e.service.First(uint(id)); err != nil {
		e.Render400(ctx, err, "")
		return
	}
	e.RenderSuccessJson(ctx, employee)
}

func (e *EmployeeController) Update(ctx iris.Context) {
	var (
		id           int
		err          error
		readData     models.Employee
		readPassword models.ReadPassword
		employee     *models.Employee
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		e.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&readData); err != nil {
		e.Render400(ctx, err, err.Error())
		return
	}
	valid := services.NewValidatorService(readData)
	errString := valid.ResultError(ctx.GetLocale().Language())
	if errString != "" {
		e.Render400(ctx, nil, errString)
		return
	}
	_ = ctx.ReadJSON(&readPassword)
	//设置初始密码
	if readPassword.Password != "" {
		if readPassword.Password != readPassword.ConfirmPassword {
			e.Render400(ctx, nil, ctx.GetLocale().GetMessage("error.password_error"))
			return
		}
		readData.EncryptedPassword, _ = e.generatePassword(readPassword.Password)
	}
	if employee, err = e.service.First(uint(id)); err != nil {
		e.Render400(ctx, err, "")
		return
	}
	if err = e.service.UpdateRecord(employee, readData); err != nil {
		e.Render500(ctx, err, "")
		return
	}
	returnData, _ := e.StructToMap(employee, ctx)
	e.RenderSuccessJson(ctx, returnData)
	return
}

func (e *EmployeeController) Delete(ctx iris.Context) {
	var (
		id  int
		err error
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		e.Render400(ctx, err, err.Error())
		return
	}
	if err = e.service.Delete(uint(id)); err != nil {
		e.Render500(ctx, err,"")
		return
	}
	e.RenderSuccessJson(ctx, iris.Map{})
}

func (e *EmployeeController) Before(ctx iris.Context) {
	e.service = services.NewEmployeeService()
	e.ctx = ctx
	ctx.Next()
}

func (e *EmployeeController) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["name-rCount"] = e.ctx.URLParamDefault("name", "")
	searchColumn["phone-rCount"] = e.ctx.URLParamDefault("phone", "")
	searchColumn["email-rCount"] = e.ctx.URLParamDefault("email", "")
	searchColumn["users.user_company_id-eq"] = e.ctx.URLParamDefault("user_company_id", "")
	searchColumn["users.department_id-eq"] = e.ctx.URLParamDefault("department_id", "")
	return searchColumn
}

func (e *EmployeeController) selectDepartment() []interface{} {
	service := services.NewDepartmentService()
	var selectKeys = []string{"departments.id", "user_company_id"}
	if e.ctx.GetLocale().Language() == "en" {
		selectKeys = append(selectKeys, "departments.name_en")
	} else {
		selectKeys = append(selectKeys, "departments.name_cn")
	}
	result, _, _ := service.FindAll(map[string]interface{}{}, selectKeys, []string{}, false)
	
	return result
}

func (e *EmployeeController) generatePassword(password string) (hashPassword string, err error) {
	session := services.NewSessionService()
	hashPassword, err = session.GeneratePassword(password)
	return
}
