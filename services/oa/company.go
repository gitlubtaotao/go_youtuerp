package oa

import (
	"youtuerp/models"
	"youtuerp/repositories/oa"
)

type ICompanyService interface {
	FindCompany(per, page uint, attr map[string]interface{}, selectKeys []string, order []string) ([]*models.UserCompany, error)
}

type CompanyService struct {
	repo oa.ICompanyRepository
}

func (c *CompanyService) FindCompany(per, page uint, attr map[string]interface{}, selectKeys []string, order []string) ([]*models.UserCompany, error) {
	return c.repo.FindCompany(per, page, attr, selectKeys, order)
}

func NewCompanyService() ICompanyService {
	return &CompanyService{repo: oa.NewCompanyRepository()}
}
