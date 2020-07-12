package services

import (
	"youtuerp/redis"
	"youtuerp/tools"
)

type IBaseService interface {
}

type BaseService struct {
}

var (
	toolOther = tools.OtherHelper{}
	red       = redis.Redis{}
	toolTime  = tools.TimeHelper{}
)
