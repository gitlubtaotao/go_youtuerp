package services

import (
	"github.com/kataras/iris/v12/context"
	"strconv"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ICompanyService interface {
	FindCompany(per, page uint, filters map[string]interface{}, selectKeys []string, orders []string, isCount bool) ([]*models.UserCompany, uint, error)
	AllCompany(filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, uint, error)
	LimitCompany(limit uint, filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, uint, error)
	CreateCompany(company models.UserCompany) (models.UserCompany, error)
	ShowTransportType(loader context.Locale, value interface{}, trans []map[string]interface{}) interface{}
	TransportTypeArrays(loader context.Locale) []map[string]interface{}
}

type CompanyService struct {
	repo repositories.ICompanyRepository
	BaseService
}

func (c *CompanyService) CreateCompany(company models.UserCompany) (models.UserCompany, error) {
	return c.repo.Create(company)
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
