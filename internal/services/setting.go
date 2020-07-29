package services

import (
	"sync"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
	"youtuerp/redis"
)

type ISettingService interface {
	UpdateSystem(key string, setting []models.ResponseSetting) error
	Find(key string) ([]models.Setting, error)
}

type SettingService struct {
	BaseService
	repo dao.ISettingRepository
	sy   sync.Mutex
}

func (s SettingService) Find(key string) ([]models.Setting, error) {
	return s.repo.Find(key)
}

func (s SettingService) UpdateSystem(key string, setting []models.ResponseSetting) error {
	var err error
	err = s.repo.UpdateSystemSetting(key, setting)
	if err != nil {
		return err
	}
	values := make(map[string]interface{})
	s.sy.Lock()
	defer s.sy.Unlock()
	for _, record := range setting {
		values[record.Field] = record.Value
	}
	go s.saveRedis(key, values)
	return nil
}

func (s SettingService) saveRedis(key string, values map[string]interface{}) error {
	red := redis.NewRedis()
	return red.HSetValue(models.Setting{}.TableName(), key, values)
}

func NewSettingService() ISettingService {
	return SettingService{repo: dao.NewSettingRepository()}
}
