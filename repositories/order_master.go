package repositories

import (
	"database/sql"
	"gorm.io/gorm"
	"sync"
	"youtuerp/database"
	"youtuerp/models"
)

type IOrderMaster interface {
	//将查询到的订单数据存储到对应的切片中
	
	//通过ids查询订单
	FindMasterByIds(ids []uint, otherFilter ...string) ([]models.ResponseOrderMaster, error)
	//获取表单so的信息
	GetFormerSoNo(orderId uint, formerType string, attr ...map[string]interface{}) (interface{}, error)
	//获取订舱单的信息
	GetFormerBooking(orderId uint, formerType string, attr ...map[string]interface{}) (interface{}, error)
	//获取海运委托单
	GetFormerInstruction(orderMasterId uint, formerType interface{}, attr map[string]interface{}) (models.FormerSeaInstruction, error)
	//删除订单
	DeleteMaster(id uint) error
	//更新订单状态
	ChangeStatus(id uint, status string) error
	//查询订单信息
	FindMaster(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string, isTotal bool) ([]models.ResponseOrderMaster, int64, error)
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

func (o OrderMasterRepository) FindMasterByIds(ids []uint, otherKeys ...string) ([]models.ResponseOrderMaster, error) {
	var orderMasters []models.ResponseOrderMaster
	sqlConn := database.GetDBCon().Model(&models.OrderMaster{}).Where("order_masters.id IN (?)", ids).Scopes(o.joinExtendInfo)
	if len(otherKeys) >= 0 {
		sqlConn = sqlConn.Select(otherKeys)
	}
	rows, err := sqlConn.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data models.ResponseOrderMaster
		_ = sqlConn.ScanRows(rows, &data)
		orderMasters = append(orderMasters, data)
	}
	return orderMasters, err
}

func (o OrderMasterRepository) GetFormerSoNo(orderId uint, formerType string, attr ...map[string]interface{}) (interface{}, error) {
	var data models.FormerSeaSoNo
	err := database.GetDBCon().Where(models.FormerSeaSoNo{OrderMasterId: orderId}).Attrs(map[string]interface{}{"order_master_id": orderId}).FirstOrCreate(&data).Error
	return data, err
}

func (o OrderMasterRepository) GetFormerBooking(orderId uint, formerType string, attr ...map[string]interface{}) (result interface{}, err error) {
	if formerType == "former_sea_book" {
		var booking models.FormerSeaBook
		if len(attr) > 0 {
			return o.createSeaBooking(orderId, attr[0])
		} else {
			err = database.GetDBCon().Preload("SeaCapLists").Preload("SeaCargoInfos").First(&booking, "order_master_id = ?", orderId).Error
		}
		return booking, err
	}
	return
}

func (o OrderMasterRepository) GetFormerInstruction(orderMasterId uint, formerType interface{}, attr map[string]interface{}) (models.FormerSeaInstruction, error) {
	attr["order_master_id"] = orderMasterId
	var data models.FormerSeaInstruction
	err := database.GetDBCon().Where(models.FormerSeaInstruction{OrderMasterId: orderMasterId, Type: formerType.(string)}).Preload("SeaCapLists").Preload("SeaCargoInfos").Attrs(attr).FirstOrCreate(&data).Error
	return data, err
}

func (o OrderMasterRepository) DeleteMaster(id uint) error {
	return o.crud.Delete(&models.OrderMaster{}, id)
}

func (o OrderMasterRepository) ChangeStatus(id uint, status string) error {
	return database.GetDBCon().Model(&models.OrderMaster{ID: id}).Updates(map[string]interface{}{"status": status}).Error
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

func (o OrderMasterRepository) FindMaster(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string, isTotal bool) (masters []models.ResponseOrderMaster, total int64, err error) {
	var rows *sql.Rows
	sqlConn := database.GetDBCon().Model(&models.OrderMaster{}).Scopes(o.joinExtendInfo)
	if isTotal {
		if total, err = o.Count(database.GetDBCon().Model(&models.OrderMaster{}).Scopes(o.joinExtendInfo), filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"order_masters.*", "order_extend_infos.*"}
	}
	if len(orders) == 0 {
		orders = []string{"order_masters.id desc"}
	}
	rows, err = sqlConn.Scopes(o.CustomerWhere(filter, selectKeys, o.Paginate(per, page), o.OrderBy(orders))).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.ResponseOrderMaster
		_ = sqlConn.ScanRows(rows, &data)
		masters = append(masters, data)
	}
	return
}


//
func (o OrderMasterRepository) joinExtendInfo(db *gorm.DB) *gorm.DB {
	return db.Joins("inner join order_extend_infos on order_extend_infos.order_master_id = order_masters.id")
}

func (o OrderMasterRepository) UpdateMaster(id uint, order models.OrderMaster) error {
	var record models.OrderMaster
	sqlCon := database.GetDBCon().First(&record, "id = ? ", id)
	return sqlCon.Transaction(func(tx *gorm.DB) error {
		if err := sqlCon.Association("Roles").Replace(order.Roles); err != nil {
			return err
		}
		err := sqlCon.Set("gorm:association_autocreate", false).Updates(order).Error
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

//创建海运订舱单
func (o OrderMasterRepository) createSeaBooking(orderId uint, attr map[string]interface{}) (models.FormerSeaBook, error) {
	var (
		capList       []models.SeaCapList
		booking       models.FormerSeaBook
		seasCargoInfo []models.SeaCargoInfo
		err           error
	)
	if item, ok := attr["sea_cap_lists"]; ok {
		if value, ok := item.([]models.SeaCapList); ok {
			for _, v := range value {
				capList = append(capList, models.SeaCapList{
					OrderMasterId: v.OrderMasterId,
					Number:        v.Number,
					CapType:       v.CapType,
				})
			}
			attr["sea_cap_lists"] = nil
		}
	}
	if value, ok := attr["sea_cargo_infos"]; ok {
		if seaCargoInfo, ok := value.([]models.SeaCargoInfo); ok {
			for i := 0; i < len(seaCargoInfo); i++ {
				seaCargoInfo[i].ID = 0
				seaCargoInfo[i].SourceID = booking.ID
				seaCargoInfo[i].SourceType = "former_sea_books"
			}
			seasCargoInfo = seaCargoInfo
		}
		attr["sea_cargo_infos"] = nil
	}
	err = database.GetDBCon().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_master_id = ?", orderId).Attrs(attr).FirstOrCreate(&booking).Error; err != nil {
			return err
		}
		if err := tx.Model(&booking).Association("SeaCapLists").Append(capList); err != nil {
			return err
		}
		if err := tx.Model(&booking).Association("SeaCargoInfos").Append(seasCargoInfo); err != nil {
			return err
		}
		return nil
	})
	return booking, err
}

func NewOrderMasterRepository() IOrderMaster {
	return OrderMasterRepository{}
}
