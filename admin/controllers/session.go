package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"time"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
	"youtuerp/tools"
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
	user, err := s.EService.FirstByPhoneOrEmail(loginInfo.UserName)
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
	
	if err = s.updateLoginInfo(ctx, user); err != nil {
		s.RenderJson(ctx, s.RenderErrorJson(http.StatusBadRequest,
			ctx.GetLocale().GetMessage("devise.invalid")))
		return
	}
	tokenString, err := s.SService.JwtGenerateToken(map[string]interface{}{
		"email": user.Email,
		"phone": user.Phone,
	})
	if err != nil {
		_, _ = ctx.JSON(s.RenderErrorJson(http.StatusBadRequest, err.Error()))
	}
	_, _ = ctx.JSON(s.RenderSuccessJson(iris.Map{"token": tokenString}))
}

func (s *SessionController) Show(ctx iris.Context) {
	currentUser, err := s.CurrentUser(ctx)
	if err != nil {
		conf.IrisApp.Logger().Error(err)
		_, _ = ctx.JSON(s.RenderErrorJson(http.StatusInternalServerError,
			ctx.GetLocale().GetMessage("error.inter_error")))
		return
	}
	userMap, err := s.StructToMap(currentUser, ctx)
	if err != nil {
		s.RenderJson(ctx, s.RenderErrorJson(http.StatusInternalServerError, err.Error()))
	}
	userMap["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80"
	_, _ = ctx.JSON(s.RenderSuccessJson(userMap))
	return
}

func (s *SessionController) Logout(ctx iris.Context) {
	
	_, _ = ctx.JSON(s.RenderSuccessJson(iris.Map{"message": "Logout is successful"}))
}

func (s *SessionController) ResetToken(ctx iris.Context) {

}

func (s *SessionController) initSession(ctx iris.Context) {
	s.SService = services.NewSessionService()
	s.EService = services.NewEmployeeService()
}

//保存用户信息
func (s *SessionController) updateLoginInfo(ctx iris.Context, employee *models.Employee) error {
	otherHelper := tools.OtherHelper{}
	ipAddress, _ := otherHelper.GetIPAddress(ctx.Request())
	updateColumn := map[string]interface{}{
		"sign_in_count":      employee.SignInCount + 1,
		"current_sign_in_at": time.Now(),
		"last_sign_in_at":    employee.CurrentSignInAt,
		"current_sign_in_ip": ipAddress,
		"last_sign_in_ip":    employee.CurrentSignInIp,
	}
	return s.EService.UpdateColumn(employee, updateColumn)
}
