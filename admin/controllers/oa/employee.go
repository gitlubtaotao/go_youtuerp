package oa

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

type EmployeeController struct {
	Ctx iris.Context
}

func (e *EmployeeController) Get() interface{} {
	return map[string]interface{}{
		"status": http.StatusOK,
	}
}
