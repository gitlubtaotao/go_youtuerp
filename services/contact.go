package services

import "youtuerp/repositories"

type IContactService interface {
}
type ContactService struct {
	repo repositories.IContactRepository
	BaseService
}

func NewContactService() IContactService {
	return &ContactService{repo: repositories.NewContactRepository()}
}
