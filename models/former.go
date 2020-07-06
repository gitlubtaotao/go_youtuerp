package models

import "time"

type FormerSeaInstruction struct {
	ID                  uint           `gorm:"primary_key"json:"id"`
	CreatedAt           time.Time      `json:"created_at"`
	Type                string         `gorm:"size:16" comment:"委托类型" sql:"index" json:"type"`
	OrderMasterId       uint           `sql:"index" json:"order_master_id"`
	UpdatedAt           time.Time      `json:"updated_at"`
	InstructionId       uint           `sql:"index" json:"instruction_id"`
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
	ShapingOnBoardDate  *time.Time     `json:"shaping_on_board_date"`
	SeaCapLists         []SeaCapList   `gorm:"polymorphic:Source;" json:"sea_cap_lists"`
	SeaCargoInfos       []SeaCargoInfo `gorm:"polymorphic:Source;" json:"sea_cargo_infos"`
}

type FormerSeaBook struct {
	ID                 uint           `gorm:"primary_key"json:"id"`
	CreatedAt          time.Time      `json:"created_at"`
	OrderMasterId      uint           `sql:"index" json:"order_master_id"`
	UpdatedAt          time.Time      `json:"updated_at"`
	InstructionId      uint           `sql:"index" json:"company_instruction_id"`
	ShipperId          uint           `json:"shipper_id"`
	ShipperContent     string         `gorm:"size:1024" json:"shipper_content"`
	ConsigneeId        uint           `json:"consignee_id"`
	ConsigneeContent   string         `gorm:"size:1024" json:"consignee_content"`
	NotifyPartyId      uint           `json:"notify_party_id"`
	NotifyPartyContent string         `gorm:"size:1024" json:"notify_party_content"`
	PodAgentId         uint           `json:"pod_agent_id"`
	PodAgentContent    string         `gorm:"size:1024" json:"pod_agent_content"`
	MblNo              string         `gorm:"size:16" json:"mbl_no"`
	OceanChangePayId   uint           `json:"ocean_change_pay_id"`
	OtherChangePayId   uint           `json:"other_change_pay_id"`
	VerifyDate         *time.Time     `json:"verify_date"`
	UserVerifyId       uint           `json:"user_verify_id"`
	PayPolId           uint           `json:"pay_pol_id"`
	TlxNo              string         `gorm:"size:16" json:"tlx_no"`
	MblNoDate          *time.Time     `json:"mbl_no_date"`
	PlaceOfIssue       string         `gorm:"size:64" json:"place_of_issue"`
	DateOfIssue        *time.Time     `json:"date_of_issue"`
	CargoReceivedDate  *time.Time     `json:"cargo_received_date"`
	Marks              string         `gorm:"size:2048" json:"marks"`
	DescriptionOfGood  string         `gorm:"size:2048" json:"description_of_good"`
	Size               string         `gorm:"size:522" json:"size"`
	Number             int            `json:"number"`
	PackageTypeId      uint           `json:"package_type_id"`
	GrossWeight        string         `gorm:"size:64" json:"gross_weight"`
	Volume             string         `gorm:"size:64" comment:"体积" json:"volume"`
	PlaceOfDelivery    string         `gorm:"size:1024" json:"place_of_delivery"`
	PlaceOfReceipt     string         `gorm:"size:1024" json:"place_of_receipt"`
	ShapingOnBoardDate *time.Time     `json:"shaping_on_board_date"`
	MblRemarks         string         `gorm:"size:1024" json:"mbl_remarks"`
	SupplyAgentId      uint           `sql:"index" json:"supply_agent_id"`
	MiscBillId         uint           `json:"misc_bill_id"`
	TradeTermsId       uint           `comment:"贸易条款" json:"trade_terms_id"`
	ShipmentItemId     uint           `comment:"装运条款" json:"shipment_item_id"`
	Remarks            string         `gorm:"size:1024" json:"remarks"`
	ChargeDescription  string         `gorm:"size:1024" json:"charge_description"`
	SeaCapLists        []SeaCapList   `gorm:"polymorphic:Source;" json:"sea_cap_lists"`
	SeaCargoInfos      []SeaCargoInfo `gorm:"polymorphic:Source;" json:"sea_cargo_infos"`
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

//其他服务类型
type FormerOtherService struct {
	ID                     uint      `gorm:"primary_key"json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	InstructionId          uint      `gorm:"index:idx_instruction_id" json:"instruction_id"`
	FumigationCompanyId    uint      `json:"fumigation_company_id" comment:"熏蒸公司" `
	TraderCompanyId        uint      `json:"trader_company_id"`
	CommodityInspectionId  uint      `json:"commodity_inspection_id"`
	MagneticInspectionId   uint      `json:"magnetic_inspection_id"`
	AccreditationCompanyId uint      `json:"accreditation_company_id"`
	InsuranceCompanyId     uint      `json:"insurance_company_id"`
	Remarks                string    `gorm:"size:512" json:"remarks"`
	OrderMasterId          uint      `gorm:"index:idx_order_master_id" json:"order_master_id"`
}

//拖车单
type FormerTrailerOrder struct {
	ID                    uint                   `gorm:"primary_key"json:"id"`
	CreatedAt             time.Time              `json:"created_at"`
	OrderMasterId         uint                   `gorm:"index:idx_order_master_id" json:"order_master_id"`
	InstructionId         uint                   `gorm:"index:idx_instruction_id" json:"instruction_id"`
	TrailerCompanyId      uint                   `json:"trailer_company_id"`
	TrailerContactName    string                 `gorm:"size:32" json:"trailer_contact_name"`
	TrailerContactPhone   string                 `gorm:"size:16" json:"trailer_contact_phone"`
	TrailerNumber         string                 `gorm:"size:16" json:"trailer_number" comment:"车牌号码"`
	OfWay                 uint                   `json:"of_way"`
	OfType                uint                   `gorm:"index:idx_of_type" json:"of_type"`
	SoNo                  string                 `gorm:"size:16" json:"so_no"`
	LoadingDate           *time.Time             `json:"loading_date"`
	PolId                 uint                   `json:"pol_id"`
	IsDrivingLicense      bool                   `json:"is_driving_license" comment:"转关带司机本"`
	IsDeclare             bool                   `json:"is_declare" comment:"报关单证随车"`
	IsWeighing            bool                   `json:"is_weighing" comment:"需要过磅"`
	IsLockers             bool                   `json:"is_lockers"`
	Marks                 string                 `gorm:"size:2048" json:"marks"`
	DescriptionOfGood     string                 `gorm:"size:2048" json:"description_of_good"`
	Size                  string                 `gorm:"size:522" json:"size"`
	Number                int                    `json:"number"`
	PackageTypeId         uint                   `json:"package_type_id"`
	GrossWeight           string                 `gorm:"size:64" json:"gross_weight"`
	Volume                string                 `gorm:"size:64" comment:"体积" json:"volume"`
	Remarks               string                 `gorm:"size:512" json:"remarks"`
	SeaCapLists           []SeaCapList           `gorm:"polymorphic:Source;" json:"sea_cap_lists"`
	TrailerCabinetNumbers []TrailerCabinetNumber `gorm:"foreignkey:FormerTrailerOrderId" json:"trailer_cabinet_numbers"`
	Departure             string                 `gorm:"size:256" json:"departure"`
	Destination           string                 `gorm:"size:256" json:"destination"`
}

type TrailerCabinetNumber struct {
	ID                   uint      `gorm:"primary_key"json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	FormerTrailerOrderId uint      `gorm:"index:idx_former_trailer_order_id" json:"former_trailer_order_id"`
	CabinetNumber        string    `gorm:"size:32" json:"cabinet_number"`
	SealNumber           string    `gorm:"size:32" json:"seal_number"`
}

type SeaCapList struct {
	ID            uint      `gorm:"primary_key"json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	OrderMasterId uint      `sql:"index" json:"order_master_id"`
	SourceId      uint      `gorm:"index:source_id_and_source_type" json:"source_id"`
	SourceType    string    `gorm:"size:32; index:source_id_and_source_type" json:"source_type"`
	Number        int       `json:"number"`
	CapType       string    `gorm:"size:32" json:"cap_type"`
}

//海运装货信息
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
	GrossUnit         uint      `comment:"重量单位" json:"gross_unit"`
	GrossWeight       string    `gorm:"size:32" comment:"毛重" json:"gross_weight"`
	ContainerWeight   string    `gorm:"size:32" comment:"柜重" json:"container_weight"`
	VgmWeight         string    `gorm:"size:32" comment:"VGM重量" json:"vgm_weight"`
	Volume            string    `gorm:"size:32" json:"volume"`
	Number            uint      `json:"number"`
	PackageTypeId     uint      `json:"package_type_id"`
	CapType           string    `gorm:"size:16" json:"cap_type"`
	DescriptionOfGood string    `gorm:"size:256" json:"description_of_good"`
	Marks             string    `gorm:"size:256" json:"marks"`
}

//仓库/场装单
type FormerWarehouseService struct {
	ID                    uint       `gorm:"primary_key"json:"id"`
	CreatedAt             time.Time  `json:"created_at"`
	OrderMasterId         uint       `gorm:"index:idx_order_master_id" json:"order_master_id"`
	InstructionId         uint       `gorm:"index:idx_instruction_id" json:"instruction_id"`
	WarehouseNo           string     `gorm:"size:32" json:"warehouse_no"`
	WarehouseDate         *time.Time `json:"warehouse_date"`
	WarehouseName         string     `gorm:"size:64" json:"warehouse_name"`
	WarehouseAddress      string     `gorm:"size:256" json:"warehouse_address"`
	WarehouseContact      string     `gorm:"size:32" json:"warehouse_contact"`
	WarehouseContactPhone string     `gorm:"size:16" json:"warehouse_contact_phone"`
	Marks                 string     `gorm:"size:2048" json:"marks"`
	DescriptionOfGood     string     `gorm:"size:2048" json:"description_of_good"`
	Size                  string     `gorm:"size:522" json:"size"`
	Number                int        `json:"number"`
	PackageTypeId         uint       `json:"package_type_id"`
	GrossWeight           string     `gorm:"size:64" json:"gross_weight"`
	Volume                string     `gorm:"size:64" comment:"体积" json:"volume"`
	DeliveryRemarks       string     `gorm:"size:256" json:"delivery_remarks"`
	DistributionRemarks   string     `gorm:"size:256" json:"distribution_remarks"`
}

const (
	InstructionMaster = "master"
	InstructionSplit  = "split"
)

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
func (FormerOtherService) TableName() string {
	return "former_other_services"
}

func (FormerTrailerOrder) TableName() string {
	return "former_trailer_orders"
}
func (FormerWarehouseService) TableName() string  {
	return "former_warehouse_services"
}

const (
	FormerDockTrailer = iota
	FormerInlandTrailer
	FormerHKTrailer
)
