package controllers

import (
	"errors"
	"fmt"
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

//读取密码信息
type readPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SessionController struct {
	BaseController
	SService services.ISessionService
	EService services.IEmployeeService
}

func (s *SessionController) Login(ctx iris.Context) {
	var loginInfo login
	s.initSession()
	err := ctx.ReadJSON(&loginInfo)
	if err != nil {
		s.RenderErrorJson(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err = s.validateLogin(ctx, loginInfo); err != nil {
		s.RenderErrorJson(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//查询用户是否存在
	user, err := s.EService.FirstByPhoneOrEmail(loginInfo.UserName)
	if err != nil {
		conf.IrisApp.Logger().Error(err)
		s.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("devise.invalid"))
		return
	}
	fmt.Printf("sdsdsdsds")
	//对比password 是否正确
	if ok := s.SService.ValidatePassword(loginInfo.Password, user.EncryptedPassword); ok != nil {
		s.RenderErrorJson(ctx, http.StatusBadRequest,
			ctx.GetLocale().GetMessage("devise.invalid"))
		return
	}
	
	if err = s.updateLoginInfo(ctx, user); err != nil {
		s.RenderErrorJson(ctx, http.StatusBadRequest,
			ctx.GetLocale().GetMessage("devise.invalid"))
		return
	}
	tokenString, err := s.SService.JwtGenerateToken(map[string]interface{}{
		"email": user.Email,
		"phone": user.Phone,
	})
	if err != nil {
		s.RenderErrorJson(ctx, http.StatusBadRequest, err.Error())
	}
	s.RenderSuccessJson(ctx, iris.Map{"token": tokenString})
}

func (s *SessionController) Show(ctx iris.Context) {
	currentUser, err := s.CurrentUser(ctx)
	if err != nil {
		conf.IrisApp.Logger().Error(err)
		s.RenderErrorJson(ctx, http.StatusInternalServerError,
			ctx.GetLocale().GetMessage("error.inter_error"))
		return
	}
	userMap, err := s.StructToMap(currentUser, ctx)
	if err != nil {
		s.RenderErrorJson(ctx, http.StatusInternalServerError, err.Error())
	}
	s.RenderSuccessJson(ctx, s.handleUserInfo(userMap))
	return
}

func (s *SessionController) Logout(ctx iris.Context) {
	
	s.RenderSuccessJson(ctx, iris.Map{"message": "Logout is successful"})
}

func (s *SessionController) ResetToken(ctx iris.Context) {

}

func (s *SessionController) Update(ctx iris.Context) {
	s.initSession()
	var userInfo models.Employee
	err := ctx.ReadJSON(&userInfo)
	if err != nil {
		s.RenderErrorJson(ctx, http.StatusBadRequest, err.Error())
	}
	var passwordInfo readPassword
	err = ctx.ReadJSON(&passwordInfo)
	if err != nil {
		s.RenderErrorJson(ctx, http.StatusBadRequest, err.Error())
	}
	updateModel := models.Employee{Email: userInfo.Email, Name: userInfo.Name, Phone: userInfo.Phone, Address: userInfo.Address}
	fmt.Println(passwordInfo.Password)
	if passwordInfo.Password != "" {
		if passwordInfo.Password != passwordInfo.ConfirmPassword {
			s.RenderErrorJson(ctx, http.StatusBadRequest, ctx.GetLocale().GetMessage("error.password_error"))
			return
		}
		updateModel.EncryptedPassword, _ = s.SService.GeneratePassword(passwordInfo.Password)
	}
	currentUser, _ := s.CurrentUser(ctx)
	err = s.EService.UpdateRecord(currentUser, updateModel)
	if err != nil {
		s.RenderErrorJson(ctx, http.StatusBadRequest, err.Error())
		return
	}
	userMap, _ := s.StructToMap(currentUser, ctx)
	s.RenderSuccessJson(ctx, s.handleUserInfo(userMap))
}

//初始化session
func (s *SessionController) initSession() {
	s.SService = services.NewSessionService()
	s.EService = services.NewEmployeeService()
}

//保存当前登录用户的信息
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

//获取当前登录用户信息数据进行处理
func (s *SessionController) handleUserInfo(userInfo map[string]interface{}) map[string]interface{} {
	userInfo["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80"
	userInfo["roles"] = []string{"admin"}
	return userInfo
}

//验证登录用户的信息
func (s *SessionController) validateLogin(ctx iris.Context, login2 login) error {
	if login2.Password == "" {
		return errors.New(ctx.GetLocale().GetMessage("devise.invalid"))
	}
	if login2.Password == "" {
		return errors.New(ctx.GetLocale().GetMessage("devise.invalid"))
	}
	return nil
}
