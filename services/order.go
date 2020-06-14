package services

import (
	"errors"
	"github.com/kataras/golog"
	"reflect"
	"time"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/repositories"
	"youtuerp/tools"
)

type IOrderMasterService interface {
	GetFormerData(id uint, formerType string, formerItemType string) (interface{}, error)
	DeleteMaster(id uint) error
	ShowFinanceStatus(enum conf.Enum, field string, value interface{}) string
	ShowTransport(enum conf.Enum, order interface{}) string
	ShowStatus(enum conf.Enum, value interface{}) string
	ChangeStatus(id uint, status string) error
	UpdateMaster(id uint, order models.OrderMaster, language string) error
	FirstMaster(id uint, load ...string) (models.OrderMaster, error)
	FindMaster(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.ResultOrderMaster, uint, error)
	CreateMaster(order models.OrderMaster, language string) (models.OrderMaster, error)
}

var orderStatusArray = []interface{}{
	models.OrderStatusCancel,
	models.OrderStatusPro,
	models.OrderStatusFinished,
	models.OrderStatusLocked,
}

type OrderMasterService struct {
	repo repositories.IOrderMaster
	BaseService
}

func (o OrderMasterService) GetFormerData(id uint, formerType string, formerItemType string) (interface{}, error) {
	var (
		data        interface{}
		err         error
		orderMaster models.OrderMaster
		attr        map[string]interface{}
		crmService  = NewCrmCompanyService()
	)
	if orderMaster, err = o.repo.FirstMaster(id); err != nil {
		return data, err
	}
	golog.Infof("%v",orderMaster)
	attr = make(map[string]interface{}, 2)
	attr["instruction_id"] = orderMaster.InstructionId
	switch formerType {
	case "former_sea_instruction":
		_, status := tools.ContainsSlice([]interface{}{models.InstructionMaster, models.InstructionSplit}, formerItemType)
		if !status {
			return data, errors.New("传入的参数有误")
		}
		if formerItemType == models.InstructionMaster {
			attr["hbl_no"] = orderMaster.SerialNumber
			attr["type"] = formerItemType
			attr["shipper_id"] = orderMaster.InstructionId
			attr["shipper_content"] = crmService.GetOperationInfo(orderMaster.InstructionId)
			data, err = o.repo.FormerSeaInstruction(id, formerItemType, attr)
		} else {
			data, err = o.repo.FormerSeaInstruction(id, formerItemType, attr)
		}
	}
	return data, err
}

func (o OrderMasterService) DeleteMaster(id uint) error {
	return o.repo.DeleteMaster(id)
}

//进行订单状态的更新
func (o OrderMasterService) ChangeStatus(id uint, status string) error {
	_, b := tools.ContainsSlice(orderStatusArray, status)
	if !b {
		return errors.New("状态有误,请重新确认")
	}
	return o.repo.ChangeStatus(id, status)
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

func (o OrderMasterService) ShowStatus(enum conf.Enum, value interface{}) string {
	return enum.DefaultText("order_masters_status.", value)
}

func (o OrderMasterService) ShowTransport(enum conf.Enum, value interface{}) string {
	var (
		TransportType uint
		MainTransport uint
	)
	if reflect.TypeOf(value).Name() == "OrderMaster" {
		record := value.(models.OrderMaster)
		TransportType = record.TransportType
		MainTransport = record.MainTransport
	} else {
		record := value.(models.ResultOrderMaster)
		TransportType = record.TransportType
		MainTransport = record.MainTransport
	}
	temp := enum.DefaultText("order_masters_transport_type.", TransportType)
	if TransportType == models.OrderMasterTransportType3 {
		return temp + enum.DefaultText("order_masters_main_transport.", MainTransport)
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
		repo: repositories.NewOrderMasterRepository(),
	}
}
