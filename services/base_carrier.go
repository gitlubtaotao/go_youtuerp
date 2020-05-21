package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
)

type IBaseCarrier interface {
	Update(id uint, code models.BaseDataCarrier,language string) error
	Delete(id uint) error
	Create(code models.BaseDataCarrier,language string) (models.BaseDataCarrier, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) (codes []models.BaseDataCarrier, total uint, err error)
}

type BaseCarrier struct {
	BaseService
	repo repositories.IBaseCarrier
}

func (b BaseCarrier) Update(id uint, code models.BaseDataCarrier, language string) error {
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
	go redis.HSetValue("base_data_carriers", result.ID, map[string]interface{}{
		"id":   result.ID,
		"name": result.Name,
	})
	return result, err
}

func (b BaseCarrier) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) (codes []models.BaseDataCarrier, total uint, err error) {
	return b.repo.Find(per, page, filter, selectKeys, orders, true)
}

func NewBaseCarrier() IBaseCarrier {
	return &BaseCarrier{repo: repositories.NewBaseCarrier()}
}
