package api

import (
	"errors"
	"github.com/kataras/iris/v12"
	"strconv"
	"youtuerp/models"
	"youtuerp/services"
	"youtuerp/tools/uploader"
)

type Attachment struct {
	BaseApi
	service services.IAttachment
}

//上传订单附件
func (a *Attachment) UploadOrder(ctx iris.Context) {
	label := ctx.FormValue("label")
	sourceId := ctx.FormValue("source_id")
	if label == "" {
		label = "internal"
	}
	if sourceId == "" {
		err := errors.New(ctx.GetLocale().GetMessage("error.params_error"))
		a.Render400(ctx, err, err.Error())
		return
	}
	file, header, err := ctx.FormFile("file")
	if err != nil {
		a.Render400(ctx, err, err.Error())
		return
	}
	up := uploader.NewQiNiuUploaderDefault()
	_, key, err := up.Upload(file, header)
	if err != nil {
		a.Render500(ctx, err, "")
		return
	}
	sourceID, _ := strconv.Atoi(sourceId)
	attachment := models.Attachment{
		Name: header.Filename, Size: header.Size,
		TypeOf: "", Key: key, Url: up.PrivateReadURL(key),
		Label: label, SourceID: uint(sourceID),
		SourceType: "order_masters",
	}
	attachment, err = a.service.Create(attachment)
	if err != nil {
		a.Render500(ctx, err, "")
	} else {
		a.RenderSuccessJson(ctx, attachment)
	}
}

//获取订单操作的附件
func (a *Attachment) GetOrderFile(ctx iris.Context) {
	orderMasterId, err := ctx.Params().GetUint("id")
	if err != nil {
		a.Render400(ctx, err, "")
		return
	}
	internalFiles, err := a.service.FindAll(map[string]interface{}{"source_id": orderMasterId, "source_type": "order_masters", "label": "internal"})
	if err != nil {
		a.Render500(ctx, err, "")
		return
	}
	externalFiles, err := a.service.FindAll(map[string]interface{}{"source_id": orderMasterId, "source_type": "order_masters", "label": "external"})
	if err != nil {
		a.Render500(ctx, err, "")
		return
	}
	_, _ = ctx.JSON(iris.Map{"code": 200, "internal": internalFiles, "external": externalFiles})
}

func (a *Attachment) Delete(ctx iris.Context) {
	var (
		id  uint
		err error
	)
	id, err = ctx.Params().GetUint("id")
	if err != nil {
		a.Render400(ctx, err, "")
		return
	}
	err = a.service.Delete(id)
	if err != nil {
		a.Render500(ctx, err, "")
	} else {
		a.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (a *Attachment) Before(ctx iris.Context) {
	a.service = services.NewAttachment()
	ctx.Next()
}
