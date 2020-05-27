package services

import "youtuerp/repositories"

type IOrderMasterService interface {
}
type OrderMasterService struct {
	repo repositories.IOrderMaster
	BaseService
}

func NewOrderMasterService() IOrderMasterService  {
	return OrderMasterService{repo: repositories.NewOrderMasterRepository()}
}
