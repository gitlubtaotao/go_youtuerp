package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type ICrmCompany interface {
	Delete(id uint) error
	Update(id uint, company models.CrmCompany) (models.CrmCompany, error)
	UpdateByMap(id uint, attr map[string]interface{}) error
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) ([]models.CrmCompany, uint, error)
	First(id uint, isRole bool,isUser bool) (models.CrmCompany, error)
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
	err := database.GetDBCon().Model(&company).Update(attr).Error
	if err != nil {
		return err
	}
	if companyType, ok = attr["company_type"]; !ok {
		return nil
	}
	if companyType == 3 {
		return database.GetDBCon().Model(&company).Association("Roles").Append(models.Role{
			UserId: userSalesmanId.(uint),
			Name:   models.RoleNameSale}).Error
	}
	return nil
}

func (c CrmCompany) Delete(id uint) error {
	return c.crud.Delete(&models.CrmCompany{}, id)
}
func (c CrmCompany) Update(id uint, company models.CrmCompany) (models.CrmCompany, error) {
	var record models.CrmCompany
	sqlCon := database.GetDBCon().First(&record, "id = ? ", id)
	sqlCon.Association("Roles").Replace(company.Roles)
	err := sqlCon.Set("gorm:association_autocreate", false).Update(company).Error
	return record, err
}

func (c CrmCompany) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (companies []models.CrmCompany, total uint, err error) {
	sqlCon := database.GetDBCon().Model(&models.CrmCompany{})
	if isTotal {
		if total, err = c.Count(sqlCon, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"user_companies.*"}
	}
	sqlCon = sqlCon.Preload("Roles")
	sqlCon = c.crud.Where(sqlCon, filter, selectKeys, c.Paginate(per, page), c.OrderBy(orders))
	err = sqlCon.Find(&companies).Error
	return
}

func (c CrmCompany) Create(company models.CrmCompany) (models.CrmCompany, error) {
	err := database.GetDBCon().Model(&models.CrmCompany{}).Create(&company).Error
	return company, err
}

func (c CrmCompany) First(id uint, isRole bool,isUser bool) (company models.CrmCompany, err error) {
	sqlConn := database.GetDBCon().Model(&models.CrmCompany{})
	if isRole {
		sqlConn = sqlConn.Preload("Roles")
	}
	if isUser{
		sqlConn = sqlConn.Preload("CrmUsers")
	}
	err = sqlConn.First(&company, "user_companies.id = ?", id).Error
	return
}
func NewCrmCompany() ICrmCompany {
	return &CrmCompany{}
}
