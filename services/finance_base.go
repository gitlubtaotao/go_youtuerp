package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type IFinanceBase interface {
	Update(id uint, record interface{}, language string) error
	Delete(id uint,model interface{}) error
	FindRate(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.FinanceRate, uint, error)
	FindFeeType(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.FinanceFeeType, uint, error)
	Create(record interface{}, language string) (interface{}, error)
}

type FinanceBase struct {
	BaseService
	repo repositories.IFinanceBase
}

func (f FinanceBase) Update(id uint, record interface{}, language string) error {
	valid := NewValidatorService(record)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	return f.repo.Update(id, record)
}

func (f FinanceBase) Delete(id uint, model interface{}) error {
	return f.repo.Delete(id, model)
}


func (f FinanceBase) FindFeeType(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.FinanceFeeType, uint, error) {
	return f.repo.FindFeeType(per, page, filter, selectKeys, orders)
}

func (f FinanceBase) FindRate(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.FinanceRate, uint, error) {
	return f.repo.FindRate(per, page, filter, selectKeys, orders)
}

func (f FinanceBase) Create(record interface{}, language string) (interface{}, error) {
	valid := NewValidatorService(record)
	if message := valid.ResultError(language); message != "" {
		return record, errors.New(message)
	}
	return f.repo.Create(record)
}

func NewFinanceBase() IFinanceBase {
	return &FinanceBase{repo: repositories.NewFinanceBase()}
}