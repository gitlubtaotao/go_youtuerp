package dao

import (
	"youtuerp/global"
	"youtuerp/internal/models"
)

type IBaseCarrier interface {
	Update(id uint, code models.BaseDataCarrier) error
	Delete(id uint) error
	Create(code models.BaseDataCarrier) (models.BaseDataCarrier, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) (codes []models.BaseDataCarrier, total int64, err error)
}

type BaseCarrier struct {
	BaseRepository
}

func (b BaseCarrier) Update(id uint, carrier models.BaseDataCarrier) error {
	return global.DataEngine.Model(&models.BaseDataCarrier{ID: id}).Updates(carrier).Error
}

func (b BaseCarrier) Delete(id uint) error {
	return b.crud.Delete(&models.BaseDataCarrier{}, id)
}
func (b BaseCarrier) Create(code models.BaseDataCarrier) (models.BaseDataCarrier, error) {
	err := global.DataEngine.Create(&code).Error
	return code, err
}

func (b BaseCarrier) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (codes []models.BaseDataCarrier, total int64, err error) {
	sqlConn := global.DataEngine.Model(&models.BaseDataCarrier{})
	if isTotal {
		if total, err = b.Count(sqlConn, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"base_data_carriers.*"}
	}
	sqlConn = b.crud.Where(sqlConn, filter, selectKeys, b.Paginate(per, page), b.OrderBy(orders))
	err = sqlConn.Find(&codes).Error
	return codes, total, err
}
func NewBaseCarrier() IBaseCarrier {
	return &BaseCarrier{}
}
