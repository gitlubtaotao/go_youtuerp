//管理员工信息

package api

import (
	"github.com/kataras/iris/v12"
	"sync"
	"youtuerp/models"
	"youtuerp/services"
)

type Employee struct {
	ctx     iris.Context
	service services.IEmployeeService
	BaseApi
}

func (e *Employee) GetColumn(ctx iris.Context) {
	e.RenderModuleColumn(ctx, models.ResponseEmployee{})
}

func (e *Employee) Get(ctx iris.Context) {
	var (
		sy        sync.Mutex
		sw        sync.WaitGroup
		result    []interface{}
		dataArray = make([]map[string]interface{}, 0)
	)
	employees, total, err := e.service.Find(e.GetPer(ctx), e.GetPage(ctx), e.handlerGetParams(), []string{}, []string{}, true)
	if err != nil {
		e.Render500(ctx, err, "")
		return
	}
	sw.Add(2)
	go func() {
		sy.Lock()
		defer sy.Unlock()
		for _, v := range employees {
			result, _ := e.StructToMap(v, ctx)
			dataArray = append(dataArray, result)
		}
		sw.Done()
	}()
	go func() {
		sy.Lock()
		defer sy.Unlock()
		result = e.selectDepartment()
		sw.Done()
	}()
	sw.Wait()
	e.RenderSuccessJson(ctx, iris.Map{
		"data":        dataArray,
		"total":       total,
		"departments": result,
	})
}

func (e *Employee) Create(ctx iris.Context) {
	var (
		employee     models.Employee
		readPassword models.ReadPassword
		err          error
	)
	if err = ctx.ReadJSON(&employee); err != nil {
		e.Render500(ctx, err, "")
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

func (e *Employee) Update(ctx iris.Context) {
	var (
		id           uint
		err          error
		readData     models.Employee
		readPassword models.ReadPassword
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		e.Render400(ctx, err, err.Error())
		return
	}
	if err = ctx.ReadJSON(&readData); err != nil {
		e.Render400(ctx, err, err.Error())
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
	readData.ID = id
	if err = e.service.UpdateRecord(id, readData); err != nil {
		e.Render500(ctx, err, "")
		return
	}
	returnData, _ := e.StructToMap(readData, ctx)
	e.RenderSuccessJson(ctx, returnData)
	return
}

func (e *Employee) Delete(ctx iris.Context) {
	var (
		id  int
		err error
	)
	if id, err = ctx.Params().GetInt("id"); err != nil {
		e.Render400(ctx, err, err.Error())
		return
	}
	if err = e.service.Delete(uint(id)); err != nil {
		e.Render500(ctx, err, "")
		return
	}
	e.RenderSuccessJson(ctx, iris.Map{})
}

func (e *Employee) Before(ctx iris.Context) {
	e.service = services.NewEmployeeService()
	e.ctx = ctx
	ctx.Next()
}

func (e *Employee) handlerGetParams() map[string]interface{} {
	searchColumn := make(map[string]interface{})
	searchColumn["name-rCount"] = e.ctx.URLParamDefault("name", "")
	searchColumn["phone-rCount"] = e.ctx.URLParamDefault("phone", "")
	searchColumn["email-rCount"] = e.ctx.URLParamDefault("email", "")
	searchColumn["users.user_company_id-eq"] = e.ctx.URLParamDefault("user_company_id", "")
	searchColumn["users.department_id-eq"] = e.ctx.URLParamDefault("department_id", "")
	return searchColumn
}

func (e *Employee) selectDepartment() []interface{} {
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

func (e *Employee) generatePassword(password string) (hashPassword string, err error) {
	session := services.NewSessionService()
	hashPassword, err = session.GeneratePassword(password)
	return
}
