package repositories

import (
	"database/sql"
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
)

type ICrmClue interface {
	Delete(id uint) error
	Update(id uint, clue models.CrmClue) error
	First(id uint, isTacks bool) (models.CrmClue, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string, isCount bool) ([]models.CrmClue, int64, error)
	Create(clue models.CrmClue) (models.CrmClue, error)
}
type CrmClue struct {
	BaseRepository
}

func (c CrmClue) Delete(id uint) error {
	return c.crud.Delete(&models.NumberSetting{}, id)
}

func (c CrmClue) Update(id uint, clue models.CrmClue) error {
	var record models.CrmClue
	err := database.GetDBCon().First(&record, "id = ? ", id).Updates(tools.StructToChange(clue)).Error
	return err
}

func (c CrmClue) First(id uint, isTacks bool) (models.CrmClue, error) {
	var data models.CrmClue
	sqlConn := database.GetDBCon().Model(&models.CrmClue{})
	if isTacks {
		sqlConn = sqlConn.Preload("UserCreate")
	}
	err := sqlConn.First(&data, "crm_clues.id = ?", id).Error
	return data, err
}
func (c CrmClue) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string, isCount bool) (clues []models.CrmClue, total int64, err error) {
	var rows *sql.Rows
	sqlCon := database.GetDBCon().Model(&models.CrmClue{}).Scopes(c.defaultScope())
	if isCount {
		if total, err = c.Count(database.GetDBCon().Model(&models.CrmClue{}).Scopes(c.defaultScope()), filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = c.defaultSelectKeys()
	}
	rows, err = sqlCon.Scopes(c.CustomerWhere(filter, selectKeys, c.Paginate(per, page), c.OrderBy(orders))).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.CrmClue
		err = sqlCon.ScanRows(rows, &data)
		clues = append(clues, data)
	}
	return
}

func (c CrmClue) Create(clue models.CrmClue) (models.CrmClue, error) {
	sqlCon := database.GetDBCon().Set("gorm:association_autoupdate", false)
	sqlCon = sqlCon.Set("gorm:association_autocreate", false)
	err := sqlCon.Create(&clue).Error
	if err != nil {
		return models.CrmClue{}, err
	}
	return clue, err
}

func (c CrmClue) defaultScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("LEFT JOIN users ON users.id = crm_clues.create_id")
	}
}
func (c CrmClue) defaultSelectKeys() []string {
	return []string{"crm_clues.*", "users.name as create_name",}
}

func NewCrmClue() ICrmClue {
	return &CrmClue{}
}
