package services

import (
	"youtuerp/models"
	"youtuerp/repositories"
)



type ICompanyService interface {
	FindCompany(per, page uint, filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, error)
	AllCompany(filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, error)
	LimitCompany(limit uint, filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, error)
}

type CompanyService struct {
	repo repositories.ICompanyRepository
}

func (c *CompanyService) FindCompany(per, page uint, filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, error) {
	return c.repo.FindCompany(per, page, filters, selectKeys, orders)
}

func (c *CompanyService) AllCompany(filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, error) {
	return c.repo.FindCompany(0, 0, filters, selectKeys, orders)
}


func (c *CompanyService) LimitCompany(limit uint, filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, error) {
	return c.repo.FindCompany(limit, 0, filters, selectKeys, orders)
}

func NewCompanyService() ICompanyService {
	return &CompanyService{repo: repositories.NewCompanyRepository()}
}
