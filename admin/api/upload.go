package api

import (
	"github.com/kataras/iris/v12"
	"youtuerp/tools/uploader"
)

type UploadApi struct {
	BaseApi
}

//上传文件
func (u *UploadApi) Upload(ctx iris.Context) {
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
