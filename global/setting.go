package global

import (
	"github.com/go-redis/redis/v7"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var (
	DataEngine    *gorm.DB // database engine
	RedisEngine   *redis.Client
	IrisAppEngine *iris.Application
)
