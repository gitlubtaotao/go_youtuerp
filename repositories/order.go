package repositories

import (
	"database/sql"
	"sync"
	"youtuerp/database"
	"youtuerp/models"
)

type IOrderMaster interface {
	FormerSeaInstruction(orderMasterId uint, formerType interface{}, attr map[string]interface{}) (models.FormerSeaInstruction, error)
	//删除订单
	DeleteMaster(id uint) error
	//更新订单状态
	ChangeStatus(id uint, status string) error
	//查询订单信息
	FindMaster(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) ([]models.ResultOrderMaster, uint, error)
	//更新订单数据
	UpdateMaster(id uint, order models.OrderMaster) error
	//创建订单
	CreateMaster(order models.OrderMaster) (models.OrderMaster, error)
	//查询第一条订单信息
	FirstMaster(id uint, load ...string) (models.OrderMaster, error)
}
type OrderMasterRepository struct {
	BaseRepository
	mu sync.Mutex
}

func (o OrderMasterRepository) FormerSeaInstruction(orderMasterId uint, formerType interface{}, attr map[string]interface{}) (models.FormerSeaInstruction, error) {
	attr["order_master_id"] = orderMasterId
	var data models.FormerSeaInstruction
	err := database.GetDBCon().Where(models.FormerSeaInstruction{OrderMasterId: orderMasterId, Type: formerType.(string)}).Attrs(attr).FirstOrCreate(&data).Error
	return data, err
}

func (o OrderMasterRepository) DeleteMaster(id uint) error {
	return o.crud.Delete(&models.OrderMaster{}, id)
}

func (o OrderMasterRepository) ChangeStatus(id uint, status string) error {
	return database.GetDBCon().Model(&models.OrderMaster{ID: id}).Update(map[string]interface{}{"status": status}).Error
}

func (o OrderMasterRepository) FirstMaster(id uint, load ...string) (models.OrderMaster, error) {
	var order models.OrderMaster
	sqlConn := database.GetDBCon().Model(&models.OrderMaster{})
	for i := 0; i < len(load); i++ {
		sqlConn = sqlConn.Preload(load[i])
	}
	err := sqlConn.First(&order, "id = ?", id).Error
	return order, err
}

func (o OrderMasterRepository) FindMaster(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (masters []models.ResultOrderMaster, total uint, err error) {
	var rows *sql.Rows
	sqlConn := database.GetDBCon().Model(&models.OrderMaster{})
	sqlConn = sqlConn.Joins("inner join order_extend_infos on order_extend_infos.order_master_id = order_masters.id")
	if isTotal {
		if total, err = o.Count(sqlConn, filter); err != nil {
			return
		}
	}
	sqlConn = o.crud.Where(sqlConn, filter, selectKeys, o.Paginate(per, page), o.OrderBy(orders)).Preload("Roles")
	_ = sqlConn.Error
	rows, err = sqlConn.Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.ResultOrderMaster
		_ = sqlConn.ScanRows(rows, &data)
		masters = append(masters, data)
	}
	return
}

func (o OrderMasterRepository) UpdateMaster(id uint, order models.OrderMaster) error {
	var record models.OrderMaster
	sqlCon := database.GetDBCon().First(&record, "id = ? ", id)
	sqlCon.Association("Roles").Replace(order.Roles)
	err := sqlCon.Set("gorm:association_autocreate", false).Update(order).Error
	return err
}

func (o OrderMasterRepository) CreateMaster(order models.OrderMaster) (models.OrderMaster, error) {
	var (
		serialNumber string
		err          error
	)
	tx := database.GetDBCon().Begin()
	serialNumber, err = NumberSettingRepository{}.GenerateOrderNo(*order.CreatedAt)
	if err != nil {
		tx.Rollback()
		return models.OrderMaster{}, err
	}
	order.SerialNumber = serialNumber
	err = database.GetDBCon().Create(&order).Error
	if err != nil {
		tx.Rollback()
		return models.OrderMaster{}, err
	}
	err = database.GetDBCon().Create(&models.OrderExtendInfo{OrderMasterId: order.ID}).Error
	if err != nil {
		tx.Rollback()
		return models.OrderMaster{}, err
	}
	tx.Commit()
	return order, nil
}

func NewOrderMasterRepository() IOrderMaster {
	return OrderMasterRepository{}
}
