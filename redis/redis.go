package redis

import (
	"github.com/go-redis/redis/v7"
	"strconv"
	"youtuerp/conf"
	"youtuerp/tools/uploader"
)

//链接redis
func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "", // no password set
		DB:         0,  // use default DB
		MaxRetries: 2,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

type IRedis interface {
	HGet(key string, field string) string
	HGetAll(key string) map[string]string
}
type Redis struct {
}

func NewRedis() IRedis {
	return &Redis{}
}
func (r Redis) HGet(key string, field string) string {
	val, err := conf.ReisCon.HGet(key, field).Result()
	if err != nil {
		conf.IrisApp.Logger().Warnf("redis get key is %v, field is %v, error is %v", key, field, err)
		return ""
	}
	return val
}
func (r Redis) HGetAll(key string) map[string]string {
	val, err := conf.ReisCon.HGetAll(key).Result()
	if err != nil {
		conf.IrisApp.Logger().Warnf("redis get all key %v, error is %v", key, err)
		return map[string]string{}
	}
	return val
}

//获取公司的Logo
func CompanyLogo() string {
	red := NewRedis()
	value := red.HGet("base", "company_logo")
	if value == "" {
	}
	upload := uploader.NewQiNiuUploaderDefault()
	return upload.PrivateReadURL(value)
}

//获取数据安全控制
func DataSecurityControl() bool {
	red := NewRedis()
	value := red.HGet("base", "data_security_control")
	if value == "" {
	}
	if value == "false" {
		return false
	} else {
		return true
	}
}

//计费的计算方式
func ConversionMethod() (method string, remain int) {
	red := NewRedis()
	method = red.HGet("base", "conversion_method_calu")
	tempR := red.HGet("base", "conversion_method_remain")
	if method == "" || tempR == "" {
	
	}
	remain, _ = strconv.Atoi(tempR)
	return
}

//获取本位币计算方式
func SystemFinanceCurrency() int {
	red := NewRedis()
	value := red.HGet("finance", "system_finance_currency")
	if value == "" {
	}
	currency, _ := strconv.Atoi(value)
	return currency
}
