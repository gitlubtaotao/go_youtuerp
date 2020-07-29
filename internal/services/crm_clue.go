package services

import (
	"errors"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
)

type ICrmClueService interface {
	Delete(id uint) error
	Update(id uint, clue models.CrmClue) error
	First(id uint, isTacks bool) (models.CrmClue, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) (clues []models.CrmClue, total int64, err error)
	Create(clue models.CrmClue, language string) (models.CrmClue, error)
}
type CrmClueService struct {
	repo dao.ICrmClue
}

func (c CrmClueService) Delete(id uint) error {
	return c.repo.Delete(id)
}

func (c CrmClueService) Update(id uint, clue models.CrmClue) error {
	return c.repo.Update(id, clue)
}

func (c CrmClueService) First(id uint, isTacks bool) (clue models.CrmClue, err error) {
	clue, err = c.repo.First(id, isTacks)
	if err != nil {
		return
	}
	//查询跟进记录
	if !isTacks {
		return
	}
	service := dao.NewCrmTrack()
	filter := map[string]interface{}{
		"source_id-eq":   clue.ID,
		"source_type-eq": "crm_clues",
	}
	tracks, _, err := service.Find(10, 1, filter, []string{}, []string{}, false)
	if err != nil {
		return
	}
	clue.CrmTracks = tracks
	return
}

func (c CrmClueService) Create(clue models.CrmClue, language string) (models.CrmClue, error) {
	valid := NewValidatorService(clue)
	if message := valid.ResultError(language); message != "" {
		err := errors.New(message)
		return models.CrmClue{}, err
	}
	return c.repo.Create(clue)
}

func (c CrmClueService) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (clues []models.CrmClue, total int64, err error) {
	return c.repo.Find(per, page, filter, selectKeys, orders, true)
}

func NewCrmClueService() ICrmClueService {
	return &CrmClueService{repo: dao.NewCrmClue()}
}
