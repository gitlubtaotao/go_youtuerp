package services

import (
	"github.com/iris-contrib/middleware/jwt"
	"sync"
	"time"
	"youtuerp/conf"
)

type ISessionService interface {
	GeneratePassword(password string) (hashPassword string)
	ValidatePassword(password string, hashPassword string) bool
	JwtGenerateToken(data map[string]interface{}) (tokenString string, err error)
}

type SessionService struct {
	sy sync.Mutex
}

func (s *SessionService) GeneratePassword(password string) (hashPassword string) {
	panic("implement me")
}

func (s *SessionService) ValidatePassword(password string, hashPassword string) bool {
	panic("implement me")
}

//通过jwt 插件生成对于的token
func (s *SessionService) JwtGenerateToken(data map[string]interface{}) (tokenString string, err error) {
	exp := conf.Configuration.ExpireTime
	s.sy.Lock()
	defer s.sy.Unlock()
	mapClaims := jwt.MapClaims{
		"iss": "youtuerp",
		// 签发时间
		"iat":         time.Now().Unix(),
		"expire_time": time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
	}
	for k, v := range data {
		mapClaims[k] = v
	}
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, mapClaims)
	// 使用密码签名并获取完整的编码令牌作为字符串
	secret := conf.Configuration.TokenSecret
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func NewSessionService() ISessionService {
	return &SessionService{}
}
