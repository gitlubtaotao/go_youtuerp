package services

import (
	"errors"
	"time"
	"youtuerp/models"
	"youtuerp/repositories"
)

type IOrderMasterService interface {
	Create(order models.OrderMaster, language string) (models.OrderMaster, error)
}
type OrderMasterService struct {
	repo repositories.IOrderMaster
	BaseService
}

func (o OrderMasterService) Create(order models.OrderMaster, language string) (models.OrderMaster, error) {
	order.Status = models.OrderStatusPro
	valid := NewValidatorService(order)
	if message := valid.ResultError(language); message != "" {
		return models.OrderMaster{}, errors.New(message)
	}
	if order.CreatedAt == nil {
		t := time.Now()
		order.CreatedAt = &t
	}
	
	return o.repo.Create(order)
}

func NewOrderMasterService() IOrderMasterService {
	return OrderMasterService{repo: repositories.NewOrderMasterRepository()}
}
