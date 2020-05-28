package services

import (
	"github.com/kataras/iris/v12/context"
	"strconv"
	"sync"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
)

type ICompanyService interface {
	FindCompany(per, page uint, filters map[string]interface{}, selectKeys []string, orders []string, isCount bool) ([]*models.UserCompany, uint, error)
	FirstCompany(id uint) (*models.UserCompany, error)
	FirstCompanyByRelated(id uint, related ...string) (models.UserCompany, error)
	AllCompany(filters map[string]interface{}, selectKeys []string, orders []string) ([]*models.UserCompany, uint, error)
	AllCompanyRedis() []map[string]string
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
	mu sync.Mutex
}

//获取所有的redis 数据
func (c *CompanyService) AllCompanyRedis() []map[string]string {
	red := redis.NewRedis()
	tableName := models.UserCompany{}.TableName()
	data := make([]map[string]string, 0)
	data = red.HCollectOptions(tableName)
	if len(data) > 0 {
		return data
	}
	records, _, err := c.AllCompany(map[string]interface{}{}, []string{}, []string{})
	if err != nil {
		return []map[string]string{}
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, k := range records {
		go c.SaveRedisData(*k)
		temp := map[string]string{"id": strconv.Itoa(int(k.ID)),
			"name_nick": k.NameNick,
			"name_cn":   k.NameCn,
			"name_en":   k.NameEn,
		}
		data = append(data, temp)
	}
	return data
}

func (c *CompanyService) FirstCompanyByRelated(id uint, related ...string) (models.UserCompany, error) {
	return c.repo.FirstCompanyByRelated(id, related...)
}

func (c *CompanyService) Delete(id uint) error {
	return c.repo.DeleteCompany(id)
}

func (c *CompanyService) Update(company *models.UserCompany, readData models.UserCompany) error {
	err := c.repo.UpdateCompany(company, readData)
	if err != nil {
		return err
	}
	go c.SaveRedisData(*company)
	return nil
}

func (c *CompanyService) FirstCompany(id uint) (*models.UserCompany, error) {
	return c.repo.FirstCompany(id)
}

func (c *CompanyService) Create(company models.UserCompany) (models.UserCompany, error) {
	data, err := c.repo.CreateCompany(company)
	if err != nil {
		return data, err
	}
	go c.SaveRedisData(data)
	return data, err
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

func (c *CompanyService) SaveRedisData(result models.UserCompany) {
	
	redis.HSetValue(models.Company{}.TableName(), result.ID, map[string]interface{}{
		"id":        result.ID,
		"name_nick": result.NameNick,
		"name_en":   result.NameEn,
		"name_cn":   result.NameCn,
	})
}

func NewCompanyService() ICompanyService {
	return &CompanyService{repo: repositories.NewCompanyRepository()}
}
