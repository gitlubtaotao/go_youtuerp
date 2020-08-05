package redisService

import "github.com/kataras/golog"

func (r RedisService) HSetValue(tableName string, id interface{}, value map[string]interface{}) error {
	key := r.CombineKey(tableName, id)
	if err := r.HMSet(key, value); err != nil {
		return err
	}
	return r.SAdd(tableName, id)
}

func HSetValue(table string, id interface{}, value map[string]interface{}) {
	red := NewRedis()
	if err := red.HSetValue(table, id, value); err != nil {
		golog.Warnf("set hash value is error current table is %v,current id is %v,error is %v",
			table, id, err)
	}
}
