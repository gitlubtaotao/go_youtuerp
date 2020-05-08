//管理员工信息

package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
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
		conf.IrisApp.Logger().Errorf("employee is err (%v)", err)
		e.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
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
		e.RenderErrorJson(ctx, 0, "")
		return
	}
	valid := services.NewValidatorService(employee)
	if message := valid.ResultError(ctx.GetLocale().Language()); message != "" {
		e.RenderErrorJson(ctx, 0, message)
		return
	}
	_ = ctx.ReadJSON(&readPassword)
	//设置初始密码
	if readPassword.Password == "" {
		readPassword.Password = "qweqwe123"
	} else {
		if readPassword.Password != readPassword.ConfirmPassword {
			e.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("error.password_error"))
			return
		}
	}
	employee.EncryptedPassword, _ = e.generatePassword(readPassword.Password)
	employee, err = e.service.Create(employee)
	if err != nil {
		conf.IrisApp.Logger().Error(err)
		e.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
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
		e.RenderErrorJson(ctx, 0, "")
		return
	}
	if employee, err = e.service.First(uint(id)); err != nil {
		e.RenderErrorJson(ctx, 0, "")
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
		e.RenderErrorJson(ctx, 0, "")
		return
	}
	if err = ctx.ReadJSON(&readData); err != nil {
		conf.IrisApp.Logger().Error(err)
		e.RenderErrorJson(ctx, 0, "")
		return
	}
	fmt.Printf("employee is %v", readData)
	valid := services.NewValidatorService(readData)
	errString := valid.ResultError(ctx.GetLocale().Language())
	if errString != "" {
		conf.IrisApp.Logger().Error(errString)
		e.RenderErrorJson(ctx, http.StatusBadRequest, errString)
		return
	}
	_ = ctx.ReadJSON(&readPassword)
	//设置初始密码
	if readPassword.Password != "" {
		if readPassword.Password != readPassword.ConfirmPassword {
			e.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("error.password_error"))
			return
		}
		readData.EncryptedPassword, _ = e.generatePassword(readPassword.Password)
	}
	if employee, err = e.service.First(uint(id)); err != nil {
		e.RenderErrorJson(ctx, 0, "")
		return
	}
	if err = e.service.UpdateRecord(employee, readData); err != nil {
		conf.IrisApp.Logger().Error(err)
		e.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
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
		e.RenderErrorJson(ctx, 0, "")
		return
	}
	if err = e.service.Delete(uint(id)); err != nil {
		conf.IrisApp.Logger().Error(err)
		e.RenderErrorJson(ctx, http.StatusInternalServerError, ctx.GetLocale().GetMessage("error.inter_error"))
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
