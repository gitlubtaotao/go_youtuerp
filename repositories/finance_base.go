package repositories

import (
	"database/sql"
	"reflect"
	"youtuerp/database"
	"youtuerp/models"
)

type IFinanceBase interface {
	//根据系统设置的按月汇率或者按实时汇率查询对应的汇率信息
	GetAllFeeRate(otherFilter ...map[string]interface{}) (feeRate []models.FinanceRate, err error)
	Update(id uint, record interface{}) error
	Create(record interface{}) (interface{}, error)
	Delete(id uint, model interface{}) error
	FindRate(per, page int, filter map[string]interface{},
		selectKeys []string, orders []string) ([]models.FinanceRate, int64, error)
	FindFeeType(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) ([]models.FinanceFeeType, int64, error)
}
type FinanceBase struct {
	BaseRepository
}

func (f FinanceBase) GetAllFeeRate(filterOther ...map[string]interface{}) (feeRate []models.FinanceRate, err error) {
	sqlCon := database.GetDBCon().Model(&models.FinanceRate{})
	var rows *sql.Rows
	if len(filterOther) >= 1 {
		for _, filter := range filterOther {
			sqlCon = sqlCon.Where(filter)
		}
	}
	rows, err = sqlCon.Order("id desc").Group("finance_currency_id").Select("finance_currency_id,rate").Rows()
	if err != nil {
		return feeRate, err
	}
	for rows.Next() {
		var data models.FinanceRate
		_ = sqlCon.ScanRows(rows, &data)
		feeRate = append(feeRate, data)
	}
	return feeRate, err
}

func (f FinanceBase) Update(id uint, record interface{}) error {
	name := reflect.TypeOf(record).Name()
	if name == "FinanceRate" {
		return database.GetDBCon().Model(&models.FinanceRate{ID: id}).Updates(record).Error
	} else if name == "FinanceFeeType" {
		return database.GetDBCon().Model(&models.FinanceFeeType{ID: id}).Updates(record).Error
	}
	return nil
}
func (f FinanceBase) Delete(id uint, model interface{}) error {
	return database.GetDBCon().Delete(model, "id = ?", id).Error
}

func (f FinanceBase) FindRate(per, page int, filter map[string]interface{},
	selectKeys []string, orders []string) (records []models.FinanceRate, total int64, err error) {
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
func (f FinanceBase) FindFeeType(per, page int, filter map[string]interface{},
	selectKeys []string, orders []string) (records []models.FinanceFeeType, total int64, err error) {
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
	rt := reflect.TypeOf(record)
	if rt.Name() == "FinanceRate" {
		rate := record.(models.FinanceRate)
		err = database.GetDBCon().Create(&rate).Error
		return rate, err
	}
	if rt.Name() == "FinanceFeeType" {
		feeType := record.(models.FinanceFeeType)
		err = database.GetDBCon().Create(&feeType).Error
		return feeType, err
	}
	return record, err
}

func NewFinanceBase() IFinanceBase {
	return &FinanceBase{}
}
