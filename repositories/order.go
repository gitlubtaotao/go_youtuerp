package repositories

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"sync"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
)

type IOrderMaster interface {
	//保存表单数据
	UpdateFormerData(formerType string, data models.RenderFormerData) error
	UpdateExtendInfo(id uint, data models.OrderExtendInfo) error
	//获取海运委托单
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

func (o OrderMasterRepository) UpdateFormerData(formerType string, data models.RenderFormerData) error {
	sqlConn := database.GetDBCon()
	var err error
	switch formerType {
	case "former_sea_instruction":
		var record models.FormerSeaInstruction
		instruction := data.FormerSeaInstruction
		golog.Infof("sea cap list is %v",instruction.SeaCapLists)
		sqlConn = sqlConn.First(&record, "id = ? ", instruction.ID)
		return sqlConn.Transaction(func(tx *gorm.DB) error {
			if len(instruction.SeaCapLists) >= 1 {
				if err := tx.Association("SeaCapLists").Replace(instruction.SeaCapLists).Error; err != nil {
					return err
				}
			}
			return sqlConn.Set("gorm:association_autocreate", false).Update(tools.StructToChange(instruction)).Error
		})
	case "former_sea_booking":
	}
	return err
}

func (o OrderMasterRepository) UpdateExtendInfo(id uint, data models.OrderExtendInfo) error {
	return database.GetDBCon().Model(&models.OrderExtendInfo{ID: id}).Updates(tools.StructToChange(data)).Error
}

func (o OrderMasterRepository) FormerSeaInstruction(orderMasterId uint, formerType interface{}, attr map[string]interface{}) (models.FormerSeaInstruction, error) {
	attr["order_master_id"] = orderMasterId
	var data models.FormerSeaInstruction
	err := database.GetDBCon().Where(models.FormerSeaInstruction{OrderMasterId: orderMasterId, Type: formerType.(string)}).Preload("SeaCapLists").Attrs(attr).FirstOrCreate(&data).Error
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
	return sqlCon.Transaction(func(tx *gorm.DB) error {
		if err := sqlCon.Association("Roles").Replace(order.Roles).Error; err != nil {
			return err
		}
		err := sqlCon.Set("gorm:association_autocreate", false).Update(order).Error
		return err
	})
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
	err = database.GetDBCon().Create(&models.OrderExtendInfo{OrderMasterId: order.ID, HblSO: order.SerialNumber}).Error
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
