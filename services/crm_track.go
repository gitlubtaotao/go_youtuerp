package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ICrmTrack interface {
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) ([]models.CrmTrack, uint, error)
	Create(track models.CrmTrack, language string) (models.CrmTrack, error)
}

type CrmTrack struct {
	repo repositories.ICrmTrack
}

func (c CrmTrack) Create(track models.CrmTrack, language string) (models.CrmTrack, error) {
	valid := NewValidatorService(track)
	if message := valid.ResultError(language); message != "" {
		return models.CrmTrack{}, errors.New(message)
	}
	return c.repo.Create(track)
}


func (c CrmTrack) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) ([]models.CrmTrack, uint, error) {
	return c.repo.Find(per, page, filter, selectKeys, orders, true)
}

func NewCrmTrack() ICrmTrack {
	return &CrmTrack{repo: repositories.NewCrmTrack()}
}
