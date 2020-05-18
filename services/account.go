package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type IAccountService interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Account, language string) (models.Account, error)
	First(id uint) (models.Account, error)
	Create(account models.Account, language string) (models.Account, error)
	FindByOa(per, page uint, filter map[string]interface{}, selectKeys []string,
		order []string) (accounts []models.Account, total uint, err error)
	FindByCrm(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) (accounts []models.Account, total uint, err error)
}

type AccountService struct {
	BaseService
	repo repositories.IAccountRepository
}

func (a AccountService) Delete(id uint) error {
	return a.repo.Delete(id)
}

func (a AccountService) UpdateById(id uint, updateContent models.Account, language string) (models.Account, error) {
	validate := NewValidatorService(updateContent)
	if message := validate.ResultError(language); message != "" {
		return models.Account{}, errors.New(message)
	}
	return a.repo.UpdateById(id, updateContent)
}

func (a AccountService) First(id uint) (models.Account, error) {
	return a.repo.First(id)
}

func (a AccountService) FindByOa(per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string) (accounts []models.Account, total uint, err error) {
	return a.repo.FindByOa(per, page, filter, selectKeys, order)
}
func (a AccountService) FindByCrm(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) (accounts []models.Account, total uint, err error) {
	return a.repo.FindByCrm(per, page, filter, selectKeys, orders)
}

func (a AccountService) Create(account models.Account, language string) (models.Account, error) {
	validate := NewValidatorService(account)
	if message := validate.ResultError(language); message != "" {
		return models.Account{}, errors.New(message)
	}
	return a.repo.Create(account)
}

func NewAccountService() IAccountService {
	return AccountService{repo: repositories.NewAccountRepository()}
}
