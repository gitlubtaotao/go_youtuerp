package models

import "time"

//订单
type OrderMaster struct {
	ID               uint            `gorm:"primary_key"json:"id"`
	CreatedAt        *time.Time      `sql:"index" json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	SerialNumber     string          `gorm:"size:16;unique;not null;index:serial_number" comment:"订单号"  json:"serial_number"`
	InstructionId    uint            `gorm:"index:company_instruction_id" json:"instruction_id" validate:"required"`
	SalesmanId       uint            `gorm:"index:salesman_id" json:"salesman_id" validate:"required" `
	OperationId      uint            `gorm:"index:operation_id" json:"operation_id" validate:"required"`
	TransportType    uint            `sql:"index" json:"transport_type" validate:"required" `
	Status           string          `gorm:"size:16;index:status;default:processing" json:"status" validate:"required"`
	CompanyId        uint            `sql:"index" json:"company_id" validate:"required"`
	ContactId        uint            `json:"contact_id"`
	MainTransport    uint            `json:"main_transport"`
	PayableStatus    string          `gorm:"size:16;index:payable_paid_status;default:unfinished" json:"payable_status"`
	PaidStatus       string          `gorm:"size:16;index:payable_paid_status;default:unfinished" json:"paid_status"`
	ReceivableStatus string          `gorm:"size:16;index:receive_received_status;default:unfinished" sql:"index" json:"receivable_status"`
	ReceivedStatus   string          `gorm:"size:16;index:receive_received_status;default:unfinished" json:"received_status"`
	Remarks          string          `gorm:"size:522" json:"remarks"`
	SupplyAgentId    uint            `sql:"index" json:"supply_agent_id"`
	Roles            []Role          `gorm:"polymorphic:Source;" json:"roles"`
	SeaCargoInfos    []SeaCargoInfo  `gorm:"polymorphic:Source;association_autocreate:false;association_autoupdate:false" json:"sea_cargo_infos"`
	OrderExtendInfo  OrderExtendInfo `gorm:"foreignkey:order_master_id;association_foreignkey:id" json:"order_extend_info"`
}

type OrderExtendInfo struct {
	ID            uint       `gorm:"primary_key"json:"id"`
	OrderMasterId uint       `sql:"index" json:"order_master_id"`
	Number        uint       `comment:"包装数量" json:"number"`
	PackageTypeId uint       `comment:"包装类型" json:"package_type_id"`
	GrossWeight   string     `gorm:"size:64" comment:"毛重" json:"gross_weight"`
	Volume        string     `gorm:"size:64" comment:"体积" json:"volume"`
	MblSO         string     `gorm:"size:16;index:mbl_so" json:"mbl_so"`
	HblSO         string     `gorm:"size:16" json:"hbl_so"`
	SONo          string     `gorm:"size:16;index:so_no" json:"so_no"`
	CarrierId     uint       `sql:"index" json:"carrier_id"`
	POLId         uint       `gorm:"index:pol_pod_and_index" json:"pol_id"`
	PODId         uint       `gorm:"index:pol_pod_and_index" json:"pod_id"`
	POTId         uint       `json:"pot_id"`
	CutOffDay     *time.Time `sql:"index" json:"cut_off_day"`
	Departure     *time.Time `sql:"index" json:"departure"`
	Arrival       *time.Time `sql:"index" json:"arrival"`
	Vessel        string     `gorm:"size:64" json:"vessel"`
	Voyage        string     `gorm:"size:64" json:"voyage"`
	FlightNo      string     `gorm:"size:16;index:flight_no" json:"flight_no"`
	ChargedWeight float64    `gorm:"default:0.0;PRECISION:4" comment:"计费重量" json:"charged_weight"`
	CourierCodeId uint       `sql:"index" json:"courier_code_id"`
	CourierNo     string     `gorm:"size:16;index:courier_no" json:"courier_no"`
	ShipperId     uint       `json:"shipper_id"`
}

type ResultOrderMaster struct {
	ID               uint       `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	SerialNumber     string     `json:"serial_number"`
	InstructionId    uint       `json:"instruction_id"`
	SalesmanId       uint       `json:"salesman_id"`
	OperationId      uint       `json:"operation_id" validate:"required"`
	TransportType    uint       `json:"transport_type" validate:"required" `
	Status           string     `json:"status" validate:"required"`
	CompanyId        uint       `json:"company_id"`
	ContactId        uint       `json:"contact_id"`
	MainTransport    uint       `json:"main_transport"`
	PayableStatus    string     ` json:"payable_status"`
	PaidStatus       string     `json:"paid_status"`
	ReceivableStatus string     `json:"receivable_status"`
	ReceivedStatus   string     `json:"received_status"`
	Remarks          string     `gorm:"size:522" json:"remarks"`
	SupplyAgentId    uint       `sql:"index" json:"supply_agent_id"`
	Number           uint       `json:"number"`
	PackageTypeId    uint       `json:"package_type_id"`
	GrossWeight      string     `gorm:"size:64" comment:"毛重" json:"gross_weight"`
	Volume           string     `gorm:"size:64" comment:"体积" json:"volume"`
	MblSO            string     `gorm:"size:16;index:mbl_so" json:"mbl_so"`
	HblSO            string     `gorm:"size:16" json:"hbl_so"`
	SONo             string     `gorm:"size:16;index:so_no" json:"so_no"`
	CarrierId        uint       `sql:"index" json:"carrier_id"`
	POLId            uint       `gorm:"index:pol_pod_and_index" json:"pol_id"`
	PODId            uint       `gorm:"index:pol_pod_and_index" json:"pod_id"`
	POTId            uint       `json:"pot_id"`
	CutOffDay        *time.Time `sql:"index" json:"cut_off_day"`
	Departure        *time.Time `sql:"index" json:"departure"`
	Arrival          *time.Time `sql:"index" json:"arrival"`
	Vessel           string     `gorm:"size:64" json:"vessel"`
	Voyage           string     `gorm:"size:64" json:"voyage"`
	FlightNo         string     `gorm:"size:16;index:flight_no" json:"flight_no"`
	ChargedWeight    float64    `gorm:"default:0.0;PRECISION:4" comment:"计费重量" json:"charged_weight"`
	CourierCodeId    uint       `sql:"index" json:"courier_code_id"`
	CourierNo        string     `gorm:"size:16;index:courier_no" json:"courier_no"`
	ShipperId        uint       `json:"shipper_id"`
	Roles            []Role     `json:"roles"`
}

const (
	OrderStatusPro      = "processing"
	OrderStatusFinished = "finished"
	OrderStatusLocked   = "locked"
	OrderStatusCancel   = "cancel"
)
const (
	OrderMasterTransportType1 = iota + 1
	OrderMasterTransportType2
	OrderMasterTransportType3
	OrderMasterTransportType4
	OrderMasterTransportType5
)

const (
	FinanceStatusUnfinished = "unfinished"
	FinanceStatusPart       = "part_finished"
	FinanceStatusFinished   = "finished"
)

func (OrderMaster) TableName() string {
	return "order_masters"
}

func (OrderExtendInfo) TableName() string {
	return "order_extend_infos"
}

func (OrderMaster) DefaultAddColumn() []string {
	return []string{"order_extend_infos_carrier_id", "order_extend_infos_pol_id",
		"order_extend_infos_pod_id", "order_extend_infos_departure",
		"order_extend_infos_arrival", "order_extend_infos_hbl_so", "order_extend_infos_mbl_so",
		"order_extend_infos_cargo_info", "order_extend_infos_so_no",
		"order_extend_infos_shipper_id"}
}
