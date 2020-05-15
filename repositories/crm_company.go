package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type ICrmCompany interface {
	Update(id uint, company models.CrmCompany)(models.CrmCompany, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) ([]models.CrmCompany, uint, error)
	First(id uint, isRole bool) (models.CrmCompany, error)
	
	Create(company models.CrmCompany) (models.CrmCompany, error)
}

type CrmCompany struct {
	BaseRepository
}

func (c CrmCompany) Update(id uint,company models.CrmCompany)(models.CrmCompany, error) {
	var record models.CrmCompany
	sqlCon := database.GetDBCon().First(&record, "id = ? ", id)
	sqlCon.Association("Roles").Replace(company.Roles)
	err := sqlCon.Set("gorm:association_autocreate", false).Update(company).Error
	return record,err
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

func (c CrmCompany) First(id uint, isRole bool) (company models.CrmCompany, err error) {
	sqlConn := database.GetDBCon().Model(&models.CrmCompany{})
	if isRole {
		sqlConn = sqlConn.Preload("Roles")
	}
	err = sqlConn.First(&company, "user_companies.id = ?", id).Error
	return
}
func NewCrmCompany() ICrmCompany {
	return &CrmCompany{}
}
