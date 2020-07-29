package redis

import (
	"github.com/kataras/golog"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
	"youtuerp/tools"
)

//保存公司redis信息
func (r Redis) HSetCompany(key string, id interface{}) error {
	result, err := r.findRecord("user_companies", map[string]interface{}{"id": id},
		[]string{"id", "name_nick", "code", "age", "amount", "account_period", "user_salesman_id"})
	if err != nil {
		return err
	}
	for _, v := range result {
		key := r.CombineKey(key, v["id"])
		if err = r.HMSet(key, v); err != nil {
			return err
		}
		_ = r.SAdd(key, v["id"])
	}
	return nil
}

func (r Redis) HSetRecord(tableName string, id interface{}, selectKeys []string) error {
	if len(selectKeys) == 0 {
		selectKeys = []string{"id", "name"}
	}
	result, err := r.findRecord(tableName, map[string]interface{}{"id": id}, selectKeys)
	if err != nil {
		return err
	}
	for _, v := range result {
		key := r.CombineKey(tableName, v["id"])
		if err = r.HMSet(key, v); err != nil {
			return err
		}
		_ = r.SAdd(tableName, v["id"])
	}
	return nil
}

func (r Redis) HSetValue(tableName string, id interface{}, value map[string]interface{}) error {
	key := r.CombineKey(tableName, id)
	if err := r.HMSet(key, value); err != nil {
		return err
	}
	return r.SAdd(tableName, id)
}

func (r Redis) HDeleteRecord(table string, id interface{}, field ...string) error {
	key := r.CombineKey(table, id)
	return r.HDelete(key, field...)
}

func (r Redis) findRecord(tableName string, filter map[string]interface{}, selectKey []string) (data []map[string]interface{}, err error) {
	repo := dao.NewSelectRepository()
	var result []models.ReadSelectResult
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

func HSetValue(table string, id interface{}, value map[string]interface{}) {
	red := NewRedis()
	if err := red.HSetValue(table, id, value); err != nil {
		golog.Warnf("set hash value is error current table is %v,current id is %v,error is %v",
			table, id, err)
	}
}
