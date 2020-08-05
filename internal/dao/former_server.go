package dao

import (
	"github.com/kataras/golog"
	"gorm.io/gorm"
	"youtuerp/global"
	"youtuerp/internal/models"
	"youtuerp/pkg/util"
)

type IFormerServer interface {
	//删除其他综合服务对应的表单
	DeleteOtherServer(id uint, formerType string) error
	//保存其他综合服务的对应的表单
	SaveOtherServer(formerType string, data models.RenderFormerData) (id uint, err error)
	//保存操作former表单的信息
	UpdateFormerData(formerType string, data models.RenderFormerData) error
	//保存订单其他数据信息
	UpdateExtendInfo(id uint, data models.OrderExtendInfo) error
	//更新海运信息
	UpdateSeaCargoInfo(infos []models.SeaCargoInfo) (interface{}, error)
	//根据ids删除货物信息
	DeleteCargoInfo(ids []int, formerType string) error
	//获取其他综合服务
	GetFormerOtherService(orderMasterId uint) ([]models.FormerOtherService, error)
	//获取拖车单
	GetFormerTrailerOrder(orderMasterId uint) ([]models.FormerTrailerOrder, error)
	//获取场装单
	GetFormerWarehouseService(orderMasterId uint) ([]models.FormerWarehouseService, error)
	//获取报关单
	GetFormerCustomClearance(orderMasterId uint) ([]models.FormerCustomClearance, error)
}

type FormerServer struct {
}

func (f FormerServer) DeleteOtherServer(id uint, formerType string) error {
	switch formerType {
	case "former_trailer_order":
		return f.deleteFormerTrailerOrder(id)
	case "former_other_service":
		return global.DataEngine.Delete(models.FormerOtherService{}, "id = ?", id).Error
	case "former_warehouse_service":
		return global.DataEngine.Delete(models.FormerWarehouseService{}, "id = ?", id).Error
	case "former_custom_clearance":
		return global.DataEngine.Delete(models.FormerCustomClearance{}, "id = ?", id).Error
	}
	return nil
}

func (f FormerServer) SaveOtherServer(formerType string, data models.RenderFormerData) (uint, error) {
	var (
		id  uint
		err error
	)
	switch formerType {
	case "former_trailer_order":
		id, err = f.saveFormerTrailerOrder(data.FormerTrailerOrder)
	case "former_other_service":
		id, err = f.saveFormerOtherService(data.FormerOtherService)
	case "former_warehouse_service":
		id, err = f.saveFormerWarehouseService(data.FormerWarehouseService)
	case "former_custom_clearance":
		id, err = f.saveFormerCustomClearance(data.FormerCustomClearance)
	}
	return id, err
}

func (f FormerServer) UpdateFormerData(formerType string, data models.RenderFormerData) error {
	var err error
	switch formerType {
	case "former_sea_instruction":
		err = f.updateFormerSeaInstruction(data)
	case "former_sea_book":
		err = f.updateFormerSeaBooking(data)
	case "former_sea_so_no":
		err = f.updateFormerSoNo(formerType, data)
	case "sea_cargo_info":
		_, err = f.UpdateSeaCargoInfo(data.SeaCargoInfo)
	}
	return err
}

func (f FormerServer) UpdateExtendInfo(id uint, data models.OrderExtendInfo) error {
	return global.DataEngine.Model(&models.OrderExtendInfo{ID: id}).Updates(util.StructToChange(data)).Error
}
func (f FormerServer) GetFormerOtherService(orderMasterId uint) ([]models.FormerOtherService, error) {
	var formerOtherServer []models.FormerOtherService
	err := global.DataEngine.Where("order_master_id = ?", orderMasterId).Order("id desc").Find(&formerOtherServer).Error
	return formerOtherServer, err
}

func (f FormerServer) GetFormerTrailerOrder(orderMasterId uint) ([]models.FormerTrailerOrder, error) {
	var formerTrailerOrder []models.FormerTrailerOrder
	sqlConn := global.DataEngine.Where("order_master_id = ?", orderMasterId).Preload("SeaCapLists").Preload("TrailerCabinetNumbers")
	err := sqlConn.Find(&formerTrailerOrder).Error
	return formerTrailerOrder, err
}

//更新货物详情
func (f FormerServer) UpdateSeaCargoInfo(infos []models.SeaCargoInfo) (interface{}, error) {
	sqlConn := global.DataEngine
	var data []models.SeaCargoInfo
	err := sqlConn.Transaction(func(tx *gorm.DB) error {
		for _, item := range infos {
			if item.ID != 0 {
				if err := tx.Model(&models.SeaCargoInfo{ID: item.ID}).Updates(util.StructToChange(item)).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Create(&item).Error; err != nil {
					return err
				}
				data = append(data, item)
			}
		}
		// 在事务中做一些数据库操作 (这里应该使用 'tx' ，而不是 'db')
		return nil
	})
	return data, err
}

//获取仓库场装单
func (f FormerServer) GetFormerWarehouseService(orderMasterId uint) ([]models.FormerWarehouseService, error) {
	var formerWarehouseService []models.FormerWarehouseService
	err := global.DataEngine.Where("order_master_id = ?", orderMasterId).Order("id desc").Find(&formerWarehouseService).Error
	return formerWarehouseService, err
}

func (f FormerServer) GetFormerCustomClearance(orderMasterId uint) ([]models.FormerCustomClearance, error) {
	var formerCustomerClearance []models.FormerCustomClearance
	err := global.DataEngine.Where("order_master_id = ?", orderMasterId).Order("id desc").Find(&formerCustomerClearance).Error
	return formerCustomerClearance, err
}

func (f FormerServer) DeleteCargoInfo(ids []int, formerType string) error {
	if formerType == "sea_cargo_info" {
		return global.DataEngine.Where("id IN (?)", ids).Delete(models.SeaCargoInfo{}).Error
	}
	return nil
}

func (f FormerServer) updateFormerSeaInstruction(data models.RenderFormerData) error {
	sqlConn := global.DataEngine
	var record models.FormerSeaInstruction
	instruction := data.FormerSeaInstruction
	if err := sqlConn.First(&record, "id = ? ", instruction.ID).Error; err != nil {
		return err
	}
	return sqlConn.Transaction(func(tx *gorm.DB) error {
		if err := sqlConn.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Model(&record).Updates(util.StructToChange(instruction)).Error; err != nil {
			return err
		}
		if len(instruction.SeaCapLists) >= 1 {
			return tx.Model(&record).Association("SeaCapLists").Replace(instruction.SeaCapLists)
		}
		return nil
	})
}
func (f FormerServer) updateFormerSeaBooking(data models.RenderFormerData) error {
	sqlConn := global.DataEngine
	var record models.FormerSeaBook
	book := data.FormerSeaBook
	sqlConn = sqlConn.First(&record, "id = ? ", book.ID)
	return sqlConn.Transaction(func(tx *gorm.DB) error {
		if err := sqlConn.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Updates(util.StructToChange(book)).Error; err != nil {
			return err
		}
		if len(book.SeaCapLists) >= 1 {
			return tx.Association("SeaCapLists").Replace(book.SeaCapLists)
		}
		return nil
	})
}

func (f FormerServer) updateFormerSoNo(formerType string, data models.RenderFormerData) error {
	if formerType == "former_sea_so_no" {
		var soNo = data.FormerSeaSoNo
		golog.Infof("current time is %v", soNo)
		return global.DataEngine.Model(&models.FormerSeaSoNo{ID: soNo.ID}).Updates(util.StructToChange(soNo)).Error
	}
	return nil
}

//保存拖车单数据
func (f FormerServer) saveFormerTrailerOrder(data models.FormerTrailerOrder) (uint, error) {
	sqlConn := global.DataEngine
	if data.ID == 0 {
		err := sqlConn.Create(&data).Error
		return data.ID, err
	}
	var temp models.FormerTrailerOrder
	sqlConn = sqlConn.First(&temp, "id = ? ", data.ID)
	err := sqlConn.Transaction(func(tx *gorm.DB) error {
		if err := sqlConn.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Updates(util.StructToChange(data)).Error; err != nil {
			return err
		}
		if len(data.SeaCapLists) >= 1 {
			if err := tx.Association("SeaCapLists").Replace(data.SeaCapLists); err != nil {
				return err
			}
		}
		if len(data.TrailerCabinetNumbers) >= 1 {
			if err := tx.Association("TrailerCabinetNumbers").Replace(data.TrailerCabinetNumbers); err != nil {
				return err
			}
		}
		return nil
	})
	return data.ID, err
}

//删除拖车单
func (f FormerServer) deleteFormerTrailerOrder(id uint) error {
	err := global.DataEngine.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(models.FormerTrailerOrder{}, " id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Delete(models.TrailerCabinetNumber{}, "former_trailer_order_id = ?", id).Error; err != nil {
			return err
		}
		return tx.Delete(models.SeaCapList{}, "source_id = ? and source_type = ?", id, models.FormerTrailerOrder{}.TableName()).Error
	})
	return err
}

func (f FormerServer) saveFormerOtherService(data models.FormerOtherService) (uint, error) {
	if data.ID == 0 {
		err := global.DataEngine.Create(&data).Error
		return data.ID, err
	}
	err := global.DataEngine.Model(models.FormerOtherService{ID: data.ID}).Updates(util.StructToChange(data)).Error
	return data.ID, err
}

//保存仓库场装单
func (f FormerServer) saveFormerWarehouseService(data models.FormerWarehouseService) (uint, error) {
	if data.ID == 0 {
		err := global.DataEngine.Create(&data).Error
		return data.ID, err
	}
	err := global.DataEngine.Model(models.FormerWarehouseService{ID: data.ID}).Updates(util.StructToChange(data)).Error
	return data.ID, err
}

//保存对应的报关单
func (f FormerServer) saveFormerCustomClearance(data models.FormerCustomClearance) (uint, error) {
	if data.ID == 0 {
		err := global.DataEngine.Create(&data).Error
		return data.ID, err
	}
	err := global.DataEngine.Model(models.FormerCustomClearance{ID: data.ID}).Updates(util.StructToChange(data)).Error
	return data.ID, err
}

func NewFormerServerRepository() IFormerServer {
	return FormerServer{}
}
