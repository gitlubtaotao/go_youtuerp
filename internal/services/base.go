package services

import (
	"youtuerp/pkg/util"
	"youtuerp/redis"
)

type IBaseService interface {
}

type BaseService struct {
}

var (
	toolOther = util.OtherHelper{}
	red       = redis.Redis{}
	toolTime  = util.TimeHelper{}
)
