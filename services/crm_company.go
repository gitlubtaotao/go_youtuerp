package services

import (
	"errors"
	"strings"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ICrmCompanyService interface {
	GetOperationInfo(id uint) string
	UpdateByMap(id uint, attr map[string]interface{}) error
	Delete(id uint) error
	Update(id uint, company models.CrmCompany, language string) (models.CrmCompany, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) ([]models.CrmCompany, uint, error)
	Create(company models.CrmCompany, language string) (models.CrmCompany, error)
	First(id uint, preload ...string) (models.CrmCompany, error)
}

type CrmCompanyService struct {
	repo repositories.ICrmCompany
	BaseService
}

func (c CrmCompanyService) GetOperationInfo(id uint) string {
	var (
		operationInfo string
		buffer        strings.Builder
	)
	crm, err := c.repo.First(id)
	if err != nil {
		return ""
	}
	if crm.FrequentlyUseInfo != "" {
		return crm.FrequentlyUseInfo
	}
	buffer.WriteString(crm.NameEn)
	buffer.WriteString("\n")
	buffer.WriteString(crm.EnAddress)
	buffer.WriteString("\n")
	buffer.WriteString("Tel:")
	buffer.WriteString(crm.Telephone)
	buffer.WriteString("\n")
	buffer.WriteString("Email:")
	buffer.WriteString(crm.Email)
	operationInfo = buffer.String()
	go c.repo.UpdateByMap(id, map[string]interface{}{"frequently_use_info": operationInfo})
	return operationInfo
}

func (c CrmCompanyService) UpdateByMap(id uint, attr map[string]interface{}) error {
	return c.repo.UpdateByMap(id, attr)
}

func (c CrmCompanyService) Delete(id uint) error {
	return c.repo.Delete(id)
}

func (c CrmCompanyService) Update(id uint, company models.CrmCompany, language string) (models.CrmCompany, error) {
	//处理roles
	
	valid := NewValidatorService(company)
	if message := valid.ResultError(language); message != "" {
		return models.CrmCompany{}, errors.New(message)
	}
	return c.repo.Update(id, company)
}

func (c CrmCompanyService) First(id uint, preload ...string) (models.CrmCompany, error) {
	return c.repo.First(id, preload...)
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
