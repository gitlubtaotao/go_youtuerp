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
	PayableStatus    string          `gorm:"size:16;index:payable_paid_status" json:"payable_status"`
	PaidStatus       string          `gorm:"size:16;index:payable_paid_status" json:"paid_status"`
	ReceivableStatus string          `gorm:"size:16;index:receive_received_status" sql:"index" json:"receivable_status"`
	ReceivedStatus   string          `gorm:"size:16;index:receive_received_status" json:"received_status"`
	Remarks          string          `gorm:"size:522" json:"remarks"`
	SupplyAgentId    uint            `sql:"index" json:"supply_agent_id"`
	Roles            []Role          `gorm:"polymorphic:Source;" json:"roles"`
	SeaCargoInfos    []SeaCargoInfo  `gorm:"polymorphic:Source;association_autocreate:false;association_autoupdate:false" json:"sea_cargo_infos"`
	ExtendInfo       OrderExtendInfo `gorm:"foreignkey:OrderMasterId"`
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
}

const (
	OrderStatusPro = "processing"
)

func (OrderMaster) TableName() string {
	return "order_masters"
}

func (OrderExtendInfo) TableName() string {
	return "order_extend_infos"
}
