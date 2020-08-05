package services

import (
	"github.com/iris-contrib/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
	"sync"
	"time"
	"youtuerp/global"
)

type ISessionService interface {
	GeneratePassword(password string) (hashPassword string, err error)
	ValidatePassword(password string, hashPassword string) error
	JwtGenerateToken(data map[string]interface{}) (tokenString string, err error)
}

type SessionService struct {
	sy sync.Mutex
}

func (s *SessionService) GeneratePassword(password string) (hashPassword string, err error) {
	var bytes []byte
	bytes, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *SessionService) ValidatePassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

//通过jwt 插件生成对于的token
func (s *SessionService) JwtGenerateToken(data map[string]interface{}) (tokenString string, err error) {
	exp := global.AppSetting.ExpireTime
	mapClaims := jwt.MapClaims{
		"iss": "youtuerp",
		// 签发时间
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
	}
	s.sy.Lock()
	defer s.sy.Unlock()
	for k, v := range data {
		mapClaims[k] = v
	}
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, mapClaims)
	// 使用密码签名并获取完整的编码令牌作为字符串
	secret := global.AppSetting.TokenSecret
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func NewSessionService() ISessionService {
	return &SessionService{}
}
