package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
)

type IBasePort interface {
	Update(id uint, code models.BaseDataPort,language string) error
	Delete(id uint) error
	Create(code models.BaseDataPort,language string) (models.BaseDataPort, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) (codes []models.BaseDataPort, total uint, err error)
}

type BasePort struct {
	BaseService
	repo repositories.IBasePort
}

func (b BasePort) Update(id uint, code models.BaseDataPort, language string) error {
	valid := NewValidatorService(code)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	if err := b.repo.Update(id, code); err != nil {
		return err
	}
	go redis.HSetValue("base_data_carriers", id, map[string]interface{}{
		"id":   id,
		"name": code.Name,
	})
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
	go redis.HSetValue("base_data_carriers", result.ID, map[string]interface{}{
		"id":   result.ID,
		"name": result.Name,
	})
	return result, err
}

func (b BasePort) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) (codes []models.BaseDataPort, total uint, err error) {
	return b.repo.Find(per, page, filter, selectKeys, orders, true)
}

func NewBasePort() IBasePort {
	return &BasePort{repo: repositories.NewBasePort()}
}
