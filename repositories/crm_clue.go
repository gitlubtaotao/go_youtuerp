package repositories

import (
	"database/sql"
	"youtuerp/database"
	"youtuerp/models"
)

type ICrmClue interface {
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isCount bool) ([]models.CrmClue, uint, error)
	Create(clue models.CrmClue) (models.CrmClue, error)
}
type CrmClue struct {
	BaseRepository
}

func (c CrmClue) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string, isCount bool) (clues []models.CrmClue, total uint, err error) {
	var rows *sql.Rows
	sqlCon := database.GetDBCon().Model(&models.CrmClue{})
	sqlCon = sqlCon.Joins("LEFT JOIN users ON users.id = crm_clues.create_id")
	if isCount {
		if total, err = c.Count(sqlCon, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"crm_clues.*", "users.name as create_name",}
	}
	sqlCon = c.crud.Where(sqlCon, filter, selectKeys, c.Paginate(per, page), c.OrderBy(orders))
	 rows, err = sqlCon.Rows()
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
	err := database.GetDBCon().Set("gorm:association_autocreate", false).Create(&clue).Error
	if err != nil {
		return models.CrmClue{}, err
	}
	return clue, err
}

func NewCrmClue() ICrmClue {
	return &CrmClue{}
}
