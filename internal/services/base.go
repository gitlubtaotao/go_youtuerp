package services

import (
	"youtuerp/pkg/redisService"
	"youtuerp/pkg/util"
)

type IBaseService interface {
}

type BaseService struct {
}

var (
	toolOther    = util.OtherHelper{}
	RedisService = redisService.NewRedis()
	toolTime     = util.TimeHelper{}
)
