package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type IAddressService interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Address, language string) (models.Address, error)
	First(id uint) (models.Address, error)
	Create(account models.Address, language string) (models.Address, error)
	FindByOa(per, page int, filter map[string]interface{}, selectKeys []string,
		order []string) (accounts []models.Address, total int64, err error)
	FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) (accounts []models.Address, total int64, err error)
}

type AddressService struct {
	BaseService
	repo repositories.IAddressRepository
}

func (a AddressService) Delete(id uint) error {
	return a.repo.Delete(id)
}

func (a AddressService) UpdateById(id uint, updateContent models.Address, language string) (models.Address, error) {
	validate := NewValidatorService(updateContent)
	if message := validate.ResultError(language); message != "" {
		return models.Address{}, errors.New(message)
	}
	return a.repo.UpdateById(id, updateContent)
}

func (a AddressService) First(id uint) (models.Address, error) {
	return a.repo.First(id)
}

func (a AddressService) FindByOa(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string) (accounts []models.Address, total int64, err error) {
	return a.repo.FindByOa(per, page, filter, selectKeys, order)
}
func (a AddressService) FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (accounts []models.Address, total int64, err error) {
	return a.repo.FindByCrm(per, page, filter, selectKeys, orders)
}

func (a AddressService) Create(account models.Address, language string) (models.Address, error) {
	validate := NewValidatorService(account)
	if message := validate.ResultError(language); message != "" {
		return models.Address{}, errors.New(message)
	}
	return a.repo.Create(account)
}

func NewAddressService() IAddressService {
	return AddressService{repo: repositories.NewAddress()}
}
