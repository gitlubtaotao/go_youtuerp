package repositories

import (
	"database/sql"
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
	return n.crud.Delete(&models.NumberSetting{}, id)
}

func (n NumberSettingRepository) Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) (numberSettings []models.ResultNumberSetting,
	total uint, err error) {
	var rows *sql.Rows
	sqlCon := database.GetDBCon().Model(&models.NumberSetting{}).Unscoped()
	sqlCon = sqlCon.Joins("inner join user_companies on user_companies.id = number_settings.user_company_id and user_companies.company_type = 4")
	if isCount {
		if total, err = n.Count(sqlCon, filter);err != nil{
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"number_settings.*", "user_companies.name_nick as user_companies_name_nick",}
	}
	sqlCon = n.crud.Where(sqlCon, filter, selectKeys, n.Paginate(per, page), n.OrderBy(order))
	if rows, err = sqlCon.Rows();err != nil{
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
	err :=  n.crud.Create(&numberSetting)
	return numberSetting, err
}

func NewNumberSetting() INumberSettingRepository {
	return &NumberSettingRepository{}
}
