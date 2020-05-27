package models

import "time"

type FormerSeaInstruction struct {
	ID                  uint           `gorm:"primary_key"json:"id"`
	CreatedAt           time.Time      `json:"created_at"`
	Type                string         `gorm:"size:16" comment:"委托类型" sql:"index" json:"type"`
	OrderMasterId       uint           `sql:"index" json:"order_master_id"`
	UpdatedAt           time.Time      `json:"updated_at"`
	InstructionId       uint           `sql:"index" json:"company_instruction_id"`
	ShipperId           uint           `json:"shipper_id"`
	ShipperContent      string         `gorm:"size:1024" json:"shipper_content"`
	ConsigneeId         uint           `json:"consignee_id"`
	ConsigneeContent    string         `gorm:"size:1024" json:"consignee_content"`
	NotifyPartyId       uint           `json:"notify_party_id"`
	NotifyPartyContent  string         `gorm:"size:1024" json:"notify_party_content"`
	PodAgentId          uint           `json:"pod_agent_id"`
	PodAgentContent     string         `gorm:"size:1024" json:"pod_agent_content"`
	HblNo               string         `gorm:"size:16" json:"hbl_no"`
	OceanChangePayId    uint           `json:"ocean_change_pay_id"`
	OtherChangePayId    uint           `json:"other_change_pay_id"`
	VerifyDate          *time.Time     `json:"verify_date"`
	UserVerifyId        uint           `json:"user_verify_id"`
	TlxNo               string         `gorm:"size:16" json:"tlx_no"`
	HblNoDate           *time.Time     `json:"hbl_no_date"`
	PlaceOfIssue        string         `gorm:"size:64" json:"place_of_issue"`
	DateOfIssue         *time.Time     `json:"date_of_issue"`
	CargoReceivedDate   *time.Time     `json:"cargo_received_date"`
	Marks               string         `gorm:"size:2048" json:"marks"`
	DescriptionOfGood   string         `gorm:"size:2048" json:"description_of_good"`
	Size                string         `gorm:"size:522" json:"size"`
	Number              uint           `json:"number"`
	PackageTypeId       uint           `json:"package_type_id"`
	GrossWeight         string         `gorm:"size:64" json:"gross_weight"`
	Volume              string         `gorm:"size:64" comment:"体积" json:"volume"`
	PlaceOfDelivery     string         `gorm:"size:1024" json:"place_of_delivery"`
	PlaceOfReceipt      string         `gorm:"size:1024" json:"place_of_receipt"`
	HblRemarks          string         `gorm:"size:1024" json:"hbl_remarks"`
	InstructionTypeId   uint           `json:"instruction_type_id"`
	WaysOfDeclarationId uint           `json:"ways_of_declaration_id"`
	MiscBillId          uint           `comment:"出单方式" json:"misc_bill_id"`
	TransshipmentTypeId uint           `comment:"转运条款" json:"transshipment_type_id"`
	TradeTermsId        uint           `comment:"贸易条款" json:"trade_terms_id"`
	ShipmentItemId      uint           `comment:"装运条款" json:"shipment_item_id"`
	ContractNo          string         `gorm:"size:32" json:"contract_no"`
	InvoiceNo           string         `gorm:"size:32" json:"invoice_no"`
	Remarks             string         `gorm:"size:1024" json:"remarks"`
	SeaCapLists         []SeaCapList   `gorm:"polymorphic:Source;" json:"sea_cap_lists"`
	SeaCargoInfos       []SeaCargoInfo `gorm:"polymorphic:Source;" json:"sea_cargo_infos"`
}

type FormerSeaBook struct {
	ID                 uint         `gorm:"primary_key"json:"id"`
	CreatedAt          time.Time    `json:"created_at"`
	OrderMasterId      uint         `sql:"index" json:"order_master_id"`
	UpdatedAt          time.Time    `json:"updated_at"`
	InstructionId      uint         `sql:"index" json:"company_instruction_id"`
	ShipperId          uint         `json:"shipper_id"`
	ShipperContent     string       `gorm:"size:1024" json:"shipper_content"`
	ConsigneeId        uint         `json:"consignee_id"`
	ConsigneeContent   string       `gorm:"size:1024" json:"consignee_content"`
	NotifyPartyId      uint         `json:"notify_party_id"`
	NotifyPartyContent string       `gorm:"size:1024" json:"notify_party_content"`
	PodAgentId         uint         `json:"pod_agent_id"`
	PodAgentContent    string       `gorm:"size:1024" json:"pod_agent_content"`
	MblNo              string       `gorm:"size:16" json:"mbl_no"`
	OceanChangePayId   uint         `json:"ocean_change_pay_id"`
	OtherChangePayId   uint         `json:"other_change_pay_id"`
	VerifyDate         *time.Time   `json:"verify_date"`
	UserVerifyId       uint         `json:"user_verify_id"`
	TlxNo              string       `gorm:"size:16" json:"tlx_no"`
	MblNoDate          *time.Time   `json:"mbl_no_date"`
	PlaceOfIssue       string       `gorm:"size:64" json:"place_of_issue"`
	DateOfIssue        *time.Time   `json:"date_of_issue"`
	CargoReceivedDate  *time.Time   `json:"cargo_received_date"`
	Marks              string       `gorm:"size:2048" json:"marks"`
	DescriptionOfGood  string       `gorm:"size:2048" json:"description_of_good"`
	Size               string       `gorm:"size:522" json:"size"`
	Number             uint         `json:"number"`
	PackageTypeId      uint         `json:"package_type_id"`
	GrossWeight        string       `gorm:"size:64" json:"gross_weight"`
	Volume             string       `gorm:"size:64" comment:"体积" json:"volume"`
	PlaceOfDelivery    string       `gorm:"size:1024" json:"place_of_delivery"`
	PlaceOfReceipt     string       `gorm:"size:1024" json:"place_of_receipt"`
	MblRemarks         string       `gorm:"size:1024" json:"mbl_remarks"`
	SupplyAgentId      uint         `sql:"index" json:"supply_agent_id"`
	MiscBillId         uint         `json:"misc_bill_id"`
	TradeTermsId       uint         `comment:"贸易条款" json:"trade_terms_id"`
	ShipmentItemId     uint         `comment:"装运条款" json:"shipment_item_id"`
	Remarks            string       `gorm:"size:1024" json:"remarks"`
	ChargeDescription  string       `gorm:"size:1024" json:"charge_description"`
	SeaCapLists        []SeaCapList `gorm:"polymorphic:Source;" json:"sea_cap_lists"`
}

type FormerSeaSoNo struct {
	ID                    uint       `gorm:"primary_key"json:"id"`
	CreatedAt             time.Time  `json:"created_at"`
	OrderMasterId         uint       `sql:"index" json:"order_master_id"`
	UpdatedAt             time.Time  `json:"updated_at"`
	SoNo                  string     `gorm:"size:1024" json:"so_no"`
	CyOpenDate            *time.Time `json:"cy_open_date"`
	VoucherCutOff         *time.Time `json:"voucher_cut_off"`
	VgmSubmissionDeadline *time.Time `json:"vgm_submission_deadline"`
	SiCutOff              *time.Time `json:"si_cut_off"`
}
type SeaCapList struct {
	ID            uint      `gorm:"primary_key"json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	OrderMasterId uint      `sql:"index" json:"order_master_id"`
	SourceId      uint      `gorm:"index:source_id_and_source_type" json:"source_id"`
	SourceType    string    `gorm:"size:32; index:source_id_and_source_type" json:"source_type"`
	Number        uint      `json:"number"`
	CapType       uint      `json:"cap_type"`
}

type SeaCargoInfo struct {
	ID                uint      `gorm:"primary_key"json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	SourceId          uint      `gorm:"index:source_id_and_source_type" json:"source_id"`
	SourceType        string    ` gorm:"size:32;index:source_id_and_source_type" json:"source_type"`
	OrderMasterId     uint      `sql:"index" json:"order_master_id"`
	SoNo              string    `gorm:"size:64" json:"so_no"`
	ContainerNo       string    `gorm:"size:64" json:"container_no"`
	SealNo            string    `gorm:"size:64" json:"seal_no"`
	VerifiedGrossMass string    `gorm:"size:32" json:"verified_gross_mass"`
	IncludeContainer  uint      `comment:"称重方式" json:"include_container"`
	GrossUnit         uint      `comment:"重量单位"`
	GrossWeight       string    `gorm:"size:32" comment:"毛重" json:"gross_weight"`
	ContainerWeight   string    `gorm:"size:32" comment:"柜重" json:"container_weight"`
	VgmWeight         string    `gorm:"size:32" comment:"VGM重量" json:"vgm_weight"`
	Volume            string    `gorm:"size:32" json:"volume"`
	Number            uint      `json:"number"`
	PackageTypeId     uint      `json:"package_type_id"`
	CapTypeId         uint      `json:"cap_type_id"`
	DescriptionOfGood string    `gorm:"size:256" json:"description_of_good"`
	Marks             string    `gorm:"size:256" json:"marks"`
}

func (FormerSeaInstruction) TableName() string {
	return "former_sea_instructions"
}

func (FormerSeaBook) TableName() string {
	return "former_sea_books"
}
func (FormerSeaSoNo) TableName() string {
	return "former_sea_so_nos"
}

func (SeaCargoInfo) TableName() string {
	return "sea_cargo_infos"
}
func (SeaCapList) TableName() string {
	return "sea_cap_lists"
}
