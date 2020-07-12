package repositories

import (
	"database/sql"
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
)

type ICompanyRepository interface {
	/*查询公司信息
	per: 每次分页的记录数
	page: 当前页数
	attr: 查询公司信息
	*/
	FindCompany(per, page int, attr map[string]interface{}, selectKeys []string, order []string, isCount bool) (companies []*models.UserCompany,
		total int64, err error)
	FirstCompany(id uint) (*models.UserCompany, error)
	FirstCompanyByRelated(id uint, related ...string) (models.UserCompany, error)
	CreateCompany(company models.UserCompany) (models.UserCompany, error)
	UpdateCompany(company *models.UserCompany, readData models.UserCompany) error
	DeleteCompany(id uint) error
}

type CompanyRepository struct {
	BaseRepository
}

func (c CompanyRepository) FirstCompanyByRelated(id uint, related ...string) (models.UserCompany, error) {
	sqlCon := database.GetDBCon()
	var company models.UserCompany
	for _, re := range related {
		sqlCon = sqlCon.Preload(re)
	}
	err := sqlCon.Find(&company, "id = ?", id).Error
	return company, err
}

func (c CompanyRepository) DeleteCompany(id uint) error {
	var readData models.UserCompany
	return database.GetDBCon().Find(&readData).Delete(&readData).Error
}

func (c CompanyRepository) UpdateCompany(company *models.UserCompany, readData models.UserCompany) error {
	return database.GetDBCon().Model(&company).Updates(tools.StructToChange(readData)).Error
}

func (c CompanyRepository) FirstCompany(id uint) (company *models.UserCompany, err error) {
	var readData models.UserCompany
	err = database.GetDBCon().First(&readData, id).Error
	company = &readData
	return
}

func (c CompanyRepository) CreateCompany(company models.UserCompany) (models.UserCompany, error) {
	err := database.GetDBCon().Create(&company).Error
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
	isCount bool) (companies []*models.UserCompany, total int64, err error) {
	var rows *sql.Rows
	sqlCon := database.GetDBCon()
	temp := sqlCon.Scopes(c.DefaultScope)
	if len(filters) > 0 {
		temp = temp.Scopes(c.Ransack(filters))
	}
	//limit
	if isCount {
		temp.Model(&models.UserCompany{}).Count(&total)
	}
	if page == 0 && per > 0 {
		temp = temp.Limit(per)
	} else if page > 0 && per > 0 {
		temp = temp.Limit(per).Offset((page - 1) * per)
	}
	if len(orders) == 0 {
		temp = temp.Order("id desc")
	} else {
		for _, order := range orders {
			temp = temp.Order(order)
		}
	}
	if len(selectKeys) > 0 {
		temp = temp.Select(selectKeys)
	}
	rows, err = temp.Model(&models.UserCompany{}).Rows()
	if err != nil {
		return
	}
	
	for rows.Next() {
		var userCompany models.UserCompany
		_ = sqlCon.ScanRows(rows, &userCompany)
		companies = append(companies, &userCompany)
	}
	return
}

//

func NewCompanyRepository() ICompanyRepository {
	return &CompanyRepository{}
}
