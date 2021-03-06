package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type IBaseCode interface {
	Update(id uint, code models.BaseDataCode) error
	Delete(id uint) error
	Create(code models.BaseDataCode) (models.BaseDataCode, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) (codes []models.BaseDataCode, total uint, err error)
	FindAllLevel() (levels []models.BaseDataLevel, err error)
}
type BaseCode struct {
	BaseRepository
}

func (b BaseCode) Update(id uint, code models.BaseDataCode) error {
	return database.GetDBCon().Model(&models.BaseDataCode{ID: id}).Update(code).Error
}

func (b BaseCode) Delete(id uint) error {
	return b.crud.Delete(&models.BaseDataCode{}, id)
}
func (b BaseCode) Create(code models.BaseDataCode) (models.BaseDataCode, error) {
	err := database.GetDBCon().Create(&code).Error
	return code, err
}

func (b BaseCode) FindAllLevel() (levels []models.BaseDataLevel, err error) {
	sqlConn := database.GetDBCon().Model(&models.BaseDataLevel{})
	err = sqlConn.Find(&levels).Error
	return levels, err
}

func (b BaseCode) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (codes []models.BaseDataCode, total uint, err error) {
	sqlConn := database.GetDBCon().Model(&models.BaseDataCode{})
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
