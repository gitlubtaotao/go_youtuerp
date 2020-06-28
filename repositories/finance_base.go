package repositories

import (
	"database/sql"
	"github.com/kataras/golog"
	"reflect"
	"time"
	"youtuerp/database"
	"youtuerp/models"
)

type IFinanceBase interface {
	//根据系统设置的按月汇率或者按实时汇率查询对应的汇率信息
	GetAllFeeRate(rateSetting string, attr map[string]interface{}) (feeRate []models.FinanceRate, err error)
	Update(id uint, record interface{}) error
	Create(record interface{}) (interface{}, error)
	Delete(id uint, model interface{}) error
	FindRate(per, page uint, filter map[string]interface{},
		selectKeys []string, orders []string) ([]models.FinanceRate, uint, error)
	FindFeeType(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) ([]models.FinanceFeeType, uint, error)
}
type FinanceBase struct {
	BaseRepository
}

func (f FinanceBase) GetAllFeeRate(rateSetting string, attr map[string]interface{}) (feeRate []models.FinanceRate, err error) {
	sqlCon := database.GetDBCon().Model(&models.FinanceRate{})
	var rows *sql.Rows
	if rateSetting == models.SettingFeeRateNow {
		attr["start_month"] = 0
		attr["end_month"] = 0
	} else if rateSetting == models.SettingFeeRateMonth {
		attr["year"] = time.Now().Year()
		attr["start_month"] = int(time.Now().Month())
		attr["end_month"] = int(time.Now().Month()) + 1
	}
	golog.Infof("rate setting is %v current month %v", rateSetting, int(time.Now().Month()))
	rows, err = sqlCon.Where(attr).Order("id desc").Group("finance_currency_id").Select("finance_currency_id,rate").Rows()
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
		return database.GetDBCon().Model(&models.FinanceRate{ID: id}).Update(record).Error
	} else if name == "FinanceFeeType" {
		return database.GetDBCon().Model(&models.FinanceFeeType{ID: id}).Update(record).Error
	}
	return nil
}
func (f FinanceBase) Delete(id uint, model interface{}) error {
	return database.GetDBCon().Delete(model, "id = ?", id).Error
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
