package dao

import (
	"github.com/kataras/golog"
	"youtuerp/database"
	"youtuerp/internal/models"
)

type ICrmCompany interface {
	Delete(id uint) error
	Update(id uint, company models.CrmCompany) (models.CrmCompany, error)
	UpdateByMap(id uint, attr map[string]interface{}) error
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) ([]models.CrmCompany, int64, error)
	First(id uint, preload ...string) (models.CrmCompany, error)
	Create(company models.CrmCompany) (models.CrmCompany, error)
}

type CrmCompany struct {
	BaseRepository
}

func (c CrmCompany) UpdateByMap(id uint, attr map[string]interface{}) error {
	var company = models.CrmCompany{ID: id}
	var (
		userSalesmanId interface{}
		ok             bool
		companyType    interface{}
	)
	if userSalesmanId, ok = attr["user_salesman_id"]; ok {
		delete(attr, "user_salesman_id")
	}
	err := database.GetDBCon().Model(&company).Updates(attr).Error
	if err != nil {
		return err
	}
	if companyType, ok = attr["company_type"]; !ok {
		return nil
	}
	if companyType == 3 {
		_ = database.GetDBCon().Model(&company).Association("Roles").Append(models.Role{
			UserId: userSalesmanId.(uint),
			Name:   models.RoleNameSale})
	}
	return nil
}

func (c CrmCompany) Delete(id uint) error {
	return c.crud.Delete(&models.CrmCompany{}, id)
}
func (c CrmCompany) Update(id uint, company models.CrmCompany) (models.CrmCompany, error) {
	var record models.CrmCompany
	//sqlCon := database.GetDBCon().First(&record, "id = ? ", id)
	golog.Infof("roles is %v", company.Roles)
	if err := database.GetDBCon().Model(&models.CrmCompany{ID: id}).Association("Roles").Replace(company.Roles); err != nil {
		return record, err
	}
	err := database.GetDBCon().Model(&models.CrmCompany{ID: id}).Set("gorm:association_autocreate", false).Updates(company).Error
	return company, err
}

func (c CrmCompany) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (companies []models.CrmCompany, total int64, err error) {
	sqlCon := database.GetDBCon().Model(&models.CrmCompany{})
	if isTotal {
		if total, err = c.Count(database.GetDBCon().Model(&models.CrmCompany{}), filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"user_companies.*"}
	}
	sqlCon = sqlCon.Preload("Roles")
	err = sqlCon.Scopes(c.CustomerWhere(filter, selectKeys, c.Paginate(per, page), c.OrderBy(orders))).Find(&companies).Error
	return
}

func (c CrmCompany) Create(company models.CrmCompany) (models.CrmCompany, error) {
	err := database.GetDBCon().Model(&models.CrmCompany{}).Create(&company).Error
	return company, err
}

func (c CrmCompany) First(id uint, preload ...string) (company models.CrmCompany, err error) {
	sqlConn := database.GetDBCon().Model(&models.CrmCompany{})
	for _, column := range preload {
		sqlConn = sqlConn.Preload(column)
	}
	err = sqlConn.First(&company, "user_companies.id = ?", id).Error
	return
}
func NewCrmCompany() ICrmCompany {
	return &CrmCompany{}
}
