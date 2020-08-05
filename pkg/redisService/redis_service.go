package redisService

import (
	"reflect"
	"strconv"
	"strings"
	"sync"
	"youtuerp/global"
)

// customer redis service
type RedisService struct {
	sy sync.Mutex
}

func NewRedis() RedisService {
	return RedisService{}
}

func (r RedisService) HSet(key string, field string, value interface{}) error {
	return global.RedisEngine.HSet(key, field, value).Err()
}

func (r RedisService) SAdd(key string, member ...interface{}) error {
	return global.RedisEngine.SAdd(key, member...).Err()
}

func (r RedisService) SMembers(key string) []string {
	return global.RedisEngine.SMembers(key).Val()
}

func (r RedisService) SRemove(key string, member ...interface{}) error {
	return global.RedisEngine.SRem(key, member...).Err()
}

func (r RedisService) HDelete(key string, field ...string) error {
	return global.RedisEngine.HDel(key, field...).Err()
}

func (r RedisService) HGet(key string, field string) (string, error) {
	if global.RedisEngine.HExists(key, field).Val() {
		return global.RedisEngine.HGet(key, field).Result()
	}
	return "", nil
}

//Key是否存在
func (r RedisService) KeyIsExist(key string) {
	global.RedisEngine.Exists(key).Val()
}

func (r RedisService) HMSet(key string, value ...interface{}) error {
	return global.RedisEngine.HMSet(key, value...).Err()
}

func (r RedisService) HMGet(key string, fields ...string) ([]interface{}, error) {
	return global.RedisEngine.HMGet(key, fields...).Result()
}

func (r RedisService) HGetAll(key string) (map[string]string, error) {
	val, err := global.RedisEngine.HGetAll(key).Result()
	if err != nil {
		return map[string]string{}, err
	}
	return val, nil
}

func (r RedisService) Scan(tableName string, cursor uint64, count int64) ([]string, uint64, error) {
	return global.RedisEngine.Scan(cursor, tableName+"*", count).Result()
}

func (r RedisService) CombineKey(tableName string, key interface{}) string {
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
	return strings.Join([]string{tableName, dst}, "")
}
