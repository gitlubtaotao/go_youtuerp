package repositories

import (
	"sync"
	"youtuerp/database"
	"youtuerp/models"
)

type IOrderMaster interface {
	Update(id uint, order models.OrderMaster) error
	Create(order models.OrderMaster) (models.OrderMaster, error)
}
type OrderMasterRepository struct {
	BaseRepository
	mu sync.Mutex
}

func (o OrderMasterRepository) Update(id uint, order models.OrderMaster) error {
	var record models.OrderMaster
	sqlCon := database.GetDBCon().First(&record, "id = ? ", id)
	sqlCon.Association("Roles").Replace(order.Roles)
	err := sqlCon.Set("gorm:association_autocreate", false).Update(order).Error
	return err
}


func (o OrderMasterRepository) Create(order models.OrderMaster) (models.OrderMaster, error) {
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
