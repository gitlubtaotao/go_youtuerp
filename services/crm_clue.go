package services

import "youtuerp/repositories"

type ICrmClueService interface {
}
type CrmClueService struct {
	repo repositories.ICrmClue
}

func NewCrmClueService() ICrmClueService {
	return &CrmClueService{repo: repositories.NewCrmClue()}
}
