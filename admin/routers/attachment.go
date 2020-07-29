package routers

import (
	"github.com/kataras/iris/v12"
	"youtuerp/admin/api"
)

//附件管理对于的路由
type Attachment struct {
	route *Routers
}

func (a *Attachment) Index() {
	attach := api.Attachment{}
	j := a.route.jwtAccess()
	a.route.app.PartyFunc("/attachments", func(c iris.Party) {
		c.Use(attach.Before)
		c.Post("/UploadOrder", j.Serve, attach.UploadOrder)
		c.Get("/{id:uint}/GetOrderFile", j.Serve, attach.GetOrderFile)
		c.Delete("/{id:uint}/delete", j.Serve, attach.Delete)
	})
}

func (a *Attachment) order() {

}
func newAttachmentRoute(r *Routers) *Attachment {
	return &Attachment{route: r}
}
