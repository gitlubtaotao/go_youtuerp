package services

import (
	"errors"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
)

type ICrmTrack interface {
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) ([]models.CrmTrack, int64, error)
	Create(track models.CrmTrack, language string) (models.CrmTrack, error)
}

type CrmTrack struct {
	repo dao.ICrmTrack
}

func (c CrmTrack) Create(track models.CrmTrack, language string) (models.CrmTrack, error) {
	valid := NewValidatorService(track)
	if message := valid.ResultError(language); message != "" {
		return models.CrmTrack{}, errors.New(message)
	}
	return c.repo.Create(track)
}

func (c CrmTrack) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) ([]models.CrmTrack, int64, error) {
	return c.repo.Find(per, page, filter, selectKeys, orders, true)
}

func NewCrmTrack() ICrmTrack {
	return &CrmTrack{repo: dao.NewCrmTrack()}
}
