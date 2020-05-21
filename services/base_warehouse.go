package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type IBaseWarehouse interface {
	Delete(id uint) error
	Update(id uint, code models.BaseWarehouse, language string) error
	Create(code models.BaseWarehouse, language string) (models.BaseWarehouse, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string, Orders []string) ([]models.BaseWarehouse, uint, error)
}
type BaseWarehouse struct {
	BaseService
	repo repositories.IBaseWarehouse
}

func (b BaseWarehouse) Delete(id uint) error {
	return b.repo.Delete(id)
}


func (b BaseWarehouse) Update(id uint, code models.BaseWarehouse, language string) error {
	valid := NewValidatorService(code)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	return b.repo.Update(id, code)
}

func (b BaseWarehouse) Create(code models.BaseWarehouse, language string) (models.BaseWarehouse, error) {
	valid := NewValidatorService(code)
	if message := valid.ResultError(language); message != "" {
		return models.BaseWarehouse{}, errors.New(message)
	}
	return b.repo.Create(code)
}

func (b BaseWarehouse) Find(per, page uint, filter map[string]interface{}, selectKeys []string, Orders []string) ([]models.BaseWarehouse, uint, error) {
	return b.repo.Find(per, page, filter, selectKeys, Orders, true)
}

func NewBaseWarehouse() IBaseWarehouse {
	return BaseWarehouse{repo: repositories.NewBaseWarehouse()}
}
