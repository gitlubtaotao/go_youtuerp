package redisService

import (
	"github.com/kataras/golog"
	"youtuerp/internal/models"
)

func (r RedisService) HGetCrm(id interface{}, field string) (value string) {
	if key := id.(uint); key == 0 {
		return ""
	}
	var err error
	value, err = r.getUserCompany(models.CrmCompany{}.RedisKey(), id, field)
	if err != nil {
		golog.Errorf("get crm redis is error %v", err)
	}
	return value
}

func (r RedisService) HGetCompany(id interface{}, field string) (value string) {
	var err error
	value, err = r.getUserCompany(models.Company{}.TableName(), id, field)
	if err != nil {
		golog.Errorf("get crm redis is error %v", err)
	}
	return value
}

func (r RedisService) HGetRecord(table string, id interface{}, field string) (value string) {
	var err error
	if field == "" {
		field = "name"
	}
	key := r.CombineKey(table, id)

	if value, err = r.HGet(key, field); err != nil {
		golog.Errorf("HGetValue is error table_name is %v,id is %v,error is %v", table, id, err)
		return ""
	}
	if value != "" {
		return value
	}
	if err := r.HSetRecord(table, id, []string{}); err != nil {
		golog.Errorf("set %v redis is error %v", table, err)
		return ""
	}
	value, _ = r.HGet(key, field)
	return
}

//保存公司redis信息
func (r RedisService) HSetCompany(key string, id interface{}) error {
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

func (r RedisService) HSetRecord(tableName string, id interface{}, selectKeys []string) error {
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

func (r RedisService) HDeleteRecord(table string, id interface{}, field ...string) error {
	key := r.CombineKey(table, id)
	return r.HDelete(key, field...)
}

func (r RedisService) findRecord(tableName string, filter map[string]interface{}, selectKey []string) (data []map[string]interface{}, err error) {
	//repo := dao.NewSelectRepository()
	//var result []models.ReadSelectResult
	//result, err = repo.FirstRecord(tableName, filter, selectKey)
	//if err != nil {
	//	return
	//}
	//r.sy.Lock()
	//for _, v := range result {
	//	data = append(data, util.StructToMap(v))
	//}
	//r.sy.Unlock()
	//return data, nil
	return nil, nil
}
