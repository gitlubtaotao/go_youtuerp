package repositories

import (
	"database/sql"
	"github.com/kataras/golog"
	"reflect"
	"youtuerp/database"
	"youtuerp/models"
)

type IFinanceBase interface {
	Create(record interface{}) (interface{}, error)
	FindRate(per, page uint, filter map[string]interface{},
		selectKeys []string, orders []string) ([]models.FinanceRate, uint, error)
	FindFeeType(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) ([]models.FinanceFeeType, uint, error)
}
type FinanceBase struct {
	BaseRepository
}

func (f FinanceBase) FindRate(per, page uint, filter map[string]interface{},
	selectKeys []string, orders []string) (records []models.FinanceRate, total uint, err error) {
	var rows *sql.Rows
	sqlConn := database.GetDBCon().Model(&models.FinanceRate{})
	if total, err = f.Count(sqlConn, filter); err != nil {
		return
	}
	sqlConn = f.crud.Where(sqlConn, filter, selectKeys, f.Paginate(per, page), f.OrderBy(orders))
	if rows, err = sqlConn.Rows(); err != nil {
		return
	}
	for rows.Next() {
		var data models.FinanceRate
		_ = sqlConn.ScanRows(rows, &data)
		records = append(records, data)
	}
	return records, total, err
}
func (f FinanceBase) FindFeeType(per, page uint, filter map[string]interface{},
	selectKeys []string, orders []string) (records []models.FinanceFeeType, total uint, err error) {
	sqlConn := database.GetDBCon().Model(&models.FinanceFeeType{})
	if total, err = f.Count(sqlConn, filter); err != nil {
		return
	}
	sqlConn = f.crud.Where(sqlConn, filter, selectKeys, f.Paginate(per, page), f.OrderBy(orders))
	err = sqlConn.Find(&records).Error
	return records, total, err
}

func (f FinanceBase) Create(record interface{}) (interface{}, error) {
	var err error
	golog.Infof("record is %v", record)
	rt := reflect.TypeOf(record)
	golog.Infof("record type is %v", rt.Name())
	if rt.Name() == "FinanceRate" {
		rate := record.(models.FinanceRate)
		err = database.GetDBCon().Create(&rate).Error
	}
	return record, err
}

func NewFinanceBase() IFinanceBase {
	return &FinanceBase{}
}
