package repositories

import (
	"database/sql"
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type ISelectRepository interface {
	FindTable(tableName string, name string, scope map[string]interface{}, selectKeys []string) (selectResult []models.SelectResult, err error)
	FirstRecord(tableName string, filter map[string]interface{}, selectKeys []string) (selectResult []models.SelectResult, err error)
}

type SelectRepository struct {
	BaseRepository
}

func (s *SelectRepository) FirstRecord(tableName string, filter map[string]interface{}, selectKeys []string) (selectResult []models.SelectResult, err error) {
	sqlCon := database.GetDBCon().Table(tableName)
	var rows *sql.Rows
	if len(selectKeys) == 0 {
		selectKeys = []string{"id", "name"}
	}
	sqlCon = sqlCon.Where(filter).Where("deleted_at is NULL").Select(selectKeys)
	if rows, err = sqlCon.Rows(); err != nil {
		return
	}
	for rows.Next() {
		var temp models.SelectResult
		_ = sqlCon.ScanRows(rows, &temp)
		selectResult = append(selectResult, temp)
	}
	return selectResult, nil
}

func (s *SelectRepository) FindTable(tableName string, name string, scope map[string]interface{},
	selectKeys []string) (selectResult []models.SelectResult, err error) {
	sqlCon := database.GetDBCon().Table(tableName)
	if len(selectKeys) == 0 {
		selectKeys = []string{"id", "name"}
	}
	if name != "" {
		sqlCon = s.defaultScope(sqlCon, tableName, name)
	}
	sqlCon = sqlCon.Scopes(s.Ransack(scope)).Where("deleted_at is NULL").Select(selectKeys)
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

func (s SelectRepository) defaultScope(db *gorm.DB, taleName string, name string) *gorm.DB {
	if taleName == "user_companies" {
		db = db.Scopes(s.companyWhere(name))
	} else {
		db = db.Scopes(s.defaultWhere(name))
	}
	return db
}

func (s SelectRepository) defaultWhere(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ? ", "%"+name)
	}
}

func (s SelectRepository) companyWhere(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name_nick like ? ", "%"+name)
	}
}

//
func NewSelectRepository() ISelectRepository {
	return &SelectRepository{}
}
