package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type IBaseCarrier interface {
	Update(id uint, code models.BaseDataCarrier) error
	Delete(id uint) error
	Create(code models.BaseDataCarrier) (models.BaseDataCarrier, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) (codes []models.BaseDataCarrier, total uint, err error)
}

type BaseCarrier struct {
	BaseRepository
}

func (b BaseCarrier) Update(id uint, carrier models.BaseDataCarrier) error {
	return database.GetDBCon().Model(&models.BaseDataCarrier{ID: id}).Update(carrier).Error
}

func (b BaseCarrier) Delete(id uint) error {
	return b.crud.Delete(&models.BaseDataCode{}, id)
}
func (b BaseCarrier) Create(code models.BaseDataCarrier) (models.BaseDataCarrier, error) {
	err := database.GetDBCon().Create(&code).Error
	return code, err
}



func (b BaseCarrier) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (codes []models.BaseDataCarrier, total uint, err error) {
	sqlConn := database.GetDBCon().Model(&models.BaseDataCarrier{})
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
