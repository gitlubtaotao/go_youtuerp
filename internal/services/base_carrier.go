package services

import (
	"errors"
	"strconv"
	"sync"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
	"youtuerp/redis"
)

type IBaseCarrier interface {
	FindCollect(key string) []map[string]string
	Update(id uint, code models.BaseDataCarrier, language string) error
	Delete(id uint) error
	Create(code models.BaseDataCarrier, language string) (models.BaseDataCarrier, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) (codes []models.BaseDataCarrier, total int64, err error)
}

type BaseCarrier struct {
	BaseService
	repo dao.IBaseCarrier
	mu   sync.Mutex
}

func (b BaseCarrier) FindCollect(key string) []map[string]string {
	red := redis.NewRedis()
	tableName := models.BaseDataCarrier{}.TableName()
	data := make([]map[string]string, 0)
	data = red.HCollectOptions(tableName + key)
	if len(data) > 0 {
		return data
	}
	idKey, _ := strconv.Atoi(key)
	records, _, _ := b.repo.Find(0, 0, map[string]interface{}{"type-eq": idKey}, []string{}, []string{}, false)
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, k := range records {
		go b.SaveRedisData(k)
		temp := map[string]string{"id": strconv.Itoa(int(k.ID)), "name": k.Name}
		data = append(data, temp)
	}
	return data
}
func (b BaseCarrier) Update(id uint, code models.BaseDataCarrier, language string) error {
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

func (b BaseCarrier) Delete(id uint) error {
	return b.repo.Delete(id)
}

func (b BaseCarrier) Create(code models.BaseDataCarrier, language string) (models.BaseDataCarrier, error) {
	valid := NewValidatorService(code)
	if message := valid.ResultError(language); message != "" {
		return models.BaseDataCarrier{}, errors.New(message)
	}
	result, err := b.repo.Create(code)
	if err != nil {
		return result, err
	}
	go b.SaveRedisData(result)
	return result, err
}

func (b BaseCarrier) SaveRedisData(result models.BaseDataCarrier) {
	redis.HSetValue(models.BaseDataCarrier{}.TableName()+strconv.Itoa(int(result.Type)), result.ID, map[string]interface{}{
		"id":   result.ID,
		"name": result.Name,
	})
}

func (b BaseCarrier) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (codes []models.BaseDataCarrier, total int64, err error) {
	return b.repo.Find(per, page, filter, selectKeys, orders, true)
}

func NewBaseCarrier() IBaseCarrier {
	return &BaseCarrier{repo: dao.NewBaseCarrier()}
}
