package routes

import (
	"github.com/kataras/iris/v12"
	"youtuerp/admin/controllers"
)

//附件管理对于的路由
type Attachment struct {
	Route *Route
}

func (a *Attachment) Index() {
	attach := controllers.Attachment{}
	j := a.Route.jwtAccess()
	a.Route.app.PartyFunc("/attachments", func(c iris.Party) {
		c.Use(attach.Before)
		c.Post("/UploadOrder", j.Serve, attach.UploadOrder)
		c.Get("/{id:uint}/GetOrderFile", j.Serve, attach.GetOrderFile)
		c.Delete("/{id:uint}/delete", j.Serve, attach.Delete)
	})
}

func (a *Attachment) order() {

}
func NewAttachmentRoute(r *Route) *Attachment {
	return &Attachment{r}
}
