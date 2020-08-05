package services

import (
	"errors"
	"strconv"
	"sync"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
)

type IBasePort interface {
	FindCollect(key string) []map[string]string
	Update(id uint, code models.BaseDataPort, language string) error
	Delete(id uint) error
	Create(code models.BaseDataPort, language string) (models.BaseDataPort, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) (codes []models.BaseDataPort, total int64, err error)
}

type BasePort struct {
	BaseService
	repo dao.IBasePort
	mu   sync.Mutex
}

func (b BasePort) FindCollect(key string) []map[string]string {
	red := RedisService
	tableName := models.BaseDataPort{}.TableName()
	data := make([]map[string]string, 0)
	data = red.HCollectOptions(tableName + key)
	if len(data) > 0 {
		return data
	}
	records, _, _ := b.repo.Find(0, 0, map[string]interface{}{"type-eq": key}, []string{}, []string{}, false)
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, k := range records {
		go b.SaveRedisData(k)
		temp := map[string]string{"id": strconv.Itoa(int(k.ID)), "name": k.Name}
		data = append(data, temp)
	}
	return data
}
func (b BasePort) Update(id uint, code models.BaseDataPort, language string) error {
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

func (b BasePort) Delete(id uint) error {
	return b.repo.Delete(id)
}

func (b BasePort) Create(code models.BaseDataPort, language string) (models.BaseDataPort, error) {
	valid := NewValidatorService(code)
	if message := valid.ResultError(language); message != "" {
		return models.BaseDataPort{}, errors.New(message)
	}
	result, err := b.repo.Create(code)
	if err != nil {
		return result, err
	}
	go b.SaveRedisData(result)
	return result, err
}

func (b BasePort) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (codes []models.BaseDataPort, total int64, err error) {
	return b.repo.Find(per, page, filter, selectKeys, orders, true)
}

func (b BasePort) SaveRedisData(result models.BaseDataPort) {
	RedisService.HSetValue(models.BaseDataPort{}.TableName()+strconv.Itoa(int(result.Type)), result.ID, map[string]interface{}{
		"id":   result.ID,
		"name": result.Name,
	})
}

func NewBasePort() IBasePort {
	return &BasePort{repo: dao.NewBasePort()}
}
