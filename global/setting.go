package global

import (
	"github.com/go-redis/redis/v7"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	redis2 "youtuerp/redis"
)

var (
	DataEngine    *gorm.DB            // database engine
	RedisEngine   *redis.Client       // redis engine
	IrisAppEngine *iris.Application   // iris app engine
	RedSetting    = redis2.NewRedis() // redis setting instance
)
