package services

import (
	"errors"
	"sync"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
)

type IBaseCode interface {
	Update(id uint, code models.BaseDataCode, language string) error
	Delete(id uint) error
	Create(code models.BaseDataCode, language string) (models.BaseDataCode, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) (codes []models.BaseDataCode, total uint, err error)
	FindAllLevel() (levels []map[string]string, err error)
}
type BaseCode struct {
	repo repositories.IBaseCode
	BaseService
	mu sync.Mutex
}

func (b BaseCode) Update(id uint, code models.BaseDataCode, language string) error {
	valid := NewValidatorService(code)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	if err := b.repo.Update(id, code); err != nil {
		return err
	}
	redis.HSetValue("base_data_codes", id, map[string]interface{}{
		"id":   id,
		"name": code.Name,
	})
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
	redis.HSetValue("base_data_codes", result.ID, map[string]interface{}{
		"id":   result.ID,
		"name": result.Name,
	})
	return result, err
}

func (b BaseCode) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) (codes []models.BaseDataCode, total uint, err error) {
	return b.repo.Find(per, page, filter, selectKeys, orders, true)
}

func (b BaseCode) FindAllLevel() (data []map[string]string, err error) {
	red := redis.NewRedis()
	data = red.HCollectOptions("base_data_levels")
	if len(data) > 0 {
		return
	}
	var levels []models.BaseDataLevel
	if levels, err = b.repo.FindAllLevel(); err != nil {
		return
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, k := range levels {
		_ = red.HSetValue("base_data_levels", k.Code, map[string]interface{}{
			"code":    k.Code,
			"name":    k.Name,
			"en_name": k.EnName,
			"id":      k.ID,
		})
		temp := map[string]string{"code": k.Code, "name": k.Name, "en_name": k.EnName,}
		data = append(data, temp)
	}
	return
}

func NewBaseCode() IBaseCode {
	return BaseCode{repo: repositories.NewBaseCode()}
}
