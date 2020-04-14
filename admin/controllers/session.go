package controllers

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type SessionController struct {
	BaseController
}

func (s *SessionController) Get(ctx iris.Context) {
}
func (s *SessionController) Login(ctx iris.Context) {
	var user login
	err := ctx.ReadJSON(&user)
	if err != nil {
		_, _ = ctx.JSON(s.RenderErrorJson(err.Error(), 0))
		return
	}
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})
	// 使用密码签名并获取完整的编码令牌作为字符串
	tokenString, _ := token.SignedString([]byte("My Secret"))
	
	_, _ = ctx.JSON(s.RenderSuccessJson(iris.Map{"token": tokenString}))
}

func (s *SessionController) Show(ctx iris.Context) {
	_, _ = ctx.JSON(s.RenderSuccessJson(iris.Map{"name": "sdsds"}))
	return
}
