package services

import (
	"errors"
	"reflect"
	"time"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/repositories"
	"youtuerp/tools"
)

type IOrderMasterService interface {
	//通过订单Ids查询订单
	FindMasterByIds(ids []uint, otherFilter ...string) ([]models.ResultOrderMaster, error)
	DeleteCargoInfo(ids []int, formerType string) error
	//保存操作盘中的数据
	UpdateOperationInfo(id uint, formerType string, readData models.RenderFormerData) error
	//保存具体的获取详情
	UpdateCargoInfo(id uint, formerType string, readData models.RenderFormerData) (data interface{}, err error)
	//获取表单中对应的数据
	GetFormerData(id uint, formerType string, formerItemType string) (interface{}, error)
	//获取委托单的数据
	GetFormerInstruction(master models.OrderMaster, formerType string, formerItemType string) (interface{}, error)
	//删除订单
	DeleteMaster(id uint) error
	//显示订单的费用状态
	ShowFinanceStatus(enum conf.Enum, field string, value interface{}) string
	//显示订单的运输类型
	ShowTransport(enum conf.Enum, order interface{}) string
	//显示订单的状态
	ShowStatus(enum conf.Enum, value interface{}) string
	//更改订单的状态
	ChangeStatus(id uint, status string) error
	//更新订单信息
	UpdateMaster(id uint, order models.OrderMaster, language string) error
	//查询订单
	FirstMaster(id uint, load ...string) (models.OrderMaster, error)
	FindMaster(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.ResultOrderMaster, uint, error)
	//创建订单
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

func (o OrderMasterService) FindMasterByIds(ids []uint, otherFilter ...string) ([]models.ResultOrderMaster, error) {
	return o.repo.FindMasterByIds(ids, otherFilter...)
}

func (o OrderMasterService) DeleteCargoInfo(ids []int, formerType string) error {
	return o.repo.DeleteCargoInfo(ids, formerType)
}

func (o OrderMasterService) UpdateCargoInfo(id uint, formerType string, readData models.RenderFormerData) (data interface{}, err error) {
	if formerType == "sea_cargo_info" {
		return o.repo.UpdateSeaCargoInfo(readData.SeaCargoInfo)
	}
	return nil, nil
}

func (o OrderMasterService) UpdateOperationInfo(id uint, formerType string, readData models.RenderFormerData) error {
	extendInfo := readData.OrderExtendInfo
	var err error
	if extendInfo.ID != 0 {
		err = o.repo.UpdateExtendInfo(extendInfo.ID, extendInfo)
	}
	if err != nil {
		return err
	}
	return o.repo.UpdateFormerData(formerType, readData)
}

func (o OrderMasterService) GetFormerData(id uint, formerType string, formerItemType string) (interface{}, error) {
	var (
		data        interface{}
		err         error
		orderMaster models.OrderMaster
	)
	if formerItemType != "former_sea_so_no" {
		if orderMaster, err = o.repo.FirstMaster(id); err != nil {
			return data, err
		}
	}
	switch formerType {
	case "former_sea_instruction":
		data, err = o.GetFormerInstruction(orderMaster, formerType, formerItemType)
	case "former_sea_book":
		data, err = o.getFormerBooking(orderMaster, formerType)
	case "former_sea_so_no":
		data, err = o.repo.GetFormerSoNo(orderMaster.ID, formerItemType)
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
	order.PaidStatus = models.FinanceStatusUnfinished
	order.PayableStatus = models.FinanceStatusUnfinished
	order.ReceivableStatus = models.FinanceStatusUnfinished
	order.ReceivedStatus = models.FinanceStatusUnfinished
	return o.repo.CreateMaster(order)
}

//获取委托单对应的数据
func (o OrderMasterService) GetFormerInstruction(master models.OrderMaster, formerType string, formerItemType string) (interface{}, error) {
	var (
		data       interface{}
		attr       map[string]interface{}
		crmService = NewCrmCompanyService()
		err        error
	)
	attr = make(map[string]interface{}, 2)
	attr["instruction_id"] = master.InstructionId
	attr["hbl_no"] = master.SerialNumber
	attr["type"] = formerItemType
	attr["shipper_id"] = master.InstructionId
	attr["shipper_content"] = crmService.GetOperationInfo(master.InstructionId)
	_, status := tools.ContainsSlice([]interface{}{models.InstructionMaster, models.InstructionSplit}, formerItemType)
	if !status {
		return data, errors.New("传入的参数有误")
	}
	data, err = o.repo.GetFormerInstruction(master.ID, formerItemType, attr)
	return data, err
}

//得到订舱的数据
func (o OrderMasterService) getFormerBooking(order models.OrderMaster, formerType string) (interface{}, error) {
	var (
		data interface{}
		err  error
	)
	if formerType == "former_sea_book" {
		data, err = o.repo.GetFormerBooking(order.ID, formerType)
		if err == nil {
			return data, err
		}
		instruction, err := o.GetFormerInstruction(order, "former_sea_instruction", models.InstructionMaster)
		if err != nil {
			return data, err
		}
		result := o.AutoFillData(models.FormerSeaBook{}, instruction)
		data, err = o.repo.GetFormerBooking(order.ID, formerType, result)
	}
	return data, err
}

//获取自动重填的信息
func (o OrderMasterService) AutoFillData(src interface{}, dst interface{}) map[string]interface{} {
	data := tools.StructToChange(dst)
	dataTypeOf := reflect.TypeOf(dst)
	typeOf := reflect.TypeOf(src)
	result := make(map[string]interface{})
	for i := 0; i < typeOf.NumField(); i++ {
		name := typeOf.Field(i).Tag.Get("json")
		if value, ok := data[name]; ok {
			result[name] = value
		}
	}
	//海运委托单对柜型柜量特殊处理
	if dataTypeOf.Name() == "FormerSeaInstruction" {
		changeData := dst.(models.FormerSeaInstruction)
		result["sea_cap_lists"] = changeData.SeaCapLists
		//todo-tao 创建订舱单时需要把货物信息进行创建
		result["sea_cargo_infos"] = changeData.SeaCargoInfos
	}
	return result
}

func NewOrderMasterService() IOrderMasterService {
	return OrderMasterService{
		repo: repositories.NewOrderMasterRepository(),
	}
}
