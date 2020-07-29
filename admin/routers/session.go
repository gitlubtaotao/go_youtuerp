package routers

import (
	"github.com/kataras/iris/v12/versioning"
	"youtuerp/admin/api"
)

type Session struct {
	Route *Routers
}

func (s *Session) Index() {
	r := s.Route
	j := r.jwtAccess()
	session := api.Session{}
	users := s.Route.app.Party("user/")
	{
		users.Post("/login", versioning.NewMatcher(versioning.Map{
			"1.0":               session.Login,
			versioning.NotFound: r.versionNotFound,
		}))
		users.Get("/info", j.Serve, session.Show)
		users.Post("/logout", j.Serve, session.Logout)
		users.Post("/update", j.Serve, session.Update)
		users.Post("/upload", j.Serve, session.UploadAvatar)
	}
}
func newRouteSession(r *Routers) *Session {
	return &Session{r}
}
