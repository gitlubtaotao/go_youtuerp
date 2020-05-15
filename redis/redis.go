package redis

import (
	"bytes"
	"fmt"
	"github.com/go-redis/redis/v7"
	"reflect"
	"strconv"
	"sync"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/repositories"
	"youtuerp/tools"
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

type Redis struct {
	sy sync.Mutex
}

func (r Redis) FindRecord(tableName string, filter map[string]interface{}, selectKey []string) (data []map[string]interface{}, err error) {
	repo := repositories.NewSelectRepository()
	var result []models.SelectResult
	result, err = repo.FirstRecord(tableName, filter, selectKey)
	if err != nil {
		return
	}
	r.sy.Lock()
	for _, v := range result {
		data = append(data, tools.OtherHelper{}.StructToMap(v))
	}
	r.sy.Unlock()
	return data, nil
}

//通过查询数据库,将记录写入redis中
func (r Redis) HSetRecord(tableName string, filter map[string]interface{}, selectKeys []string) error {
	result, err := r.FindRecord(tableName, filter, selectKeys)
	fmt.Println(result,err)
	if err != nil {
		return err
	}
	for _, v := range result {
		key := r.CombineKey(tableName, v["id"])
		if err = r.HMSet(key, v); err != nil {
			return err
		}
	}
	return nil
}

func (r Redis) HSet(tableName string, key interface{}, field string, value interface{}) error {
	panic("implement me")
}

func (r Redis) HGet(tableName string, key interface{}, field string) string {
	dst := r.CombineKey(tableName, key)
	if conf.ReisCon.HExists(dst, field).Val() {
		val, err := conf.ReisCon.HGet(dst, field).Result()
		if err != nil {
			conf.IrisApp.Logger().Warnf("redis get key is %v, field is %v, error is %v", key, field, err.Error())
			return ""
		}
		return val
	}
	return ""
}

//Key是否存在
func (r Redis) KeyIsExist(tableName, key string) {
	conf.ReisCon.Exists(r.CombineKey(tableName, key)).Val()
}

func (r Redis) HMSet(key string, value ...interface{}) error {
	return conf.ReisCon.HMSet(key, value...).Err()
}

func (r Redis) HMGet(key string, fields ...string) ([]interface{}, error) {
	return conf.ReisCon.HMGet(key, fields...).Result()
}

func (r Redis) HGetAll(tableName string, key string) map[string]string {
	val, err := conf.ReisCon.HGetAll(r.CombineKey(tableName, key)).Result()
	if err != nil {
		conf.IrisApp.Logger().Warnf("redis get all key %v, error is %v", key, err)
		return map[string]string{}
	}
	return val
}

func (r Redis) CombineKey(tableName string, key interface{}) string {
	var dst string
	ty := reflect.TypeOf(key)
	switch ty.Kind() {
	case reflect.String:
		dst = key.(string)
	case reflect.Int:
		dst = strconv.Itoa(key.(int))
	case reflect.Uint:
		dst = strconv.Itoa(int(key.(uint)))
	}
	var b bytes.Buffer
	b.WriteString(tableName)
	b.WriteString(dst)
	return b.String()
}

func NewRedis() Redis {
	return Redis{}
}

//获取公司的Logo
func CompanyLogo() string {
	red := NewRedis()
	value := red.HGet("system_settings", "base", "company_logo")
	if value == "" {
	}
	upload := uploader.NewQiNiuUploaderDefault()
	return upload.PrivateReadURL(value)
}

//获取数据安全控制
func DataSecurityControl() bool {
	red := NewRedis()
	value := red.HGet("system_settings", "base", "data_security_control")
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
	table := "system_settings"
	method = red.HGet(table, "base", "conversion_method_calu")
	tempR := red.HGet(table, "base", "conversion_method_remain")
	if method == "" || tempR == "" {
	}
	remain, _ = strconv.Atoi(tempR)
	return
}

//获取本位币计算方式
func SystemFinanceCurrency() int {
	red := NewRedis()
	value := red.HGet("system_settings", "finance", "system_finance_currency")
	if value == "" {
	}
	currency, _ := strconv.Atoi(value)
	return currency
}
