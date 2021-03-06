package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type IBasePort interface {
	Update(id uint, code models.BaseDataPort) error
	Delete(id uint) error
	Create(code models.BaseDataPort) (models.BaseDataPort, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) (codes []models.BaseDataPort, total uint, err error)
}


type BasePort struct {
	BaseRepository
}

func (b BasePort) Update(id uint, carrier models.BaseDataPort) error {
	return database.GetDBCon().Model(&models.BaseDataPort{ID: id}).Update(carrier).Error
}

func (b BasePort) Delete(id uint) error {
	return b.crud.Delete(&models.BaseDataPort{}, id)
}
func (b BasePort) Create(code models.BaseDataPort) (models.BaseDataPort, error) {
	err := database.GetDBCon().Create(&code).Error
	return code, err
}



func (b BasePort) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (codes []models.BaseDataPort, total uint, err error) {
	sqlConn := database.GetDBCon().Model(&models.BaseDataPort{})
	if isTotal {
		if total, err = b.Count(sqlConn, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"base_data_ports.*"}
	}
	sqlConn = b.crud.Where(sqlConn, filter, selectKeys, b.Paginate(per, page), b.OrderBy(orders))
	err = sqlConn.Find(&codes).Error
	return codes, total, err
}

func NewBasePort() IBasePort {
	return &BasePort{}
}
