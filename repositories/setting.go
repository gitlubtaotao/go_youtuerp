package repositories

import (
	"github.com/jinzhu/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type ISettingRepository interface {
	UpdateSystemSetting(key string, setting []models.ResultSetting) error
}

type SettingRepository struct {
}
type Value []interface{}

func (s SettingRepository) UpdateSystemSetting(key string, setting []models.ResultSetting) error {
	return database.GetDBCon().Transaction(func(tx *gorm.DB) error {
		for _, record := range setting {
			err := database.GetDBCon().Where(models.Setting{Field: record.Field, Key: key}).Assign(models.Setting{Value: record.Value}).FirstOrCreate(&models.Setting{}).Error
			if err != nil {
				return err
			}
		}
		//返回 nil 提交事务
		return nil
	})
}

func NewSettingRepository() ISettingRepository {
	return &SettingRepository{}
}
