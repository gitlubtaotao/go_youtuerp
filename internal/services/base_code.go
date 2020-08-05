package services

import (
	"errors"
	"strconv"
	"sync"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
)

type IBaseCode interface {
	//从redis 获取对应的字段的值
	HGetValue(key string, id interface{}, field string) string
	Update(id uint, code models.BaseDataCode, language string) error
	Delete(id uint) error
	Create(code models.BaseDataCode, language string) (models.BaseDataCode, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) (codes []models.BaseDataCode, total int64, err error)
	FindAllLevel() (levels []map[string]string, err error)
	FindCollect(key string) []map[string]string
}
type BaseCode struct {
	repo dao.IBaseCode
	BaseService
	mu sync.Mutex
}

func (b BaseCode) HGetValue(key string, id interface{}, field string) string {
	red := RedisService
	return red.HGetValue(models.BaseDataCode{}.TableName()+key, id, field)
}

func (b BaseCode) FindCollect(key string) []map[string]string {
	red := RedisService
	tableName := models.BaseDataCode{}.TableName()
	data := make([]map[string]string, 1)
	data = red.HCollectOptions(tableName + key)
	if len(data) > 0 {
		return data
	}
	records, _, _ := b.repo.Find(0, 0, map[string]interface{}{"code_name-eq": key}, []string{}, []string{}, false)
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, k := range records {
		go b.SaveRedisData(k)
		temp := map[string]string{"id": strconv.Itoa(int(k.ID)), "name": k.Name}
		data = append(data, temp)
	}
	return data
}

func (b BaseCode) Update(id uint, code models.BaseDataCode, language string) error {
	valid := NewValidatorService(code)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	if err := b.repo.Update(id, code); err != nil {
		return err
	}
	go b.SaveRedisData(code)
	return nil
}

func (b BaseCode) Delete(id uint) error {
	return b.repo.Delete(id)
}

func (b BaseCode) Create(code models.BaseDataCode, language string) (models.BaseDataCode, error) {
	valid := NewValidatorService(code)
	if message := valid.ResultError(language); message != "" {
		return models.BaseDataCode{}, errors.New(message)
	}
	result, err := b.repo.Create(code)
	if err != nil {
		return result, err
	}
	go b.SaveRedisData(result)
	return result, err
}

func (b BaseCode) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (codes []models.BaseDataCode, total int64, err error) {
	return b.repo.Find(per, page, filter, selectKeys, orders, true)
}

func (b BaseCode) FindAllLevel() (data []map[string]string, err error) {
	red := RedisService
	//data = red.HCollectOptions(models.BaseDataLevel{}.TableName())
	//if len(data) > 0 {
	//	return
	//}
	var levels []models.BaseDataLevel
	if levels, err = b.repo.FindAllLevel(); err != nil {
		return
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, k := range levels {
		_ = red.HSetValue(models.BaseDataLevel{}.TableName(), k.Code, map[string]interface{}{"code": k.Code, "name": k.Name, "en_name": k.EnName, "id": k.ID})
		temp := map[string]string{"code": k.Code, "name": k.Name, "en_name": k.EnName}
		data = append(data, temp)
	}
	return
}

func (b BaseCode) SaveRedisData(result models.BaseDataCode) {
	RedisService.HSetValue(models.BaseDataCode{}.TableName()+result.CodeName, result.ID, map[string]interface{}{
		"id":   result.ID,
		"name": result.Name,
	})
}

func NewBaseCode() IBaseCode {
	return BaseCode{repo: dao.NewBaseCode()}
}
