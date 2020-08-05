package dao

import (
	"gorm.io/gorm"
	"youtuerp/internal/models"
)

type ICompanyRepository interface {
	/*查询公司信息
	per: 每次分页的记录数
	page: 当前页数
	attr: 查询公司信息
	*/
	FindCompany(per, page int, attr map[string]interface{}, selectKeys []string, order []string, isCount bool) (companies []models.UserCompany,
		total int64, err error)
	FirstCompany(id uint) (models.UserCompany, error)
	FirstCompanyByRelated(id uint, related ...string) (models.UserCompany, error)
	CreateCompany(company models.UserCompany) (models.UserCompany, error)
	UpdateCompany(id uint, readData models.UserCompany) error
	DeleteCompany(id uint) error
}

type CompanyRepository struct {
	BaseRepository
}

func (c CompanyRepository) FirstCompanyByRelated(id uint, related ...string) (models.UserCompany, error) {
	sqlCon := global.DataEngine
	var company models.UserCompany
	for _, re := range related {
		sqlCon = sqlCon.Preload(re)
	}
	err := sqlCon.Find(&company, "id = ?", id).Error
	return company, err
}

func (c CompanyRepository) DeleteCompany(id uint) error {
	var readData models.UserCompany
	return global.DataEngine.Find(&readData).Delete(&readData).Error
}

func (c CompanyRepository) UpdateCompany(id uint, readData models.UserCompany) error {
	return global.DataEngine.Model(&models.UserCompany{ID: id}).Updates(readData).Error
}

func (c CompanyRepository) FirstCompany(id uint) (models.UserCompany, error) {
	var readData models.UserCompany
	err := global.DataEngine.First(&readData, id).Error
	return readData, err
}

func (c CompanyRepository) CreateCompany(company models.UserCompany) (models.UserCompany, error) {
	err := global.DataEngine.Create(&company).Error
	return company, err
}

//
////默认查询company
func (c CompanyRepository) DefaultScope(temp *gorm.DB) *gorm.DB {
	return temp.Where("company_type = ?", 4)
}

//
func (c *CompanyRepository) FindCompany(per, page int, filters map[string]interface{},
	selectKeys []string, orders []string,
	isCount bool) (companies []models.UserCompany, total int64, err error) {
	sqlCon := global.DataEngine.Model(&models.UserCompany{}).Scopes(c.DefaultScope)
	if len(selectKeys) > 0 {
		selectKeys = []string{"user_companies.*"}
	}
	if isCount {
		countCon := global.DataEngine.Model(&models.UserCompany{}).Scopes(c.DefaultScope)
		if total, err = c.Count(countCon, filters); err != nil {
			return
		}
	}
	sqlCon = c.crud.Where(sqlCon, filters, selectKeys, c.Paginate(per, page), c.OrderBy(orders))
	err = sqlCon.Find(&companies).Error
	return
}

//

func NewCompanyRepository() ICompanyRepository {
	return &CompanyRepository{}
}
