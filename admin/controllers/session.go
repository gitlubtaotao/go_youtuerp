package controllers

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"net/http"
	"youtuerp/conf"
	"youtuerp/services"
)

type login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type SessionController struct {
	BaseController
	SService services.ISessionService
	EService services.IEmployeeService
}

func (s *SessionController) Login(ctx iris.Context) {
	var loginInfo login
	s.initSession(ctx)
	err := ctx.ReadJSON(&loginInfo)
	if err != nil {
		s.RenderJson(ctx, s.RenderErrorJson(http.StatusBadRequest, err.Error()))
		return
	}
	//查询用户是否存在
	user, err := s.EService.FirstByNameOrEmail(loginInfo.UserName)
	if err != nil {
		conf.IrisApp.Logger().Error(err)
		s.RenderJson(ctx, s.RenderErrorJson(http.StatusBadRequest,
			ctx.GetLocale().GetMessage("devise.invalid")))
		return
	}
	//对比password 是否正确
	if ok := s.SService.ValidatePassword(loginInfo.Password, user.EncryptedPassword); ok != nil {
		s.RenderJson(ctx, s.RenderErrorJson(http.StatusBadRequest,
			ctx.GetLocale().GetMessage("devise.invalid")))
		return
	}
	tokenString, err := s.SService.JwtGenerateToken(map[string]interface{}{})
	if err != nil {
		_, _ = ctx.JSON(s.RenderErrorJson(http.StatusBadRequest, err.Error()))
	}
	_, _ = ctx.JSON(s.RenderSuccessJson(iris.Map{"token": tokenString}))
}

func (s *SessionController) Show(ctx iris.Context) {
	_, _ = ctx.JSON(s.RenderSuccessJson(iris.Map{"name": "sdsds"}))
	return
}

func (s *SessionController) Logout(ctx iris.Context) {
	
	_, _ = ctx.JSON(s.RenderSuccessJson(iris.Map{"message": "Logout is successful"}))
}

func (s *SessionController) ResetToken(ctx iris.Context) {

}

func myAuthenticatedHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")
	foobar := user.Claims.(jwt.MapClaims)
	for key, value := range foobar {
		ctx.Writef("%s = %s", key, value)
	}
}

func (s *SessionController) initSession(ctx iris.Context) {
	s.SService = services.NewSessionService()
	s.EService = services.NewEmployeeService()
}
