package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ICrmCompanyService interface {
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) ([]models.CrmCompany, uint, error)
	Create(company models.CrmCompany, language string) (models.CrmCompany, error)
}

type CrmCompanyService struct {
	repo repositories.ICrmCompany
	BaseService
}

func (c CrmCompanyService) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) ([]models.CrmCompany, uint, error) {
	return c.repo.Find(per, page, filter, selectKeys, orders, true)
}

func (c CrmCompanyService) Create(company models.CrmCompany, language string) (models.CrmCompany, error) {
	valid := NewValidatorService(company)
	if message := valid.ResultError(language); message != "" {
		return company, errors.New(message)
	}
	//
	//将记录设置缓存
	return c.repo.Create(company)
}


func NewCrmCompanyService() ICrmCompanyService {
	return &CrmCompanyService{repo: repositories.NewCrmCompany()}
}
