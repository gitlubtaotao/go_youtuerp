package services

import (
	"sync"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ISettingService interface {
	UpdateSystem(key string, setting []models.ResultSetting) error
}

type SettingService struct {
	BaseService
	repo repositories.ISettingRepository
	sy   sync.Mutex
}

func (s SettingService) UpdateSystem(key string, setting []models.ResultSetting) error {
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
	return s.saveRedis(key, values)
}

func (s SettingService) saveRedis(key string, values map[string]interface{}) error {
	return conf.ReisCon.HSet(key, values).Err()
}

func NewSettingService() ISettingService {
	return SettingService{repo: repositories.NewSettingRepository()}
}
