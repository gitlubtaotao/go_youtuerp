package api

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

type HomeApi struct {
	Ctx iris.Context
}

func (i *HomeApi) Get() interface{} {
	return map[string]interface{}{
		"status": http.StatusOK,
		"path":   i.Ctx.Path(),
	}
}
