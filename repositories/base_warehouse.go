package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type IBaseWarehouse interface {
	//根据不同的select keys 和orderBy 查询所有的仓库地址信息
	FindAll(selectKeys []string, orderBy string) ([]models.BaseWarehouse, error)
	//更新仓库地址信息
	Update(id uint, code models.BaseWarehouse) error
	//删除创建地址信息
	Delete(id uint) error
	//创建仓库地址信息
	Create(code models.BaseWarehouse) (models.BaseWarehouse, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) (codes []models.BaseWarehouse, total uint, err error)
}
type BaseWarehouse struct {
	BaseRepository
}

func (b BaseWarehouse) FindAll(selectKeys []string, orderBy string) ([]models.BaseWarehouse, error) {
	var records [] models.BaseWarehouse
	sqlCon := database.GetDBCon()
	if len(selectKeys) > 0 {
		sqlCon = sqlCon.Select(selectKeys)
	}
	if orderBy != "" {
		sqlCon = sqlCon.Order(orderBy)
	}
	err := sqlCon.Find(&records).Error
	return records, err
}

func (b BaseWarehouse) Update(id uint, code models.BaseWarehouse) error {
	return database.GetDBCon().Model(&models.BaseWarehouse{ID: id}).Update(code).Error
}

func (b BaseWarehouse) Delete(id uint) error {
	return database.GetDBCon().Delete(&models.BaseWarehouse{}, "id = ?", id).Error
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
