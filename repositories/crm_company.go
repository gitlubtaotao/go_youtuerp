package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type ICrmCompany interface {
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) ([]models.CrmCompany, uint, error)
	Create(company models.CrmCompany) (models.CrmCompany, error)
}

type CrmCompany struct {
	BaseRepository
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
	err =  sqlCon.Find(&companies).Error
	return
}

func (c CrmCompany) Create(company models.CrmCompany) (models.CrmCompany, error) {
	err := database.GetDBCon().Model(&models.CrmCompany{}).Create(&company).Error
	return company, err
}
func NewCrmCompany() ICrmCompany {
	return &CrmCompany{}
}
