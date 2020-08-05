package dao

import (
	"youtuerp/internal/models"
)

type IBaseCode interface {
	Update(id uint, code models.BaseDataCode) error
	Delete(id uint) error
	Create(code models.BaseDataCode) (models.BaseDataCode, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) (codes []models.BaseDataCode, total int64, err error)
	FindAllLevel() (levels []models.BaseDataLevel, err error)
}
type BaseCode struct {
	BaseRepository
}

func (b BaseCode) Update(id uint, code models.BaseDataCode) error {
	return global.DataEngine.Model(&models.BaseDataCode{ID: id}).Updates(code).Error
}

func (b BaseCode) Delete(id uint) error {
	return b.crud.Delete(&models.BaseDataCode{}, id)
}
func (b BaseCode) Create(code models.BaseDataCode) (models.BaseDataCode, error) {
	err := global.DataEngine.Create(&code).Error
	return code, err
}

func (b BaseCode) FindAllLevel() (levels []models.BaseDataLevel, err error) {
	sqlConn := global.DataEngine.Model(&models.BaseDataLevel{})
	err = sqlConn.Find(&levels).Error
	return levels, err
}

func (b BaseCode) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (codes []models.BaseDataCode, total int64, err error) {
	sqlConn := global.DataEngine.Model(&models.BaseDataCode{})
	if isTotal {
		if total, err = b.Count(sqlConn, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"base_data_codes.*"}
	}
	sqlConn = b.crud.Where(sqlConn, filter, selectKeys, b.Paginate(per, page), b.OrderBy(orders))
	err = sqlConn.Find(&codes).Error
	return codes, total, err
}

func NewBaseCode() IBaseCode {
	return &BaseCode{}
}
