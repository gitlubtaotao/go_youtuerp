package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/tools/uploader"
)

type UploadController struct {
	BaseController
}

//上传文件
func (u *UploadController) Upload(ctx iris.Context) {
	value, header, _ := ctx.FormFile("filer")
	up := uploader.NewQiNiuUploaderDefault()
	url, key, err := up.Upload(value, header)
	if err != nil {
		u.Render400(ctx, err, ctx.GetLocale().GetMessage("error.upload"))
		return
	}
	url = up.PrivateReadURL(key)
	u.RenderSuccessJson(ctx, map[string]interface{}{
		"url": url,
		"key": key,
	})
}
