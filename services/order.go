package services

import (
	"errors"
	"time"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/repositories"
)

type IOrderMasterService interface {
	ShowFinanceStatus(enum conf.Enum, field string, value interface{}) string
	ShowTransport(enum conf.Enum, record models.ResultOrderMaster) string
	ShowStatus(enum conf.Enum, record models.ResultOrderMaster) string
	UpdateMaster(id uint, order models.OrderMaster, language string) error
	FirstMaster(id uint, load ...string) (models.OrderMaster, error)
	FindMaster(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.ResultOrderMaster, uint, error)
	CreateMaster(order models.OrderMaster, language string) (models.OrderMaster, error)
}

type OrderMasterService struct {
	repo repositories.IOrderMaster
	BaseService
}

func (o OrderMasterService) UpdateMaster(id uint, order models.OrderMaster, language string) error {
	valid := NewValidatorService(order)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	return o.repo.UpdateMaster(id, order)
}

func (o OrderMasterService) FirstMaster(id uint, load ...string) (models.OrderMaster, error) {
	return o.repo.FirstMaster(id, load...)
}

func (o OrderMasterService) ShowFinanceStatus(enum conf.Enum, field string, value interface{}) string {
	if value == "" {
		value = models.FinanceStatusUnfinished
	}
	return enum.DefaultText(field+".", value)
}

func (o OrderMasterService) ShowStatus(enum conf.Enum, record models.ResultOrderMaster) string {
	return enum.DefaultText("order_masters_status.", record.Status)
}

func (o OrderMasterService) ShowTransport(enum conf.Enum, record models.ResultOrderMaster) string {
	temp := enum.DefaultText("order_masters_transport_type.", record.TransportType)
	if record.TransportType == models.OrderMasterTransportType3 {
		return temp + enum.DefaultText("order_masters_main_transport.", record.MainTransport)
	}
	return temp
}

func (o OrderMasterService) FindMaster(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.ResultOrderMaster, uint, error) {
	return o.repo.FindMaster(per, page, filter, selectKeys, orders, true)
}

func (o OrderMasterService) CreateMaster(order models.OrderMaster, language string) (models.OrderMaster, error) {
	order.Status = models.OrderStatusPro
	valid := NewValidatorService(order)
	if message := valid.ResultError(language); message != "" {
		return models.OrderMaster{}, errors.New(message)
	}
	if order.CreatedAt == nil {
		t := time.Now()
		order.CreatedAt = &t
	}
	return o.repo.CreateMaster(order)
}

func NewOrderMasterService() IOrderMasterService {
	return OrderMasterService{
		repo:     repositories.NewOrderMasterRepository(),
	}
}


