package redis

import (
	"github.com/go-redis/redis/v7"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"youtuerp/conf"
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

func (r Redis) HSet(key string, field string, value interface{}) error {
	return conf.ReisCon.HSet(key, field, value).Err()
}

func (r Redis) SAdd(key string, member ...interface{}) error {
	return conf.ReisCon.SAdd(key, member...).Err()
}

func (r Redis) SMembers(key string) []string {
	return conf.ReisCon.SMembers(key).Val()
}

func (r Redis) SRemove(key string, member ...interface{}) error {
	return conf.ReisCon.SRem(key, member...).Err()
}

func (r Redis) HDelete(key string, field ...string) error {
	return conf.ReisCon.HDel(key, field...).Err()
}

func (r Redis) HGet(key string, field string) (string, error) {
	if conf.ReisCon.HExists(key, field).Val() {
		return conf.ReisCon.HGet(key, field).Result()
	}
	return "", nil
}

//Key是否存在
func (r Redis) KeyIsExist(key string) {
	conf.ReisCon.Exists(key).Val()
}

func (r Redis) HMSet(key string, value ...interface{}) error {
	return conf.ReisCon.HMSet(key, value...).Err()
}

func (r Redis) HMGet(key string, fields ...string) ([]interface{}, error) {
	return conf.ReisCon.HMGet(key, fields...).Result()
}

func (r Redis) HGetAll(key string) (map[string]string, error) {
	val, err := conf.ReisCon.HGetAll(key).Result()
	if err != nil {
		return map[string]string{}, err
	}
	return val, nil
}

func (r Redis) Scan(tableName string, cursor uint64, count int64) ([]string, uint64, error) {
	return conf.ReisCon.Scan(cursor, tableName+"*", count).Result()
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
	return strings.Join([]string{tableName, dst}, "")
}

func NewRedis() Redis {
	return Redis{}
}

//获取公司的Logo
