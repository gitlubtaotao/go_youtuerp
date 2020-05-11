package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type INumberSettingRepository interface {
	Create(numberSetting models.NumberSetting) (models.NumberSetting, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) (numberSettings []models.ResultNumberSetting,
		total uint, err error)
	Delete(id uint) error
}
type NumberSettingRepository struct {
	BaseRepository
}


func (n NumberSettingRepository) Delete(id uint) error {
	return  database.GetDBCon().Delete(&models.Setting{},"id = ?",id).Error
}

func (n NumberSettingRepository) Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) (numberSettings []models.ResultNumberSetting,
	total uint, err error) {
	sqlCon := database.GetDBCon().Model(&models.NumberSetting{}).Unscoped()
	sqlCon = sqlCon.Joins("inner join user_companies on user_companies.id = number_settings.user_company_id and user_companies.company_type = 4")
	
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(n.Ransack(filter))
	}
	if isCount {
		err = sqlCon.Count(&total).Error
		if err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"number_settings.*", "user_companies.name_nick as user_companies_name_nick",}
	}
	rows, err := sqlCon.Scopes(n.Paginate(per, page), n.OrderBy(order)).Select(selectKeys).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.ResultNumberSetting
		_ = sqlCon.ScanRows(rows, &data)
		numberSettings = append(numberSettings, data)
	}
	return
}

func (n NumberSettingRepository) Create(numberSetting models.NumberSetting) (models.NumberSetting, error) {
	err := database.GetDBCon().Create(&numberSetting).Error
	return numberSetting, err
}

func NewNumberSetting() INumberSettingRepository {
	return &NumberSettingRepository{}
}
