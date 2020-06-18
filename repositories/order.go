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
	//获取表单so的信息
	GetFormerSoNo(orderId uint, formerType string, attr ...map[string]interface{}) (interface{}, error)
	//获取订舱单的信息
	GetFormerBooking(orderId uint, formerType string, attr ...map[string]interface{}) (interface{}, error)
	//保存表单数据
	UpdateFormerData(formerType string, data models.RenderFormerData) error
	//保存订单其他数据信息
	UpdateExtendInfo(id uint, data models.OrderExtendInfo) error
	//获取海运委托单
	GetFormerInstruction(orderMasterId uint, formerType interface{}, attr map[string]interface{}) (models.FormerSeaInstruction, error)
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

func (o OrderMasterRepository) GetFormerSoNo(orderId uint, formerType string, attr ...map[string]interface{}) (interface{}, error) {
	var data models.FormerSeaSoNo
	err := database.GetDBCon().Where(models.FormerSeaSoNo{OrderMasterId: orderId}).Attrs(map[string]interface{}{"order_master_id": orderId}).FirstOrCreate(&data).Error
	return data, err
}

func (o OrderMasterRepository) GetFormerBooking(orderId uint, formerType string, attr ...map[string]interface{}) (result interface{}, err error) {
	if formerType == "former_sea_book" {
		var booking models.FormerSeaBook
		var capList []models.SeaCapList
		if len(attr) > 0 {
			if item, ok := attr[0]["sea_cap_lists"]; ok {
				value := item.([]map[string]interface{})
				for _, v := range value {
					capList = append(capList, models.SeaCapList{
						OrderMasterId: v["order_master_id"].(uint),
						Number:        v["number"].(int),
						CapType:       v["cap_type"].(string),
					})
				}
			}
			sqlCon := database.GetDBCon().Where("order_master_id = ?", orderId).Preload("SeaCapLists").Attrs(attr[0]).FirstOrCreate(&booking)
			err = sqlCon.Association("SeaCapLists").Append(capList).Error
		} else {
			err = database.GetDBCon().Preload("SeaCapLists").First(&booking, "order_master_id = ?", orderId).Error
		}
		return booking, err
	}
	return
}

func (o OrderMasterRepository) UpdateFormerData(formerType string, data models.RenderFormerData) error {
	var err error
	switch formerType {
	case "former_sea_instruction":
		err = o.updateFormerSeaInstruction(data)
	case "former_sea_book":
		err = o.updateFormerSeaBooking(data)
	case "former_sea_so_no":
		err = o.updateFormerSoNo(formerType, data)
	}
	return err
}

func (o OrderMasterRepository) UpdateExtendInfo(id uint, data models.OrderExtendInfo) error {
	return database.GetDBCon().Model(&models.OrderExtendInfo{ID: id}).Updates(tools.StructToChange(data)).Error
}

func (o OrderMasterRepository) GetFormerInstruction(orderMasterId uint, formerType interface{}, attr map[string]interface{}) (models.FormerSeaInstruction, error) {
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

func (o OrderMasterRepository) updateFormerSeaInstruction(data models.RenderFormerData) error {
	sqlConn := database.GetDBCon()
	var record models.FormerSeaInstruction
	instruction := data.FormerSeaInstruction
	golog.Infof("sea cap list is %v", instruction.SeaCapLists)
	if err := sqlConn.First(&record, "id = ? ", instruction.ID).Error; err != nil {
		return err
	}
	return sqlConn.Transaction(func(tx *gorm.DB) error {
		if err := sqlConn.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Model(&record).Update(tools.StructToChange(instruction)).Error; err != nil {
			return err
		}
		if len(instruction.SeaCapLists) >= 1 {
			if err := tx.Model(&record).Association("SeaCapLists").Replace(instruction.SeaCapLists).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
func (o OrderMasterRepository) updateFormerSeaBooking(data models.RenderFormerData) error {
	sqlConn := database.GetDBCon()
	var record models.FormerSeaBook
	book := data.FormerSeaBook
	golog.Infof("sea cap list is %v", book.SeaCapLists)
	sqlConn = sqlConn.First(&record, "id = ? ", book.ID)
	return sqlConn.Transaction(func(tx *gorm.DB) error {
		if err := sqlConn.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Update(tools.StructToChange(book)).Error; err != nil {
			return err
		}
		if len(book.SeaCapLists) >= 1 {
			if err := tx.Association("SeaCapLists").Replace(book.SeaCapLists).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (o OrderMasterRepository) updateFormerSoNo(formerType string, data models.RenderFormerData) error {
	if formerType == "former_sea_so_no" {
		var soNo = data.FormerSeaSoNo
		golog.Infof("current time is %v",soNo)
		return database.GetDBCon().Model(&models.FormerSeaSoNo{ID: soNo.ID}).Update(tools.StructToChange(soNo)).Error
	}
	return nil
}

func NewOrderMasterRepository() IOrderMaster {
	return OrderMasterRepository{}
}
