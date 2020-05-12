package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type INumberSettingService interface {
	Create(numberSetting models.NumberSetting, language string) (models.NumberSetting, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string) (numberSettings []models.ResultNumberSetting,
		total uint, err error)
	Delete(id uint) error
}

type NumberSettingService struct {
	repo repositories.INumberSettingRepository
}

func (n NumberSettingService) Delete(id uint) error {
	return n.repo.Delete(id)
}

func (n NumberSettingService) Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string) (numberSettings []models.ResultNumberSetting,
	total uint, err error) {
	return n.repo.Find(per, page, filter, selectKeys, order, true)
}

func (n NumberSettingService) Create(numberSetting models.NumberSetting, language string) (result models.NumberSetting, err error) {
	valid := NewValidatorService(numberSetting)
	if message := valid.ResultError(language); message != "" {
		err = errors.New(message)
		return
	}
	return n.repo.Create(numberSetting)
}

func NewNumberSetting() INumberSettingService {
	return &NumberSettingService{repo: repositories.NewNumberSetting()}
}
