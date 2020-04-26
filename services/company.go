package services

import (
	"github.com/kataras/iris/v12/context"
	"strconv"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ICompanyService interface {
	FindCompany(per, page uint, filters map[string]interface{}, selectKeys []string, orders []string, isCount bool) ([]*models.UserCompany, uint, error)
	FirstCompany(id uint) (*models.UserCompany, error)
	AllCompany(filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, uint, error)
	LimitCompany(limit uint, filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, uint, error)
	Create(company models.UserCompany) (models.UserCompany, error)
	Update(company *models.UserCompany, readData models.UserCompany) error
	Delete(id uint) error
	ShowTransportType(loader context.Locale, value interface{}, trans []map[string]interface{}) interface{}
	TransportTypeArrays(loader context.Locale) []map[string]interface{}
}

type CompanyService struct {
	repo repositories.ICompanyRepository
	BaseService
}

func (c *CompanyService) Delete(id uint) error {
	return c.repo.DeleteCompany(id)
}

func (c *CompanyService) Update(company *models.UserCompany, readData models.UserCompany) error {
	return c.repo.UpdateCompany(company, readData)
}

func (c *CompanyService) FirstCompany(id uint) (*models.UserCompany, error) {
	return c.repo.FirstCompany(id)
}

func (c *CompanyService) Create(company models.UserCompany) (models.UserCompany, error) {
	return c.repo.CreateCompany(company)
}

func (c *CompanyService) ShowTransportType(loader context.Locale, value interface{}, trans []map[string]interface{}) interface{} {
	if len(trans) == 0 {
		trans = c.TransportTypeArrays(loader)
	}
	for _, v := range trans {
		if v["data"] == value {
			return v["text"]
		}
	}
	return nil
}

func (c *CompanyService) TransportTypeArrays(loader context.Locale) []map[string]interface{} {
	data := make([]map[string]interface{}, 3)
	for _, value := range []int{1, 2, 3, 4} {
		data = append(data, map[string]interface{}{
			"data": value,
			"text": loader.GetMessage("transport_type." + strconv.Itoa(value)),
		})
	}
	return data
}

func (c *CompanyService) FindCompany(per, page uint, filters map[string]interface{}, selectKeys []string, orders []string, isCount bool) ([]*models.UserCompany, uint, error) {
	return c.repo.FindCompany(per, page, filters, selectKeys, orders, isCount)
}

func (c *CompanyService) AllCompany(filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, uint, error) {
	return c.repo.FindCompany(0, 0, filters, selectKeys, orders, false)
}

func (c *CompanyService) LimitCompany(limit uint, filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, uint, error) {
	return c.repo.FindCompany(limit, 0, filters, selectKeys, orders, false)
}

func NewCompanyService() ICompanyService {
	return &CompanyService{repo: repositories.NewCompanyRepository()}
}
