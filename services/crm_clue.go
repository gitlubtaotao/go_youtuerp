package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ICrmClueService interface {
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) (clues []models.CrmClue, total uint, err error)
	Create(clue models.CrmClue, language string) (models.CrmClue, error)
}
type CrmClueService struct {
	repo repositories.ICrmClue
}


func (c CrmClueService) Create(clue models.CrmClue, language string) (models.CrmClue, error) {
	valid := NewValidatorService(clue)
	if message := valid.ResultError(language); message != "" {
		err := errors.New(message)
		return models.CrmClue{}, err
	}
	return c.repo.Create(clue)
}

func (c CrmClueService) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) (clues []models.CrmClue, total uint, err error) {
	return c.repo.Find(per, page, filter, selectKeys, orders, true)
}

func NewCrmClueService() ICrmClueService {
	return &CrmClueService{repo: repositories.NewCrmClue()}
}
