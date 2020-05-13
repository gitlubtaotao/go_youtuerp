package controllers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/models"
	"youtuerp/services"
)

//跟进记录的操作方法
type ClueTrack struct {
	BaseController
	service services.ICrmTrack
	ctx     iris.Context
}

func (c *ClueTrack) Get(ctx iris.Context) {
}

func (c *ClueTrack) Create(ctx iris.Context) {
	var (
		track models.CrmTrack
		err   error
	)
	if err = ctx.ReadJSON(&track); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	if track, err = c.service.Create(track, ctx.GetLocale().Language()); err != nil {
		c.Render400(ctx, err, err.Error())
		return
	}
	c.RenderSuccessJson(ctx, track)
}



func (c *ClueTrack) Before(ctx iris.Context) {
	c.ctx = ctx
	c.service = services.NewCrmTrack()
	ctx.Next()
}
