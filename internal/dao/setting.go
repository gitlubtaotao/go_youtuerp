package dao

import (
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/internal/models"
)

type ISettingRepository interface {
	Find(key string) ([]models.Setting, error)
	UpdateSystemSetting(key string, setting []models.ResponseSetting) error
}

type SettingRepository struct {
}

func (s SettingRepository) Find(key string) ([]models.Setting, error) {
	var result []models.Setting
	if key == "system" {
		err := database.GetDBCon().Where("user_id = 0").Find(&result).Error
		return result, err
	} else {
		return result, nil
	}
}

type Value []interface{}

func (s SettingRepository) UpdateSystemSetting(key string, setting []models.ResponseSetting) error {
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
