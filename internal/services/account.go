package services

import (
	"errors"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
)

type IAccountService interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Account, language string) (models.Account, error)
	First(id uint) (models.Account, error)
	Create(account models.Account, language string) (models.Account, error)
	FindByOa(per, page int, filter map[string]interface{}, selectKeys []string,
		order []string) (accounts []models.Account, total int64, err error)
	FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) (accounts []models.Account, total int64, err error)
}

type AccountService struct {
	BaseService
	repo dao.IAccountRepository
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

func (a AccountService) FindByOa(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string) (accounts []models.Account, total int64, err error) {
	return a.repo.FindByOa(per, page, filter, selectKeys, order)
}
func (a AccountService) FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (accounts []models.Account, total int64, err error) {
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
	return AccountService{repo: dao.NewAccountRepository()}
}
