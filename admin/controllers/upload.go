package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/tools/uploader"
)

type UploadController struct {
	BaseController
}

//上传文件
func (u *UploadController) Upload(ctx iris.Context) {
	fmt.Printf("3333333")
	value, header, _ := ctx.FormFile("filer")
	up := uploader.NewQiNiuUploaderDefault()
	url, key, err := up.Upload(value, header)
	if err != nil {
		conf.IrisApp.Logger().Error(err)
		u.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("error.upload"))
	}
	url = up.PrivateReadURL(key)
	u.RenderSuccessJson(ctx, map[string]interface{}{
		"url": url,
		"key": key,
	})
}
