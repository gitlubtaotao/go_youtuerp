package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type IBaseWarehouse interface {
	Update(id uint, code models.BaseWarehouse) error
	Delete(id uint) error
	Create(code models.BaseWarehouse) (models.BaseWarehouse, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) (codes []models.BaseWarehouse, total uint, err error)
}
type BaseWarehouse struct {
	BaseRepository
}

func (b BaseWarehouse) Update(id uint, code models.BaseWarehouse) error {
	return database.GetDBCon().Model(&models.BaseWarehouse{ID: id}).Update(code).Error
}

func (b BaseWarehouse) Delete(id uint) error {
	return  database.GetDBCon().Delete(&models.BaseWarehouse{},"id = ?",id).Error
}


func (b BaseWarehouse) Create(code models.BaseWarehouse) (models.BaseWarehouse, error) {
	err := database.GetDBCon().Create(&code).Error
	return code, err
}

func (b BaseWarehouse) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (codes []models.BaseWarehouse, total uint, err error) {
	sqlConn := database.GetDBCon().Model(&models.BaseWarehouse{})
	if isTotal {
		if total, err = b.Count(sqlConn, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"base_warehouses.*"}
	}
	sqlConn = b.crud.Where(sqlConn, filter, selectKeys, b.Paginate(per, page), b.OrderBy(orders))
	err = sqlConn.Find(&codes).Error
	return codes, total, err
}

func NewBaseWarehouse() IBaseWarehouse {
	return BaseWarehouse{}
}
