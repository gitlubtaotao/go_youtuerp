package services

import (
	"errors"
	"reflect"
	"strconv"
	"time"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
)

type IFinanceBase interface {
	Update(id uint, record interface{}, language string) error
	Delete(id uint, model interface{}) error
	FindRate(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.FinanceRate, int64, error)
	FindFeeType(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.FinanceFeeType, int64, error)
	Create(record interface{}, language string) (interface{}, error)
	//获取费用种类的redis数据
	FindFeeTypeRedis() []map[string]string
	//根据系统设置的按月汇率或者按实时汇率查询对应的汇率信息
	GetAllFeeRate(companyId uint, otherFilter ...map[string]interface{}) ([]models.FinanceRate, error)
}

type FinanceBase struct {
	BaseService
	repo repositories.IFinanceBase
}

func (f FinanceBase) GetAllFeeRate(companyId uint, filterOther ...map[string]interface{}) ([]models.FinanceRate, error) {
	rateSetting := redis.SystemRateSetting()
	//按照实时的汇率进行计算
	if rateSetting == models.SettingFeeRateNow {
		return f.repo.GetAllFeeRate(map[string]interface{}{"start_month": 0, "end_month": 0, "company_id": companyId})
	}
	if len(filterOther) == 0 {
		filterOther = append(filterOther, map[string]interface{}{
			"year":        time.Now().Year(),
			"start_month": int(time.Now().Month()),
			"end_month":   int(time.Now().Month()) + 1,
		})
	}
	filterOther = append(filterOther, map[string]interface{}{"company_id": companyId})
	return f.repo.GetAllFeeRate(filterOther...)
}

func (f FinanceBase) FindFeeTypeRedis() []map[string]string {
	red := redis.NewRedis()
	tableName := models.FinanceFeeType{}.TableName()
	data := make([]map[string]string, 0)
	data = red.HCollectOptions(tableName)
	if len(data) > 0 {
		return data
	}
	records, _, _ := f.repo.FindFeeType(0, 0, map[string]interface{}{}, []string{}, []string{})
	for _, k := range records {
		go f.saveRedisData(k)
		temp := map[string]string{"id": strconv.Itoa(int(k.ID)),
			"name": k.Name, "name_cn": k.NameCn, "name_en": k.NameEn, "finance_currency_id": strconv.Itoa(int(k.FinanceCurrencyId))}
		data = append(data, temp)
	}
	return data
}

func (f FinanceBase) Update(id uint, record interface{}, language string) error {
	valid := NewValidatorService(record)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	if err := f.repo.Update(id, record); err != nil {
		return err
	}
	go f.saveRedisData(record)
	return nil
}

func (f FinanceBase) Delete(id uint, model interface{}) error {
	err := f.repo.Delete(id, model)
	if err != nil {
		return err
	}
	go f.deleteRedisData(id, model)
	return nil
}

func (f FinanceBase) FindFeeType(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.FinanceFeeType, int64, error) {
	return f.repo.FindFeeType(per, page, filter, selectKeys, orders)
}

func (f FinanceBase) FindRate(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.FinanceRate, int64, error) {
	return f.repo.FindRate(per, page, filter, selectKeys, orders)
}

func (f FinanceBase) Create(record interface{}, language string) (data interface{}, err error) {
	valid := NewValidatorService(record)
	if message := valid.ResultError(language); message != "" {
		return record, errors.New(message)
	}
	data, err = f.repo.Create(record)
	if err != nil {
		return
	}
	go f.saveRedisData(data)
	return
}

//将币种和费用类型进行redis缓存
func (f FinanceBase) saveRedisData(data interface{}) {
	ty := reflect.TypeOf(data).Name()
	if ty != "FinanceFeeType" {
		return
	}
	record := data.(models.FinanceFeeType)
	tableName := models.FinanceFeeType{}.TableName()
	
	temp := map[string]interface{}{
		"id":                  record.ID,
		"name":                record.Name,
		"name_en":             record.NameEn,
		"name_cn":             record.NameCn,
		"finance_currency_id": record.FinanceCurrencyId,
	}
	redis.HSetValue(tableName, record.ID, temp)
}

//删除币种和费用对应的redis缓存
func (f FinanceBase) deleteRedisData(id uint, data interface{}) {
	ty := reflect.TypeOf(data).Kind().String()
	if ty != "FinanceFeeType" {
		return
	}
	tableName := models.FinanceFeeType{}.TableName()
	red := redis.NewRedis()
	red.SRemove(tableName, id)
}

func NewFinanceBase() IFinanceBase {
	return &FinanceBase{repo: repositories.NewFinanceBase()}
}
