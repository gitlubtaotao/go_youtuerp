package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type ISelectRepository interface {
	Find(tableName string, scope map[string]interface{}, selectKey []string) (selectResult []models.SelectResult, err error)
}

type SelectRepository struct {
	BaseRepository
}

//
func NewSelectRepository() ISelectRepository {
	return &SelectRepository{}
}

func (s *SelectRepository) Find(tableName string, scope map[string]interface{}, selectKey []string) (selectResult []models.SelectResult, err error) {
	temp := database.GetDBCon().Table(tableName)
	sqlCon := database.GetDBCon()
	if len(selectKey) > 0 {
		selectKey = append(selectKey, "id")
		temp = temp.Select(selectKey)
	} else {
		temp = temp.Select("*")
	}
	if len(scope) > 0 {
		temp = temp.Where(scope)
	}
	temp = temp.Scopes(s.Paginate(20, 1))
	rows, err := temp.Rows()
	if err != nil {
		return []models.SelectResult{}, nil
	}
	for rows.Next() {
		var temp models.SelectResult
		_ = sqlCon.ScanRows(rows, &temp)
		selectResult = append(selectResult, temp)
	}
	return selectResult, err
}
