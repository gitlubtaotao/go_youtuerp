package repositories

import (
	"fmt"
	"youtuerp/database"
	"youtuerp/models"
)

type ISelectRepository interface {
	FindModel(model interface{}, scope map[string]interface{}, selectKey []string) (selectResult []models.SelectResult, err error)
	FindTable(tableName string, name string, scope map[string]interface{}, selectKeys []string) (selectResult []models.SelectResult, err error)
}

type SelectRepository struct {
	BaseRepository
}

func (s *SelectRepository) FindTable(tableName string, name string, scope map[string]interface{},
	selectKeys []string) (selectResult []models.SelectResult, err error) {
	fmt.Printf("%v,%v,%v", tableName, scope, selectKeys)
	sqlCon := database.GetDBCon().Table(tableName)
	if len(selectKeys) == 0 {
		selectKeys = []string{"id", "name"}
	}
	if name != "" {
		sqlCon = sqlCon.Where("name link ? ", "%"+name+"%")
	}
	sqlCon = sqlCon.Where(scope).Scopes(s.Paginate(20, 1)).Select(selectKeys)
	rows, err := sqlCon.Rows()
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

//
func NewSelectRepository() ISelectRepository {
	return &SelectRepository{}
}
