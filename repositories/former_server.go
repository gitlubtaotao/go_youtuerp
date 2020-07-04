package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
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
}

type FormerServer struct {
}

func (f FormerServer) DeleteOtherServer(id uint, formerType string) error {
	switch formerType {
	case "former_trailer_order":
		return f.deleteFormerTrailerOrder(id)
	case "former_other_service":
		return  database.GetDBCon().Delete(models.FormerOtherService{},"id = ?",id).Error
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
	return database.GetDBCon().Model(&models.OrderExtendInfo{ID: id}).Updates(tools.StructToChange(data)).Error
}
func (f FormerServer) GetFormerOtherService(orderMasterId uint) ([]models.FormerOtherService, error) {
	var formerOtherServer []models.FormerOtherService
	err := database.GetDBCon().Where("order_master_id = ?", orderMasterId).Order("id desc").Find(&formerOtherServer).Error
	return formerOtherServer, err
}

func (f FormerServer) GetFormerTrailerOrder(orderMasterId uint) ([]models.FormerTrailerOrder, error) {
	var formerTrailerOrder []models.FormerTrailerOrder
	sqlConn := database.GetDBCon().Where("order_master_id = ?", orderMasterId).Preload("SeaCapLists").Preload("TrailerCabinetNumbers")
	err := sqlConn.Find(&formerTrailerOrder).Error
	return formerTrailerOrder, err
}

//更新货物详情
func (f FormerServer) UpdateSeaCargoInfo(infos []models.SeaCargoInfo) (interface{}, error) {
	sqlConn := database.GetDBCon()
	var data []models.SeaCargoInfo
	err := sqlConn.Transaction(func(tx *gorm.DB) error {
		for _, item := range infos {
			if item.ID != 0 {
				if err := tx.Model(&models.SeaCargoInfo{ID: item.ID}).Update(tools.StructToChange(item)).Error; err != nil {
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

func (f FormerServer) DeleteCargoInfo(ids []int, formerType string) error {
	if formerType == "sea_cargo_info" {
		return database.GetDBCon().Where("id IN (?)", ids).Delete(models.SeaCargoInfo{}).Error
	}
	return nil
}

func (f FormerServer) updateFormerSeaInstruction(data models.RenderFormerData) error {
	sqlConn := database.GetDBCon()
	var record models.FormerSeaInstruction
	instruction := data.FormerSeaInstruction
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
func (f FormerServer) updateFormerSeaBooking(data models.RenderFormerData) error {
	sqlConn := database.GetDBCon()
	var record models.FormerSeaBook
	book := data.FormerSeaBook
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

func (f FormerServer) updateFormerSoNo(formerType string, data models.RenderFormerData) error {
	if formerType == "former_sea_so_no" {
		var soNo = data.FormerSeaSoNo
		golog.Infof("current time is %v", soNo)
		return database.GetDBCon().Model(&models.FormerSeaSoNo{ID: soNo.ID}).Update(tools.StructToChange(soNo)).Error
	}
	return nil
}

//保存拖车单数据
func (f FormerServer) saveFormerTrailerOrder(data models.FormerTrailerOrder) (uint, error) {
	sqlConn := database.GetDBCon()
	if data.ID == 0 {
		err := sqlConn.Create(&data).Error
		return data.ID, err
	}
	var temp models.FormerTrailerOrder
	sqlConn = sqlConn.First(&temp, "id = ? ", data.ID)
	err := sqlConn.Transaction(func(tx *gorm.DB) error {
		if err := sqlConn.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Update(tools.StructToChange(data)).Error; err != nil {
			return err
		}
		if len(data.SeaCapLists) >= 1 {
			if err := tx.Association("SeaCapLists").Replace(data.SeaCapLists).Error; err != nil {
				return err
			}
		}
		if len(data.TrailerCabinetNumbers) >= 1 {
			if err := tx.Association("TrailerCabinetNumbers").Replace(data.TrailerCabinetNumbers).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return data.ID, err
}

//删除拖车单
func (f FormerServer) deleteFormerTrailerOrder(id uint) error {
	err := database.GetDBCon().Transaction(func(tx *gorm.DB) error {
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
		err := database.GetDBCon().Create(&data).Error
		return data.ID, err
	}
	err := database.GetDBCon().Model(models.FormerOtherService{ID: data.ID}).Update(tools.StructToChange(data)).Error
	return data.ID, err
}

func NewFormerServerRepository() IFormerServer {
	return FormerServer{}
}
