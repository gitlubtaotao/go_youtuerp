package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type ISelectRepository interface {
	FindModel(model interface{}, scope map[string]interface{}, selectKey []string) (selectResult []models.SelectResult, err error)
}

type SelectRepository struct {
	BaseRepository
}

//
func NewSelectRepository() ISelectRepository {
	return &SelectRepository{}
}

func (s *SelectRepository) FindModel(model interface{}, scope map[string]interface{}, selectKey []string) (selectResult []models.SelectResult, err error) {
	temp := database.GetDBCon().Model(model)
	sqlCon := database.GetDBCon()
	selectKey = append(selectKey, "id")
	if len(scope) > 0 {
		temp = temp.Scopes(s.Ransack(scope))
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
