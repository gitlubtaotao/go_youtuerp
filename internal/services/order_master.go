package services

import (
	"errors"
	"github.com/kataras/golog"
	"reflect"
	"strconv"
	"time"
	"youtuerp/conf"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
	"youtuerp/pkg/util"
	"youtuerp/redis"
)

type IOrderMasterService interface {
	//获取订单数据
	FindMasterByIndex(per, page int, filter map[string]interface{}, selectKeys []string, orders []string, enum conf.Enum) ([]map[string]interface{}, int64, error)
	//查询订单信息，不需要进行total的统计
	FindMasterNoTotal(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.ResponseOrderMaster, int64, error)
	//前端显示订单信息预处理
	HandlerOrderMasterShow(order interface{}, enum conf.Enum) map[string]interface{}
	//处理港口显示
	ShowPort(transportType interface{}, portId interface{}) string
	//处理承运方显示
	ShowCarrier(transportType interface{}, carrierId interface{}) string
	//通过订单Ids查询订单
	FindMasterByIds(ids []uint, otherFilter ...string) ([]models.ResponseOrderMaster, error)
	//获取表单中对应的数据
	GetFormerData(id uint, formerType string) (interface{}, error)
	//获取委托单的数据
	GetSeaFormerInstruction(master models.OrderMaster, formerItemType string) (interface{}, error)
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
	FindMaster(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.ResponseOrderMaster, int64, error)
	//创建订单
	CreateMaster(order models.OrderMaster, language string) (models.OrderMaster, error)
}

var orderStatusArray = []interface{}{
	models.OrderStatusCancel,
	models.OrderStatusPro,
	models.OrderStatusFinished,
	models.OrderStatusLocked,
	models.OrderStatusAudit,
	models.OrderStatusTakeOrder,
}

type OrderMasterService struct {
	repo dao.IOrderMaster
	BaseService
}

func (o OrderMasterService) FindMasterByIndex(per, page int, filter map[string]interface{}, selectKeys []string, orders []string, enum conf.Enum) ([]map[string]interface{}, int64, error) {
	orderMasters, total, err := o.repo.FindMaster(per, page, filter, selectKeys, orders, true)
	if err != nil {
		return []map[string]interface{}{}, 0, err
	}
	data := make([]map[string]interface{}, len(orderMasters))
	for i, record := range orderMasters {
		data[i] = o.HandlerOrderMasterShow(record, enum)
	}
	return data, total, nil
}

func (o OrderMasterService) FindMasterNoTotal(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.ResponseOrderMaster, int64, error) {
	return o.repo.FindMaster(per, page, filter, selectKeys, orders, false)
}

func (o OrderMasterService) ShowPort(transportType interface{}, portId interface{}) string {
	tableName := models.BaseDataPort{}.TableName()
	uTransportType := transportType.(uint)
	var key string
	switch uTransportType {
	case 1, 4:
		key = strconv.Itoa(models.BaseTypeSea)
	case 2:
		key = strconv.Itoa(models.BaseTypeAir)
	}
	return red.HGetValue(tableName+key, portId, "name")
}

func (o OrderMasterService) ShowCarrier(transportType interface{}, carrierId interface{}) string {
	tableName := models.BaseDataCarrier{}.TableName()
	uTransportType := transportType.(uint)
	var key string
	switch uTransportType {
	case 1, 4:
		key = strconv.Itoa(models.BaseTypeSea)
	case 2:
		key = strconv.Itoa(models.BaseTypeAir)
	}
	return red.HGetValue(tableName+key, carrierId, "name")
}

func (o OrderMasterService) HandlerOrderMasterShow(order interface{}, enum conf.Enum) map[string]interface{} {
	data := util.StructToMap(order)
	data["cut_off_day"] = toolTime.InterfaceFormat(data["cut_off_day"], "zh-CN")
	data["departure"] = toolTime.InterfaceFormat(data["departure"], "zh-CN")
	data["arrival"] = toolTime.InterfaceFormat(data["arrival"], "zh-CN")
	data["instruction_id_value"] = data["instruction_id"]
	data["instruction_id"] = red.HGetCrm(data["instruction_id"], "")
	data["supply_agent_id"] = red.HGetCrm(data["supply_agent_id"], "")
	data["salesman_id_value"] = data["salesman_id"]
	data["salesman_id"] = red.HGetRecord("users", data["salesman_id"], "")
	data["operation_id_value"] = data["operation_id"]
	data["operation_id"] = red.HGetRecord("users", data["operation_id"], "")
	data["pol_id"] = o.ShowPort(data["transport_type"], data["pol_id"])
	data["pod_id"] = o.ShowPort(data["transport_type"], data["pod_id"])
	data["carrier_id"] = o.ShowCarrier(data["transport_type"], data["carrier_id"])
	data["transport_type_value"] = data["transport_type"]
	data["transport_type"] = o.ShowTransport(enum, order)
	data["status_value"] = data["status"]
	data["status"] = o.ShowStatus(enum, data["status"])
	data["company_id"] = red.HGetCompany(data["company_id"], "")
	for _, v := range []string{"paid_status", "received_status", "payable_status", "receivable_status"} {
		data[v+"_value"] = data[v]
		data[v] = o.ShowFinanceStatus(enum, v, data[v])
	}
	return data
}

func (o OrderMasterService) FindMasterByIds(ids []uint, otherFilter ...string) ([]models.ResponseOrderMaster, error) {
	return o.repo.FindMasterByIds(ids, otherFilter...)
}

func (o OrderMasterService) GetFormerData(id uint, formerType string) (interface{}, error) {
	var (
		data        interface{}
		err         error
		orderMaster models.OrderMaster
	)
	orderMaster, err = o.FirstMaster(id)
	if err != nil {
		return nil, err
	}
	switch formerType {
	case "former_sea_instruction":
		data, err = o.GetSeaFormerInstruction(orderMaster, models.InstructionMaster)
	case "former_sea_book":
		data, err = o.getFormerBooking(orderMaster, formerType)
	case "former_sea_so_no":
		data, err = o.repo.GetFormerSoNo(orderMaster.ID, formerType)
	}
	return data, err
}

func (o OrderMasterService) DeleteMaster(id uint) error {
	return o.repo.DeleteMaster(id)
}

//进行订单状态的更新
func (o OrderMasterService) ChangeStatus(id uint, status string) error {
	_, b := util.ContainsSlice(orderStatusArray, status)
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
		record := value.(models.ResponseOrderMaster)
		TransportType = record.TransportType
		MainTransport = record.MainTransport
	}
	temp := enum.DefaultText("order_masters_transport_type.", TransportType)
	if TransportType == models.OrderMasterTransportType3 {
		return temp + enum.DefaultText("order_masters_main_transport.", MainTransport)
	}
	return temp
}

func (o OrderMasterService) FindMaster(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) ([]models.ResponseOrderMaster, int64, error) {
	golog.Infof("find master is %v", per)
	return o.repo.FindMaster(per, page, filter, selectKeys, orders, true)
}

func (o OrderMasterService) CreateMaster(order models.OrderMaster, language string) (models.OrderMaster, error) {
	if redis.OrderAuditMechanism() == "false" {
		order.Status = models.OrderStatusPro
	} else {
		order.Status = models.OrderStatusAudit
	}
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
func (o OrderMasterService) GetSeaFormerInstruction(master models.OrderMaster, formerItemType string) (interface{}, error) {
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
	_, status := util.ContainsSlice([]interface{}{models.InstructionMaster, models.InstructionSplit}, formerItemType)
	if !status {
		return data, errors.New("传入的参数有误")
	}
	data, err = o.repo.GetFormerInstruction(master.ID, formerItemType, attr)
	return data, err
}

//得到海运订舱的数据
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
		instruction, err := o.GetSeaFormerInstruction(order, models.InstructionMaster)
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
	data := util.StructToChange(dst)
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
		repo: dao.NewOrderMasterRepository(),
	}
}
