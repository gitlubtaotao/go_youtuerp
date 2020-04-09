package models

import (
	"time"
)

type FormerAirsInstructions struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	ShipperContent            string    `xorm:"comment('发货人信息') TEXT"`
	ConsigneeContent          string    `xorm:"comment('收货人信息') TEXT"`
	NotifyPartyContent        string    `xorm:"TEXT"`
	DepartureAirportId        int       `xorm:"index INT(11)"`
	CodeOfTwoId               int       `xorm:"index INT(11)"`
	OceanChangesPaytypeId     string    `xorm:"VARCHAR(255)"`
	OtherChangesPaytypeId     string    `xorm:"VARCHAR(255)"`
	Marks                     string    `xorm:"comment('唛头') TEXT"`
	Number                    string    `xorm:"TEXT"`
	GrossWeight               string    `xorm:"comment('毛重') TEXT"`
	Size                      string    `xorm:"TEXT"`
	ShipperDeclaration        string    `xorm:"comment('发货人声明') TEXT"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	OrderMasterId             int64     `xorm:"index BIGINT(20)"`
	SerialNo                  string    `xorm:"VARCHAR(255)"`
	CompanyInstructionType    string    `xorm:"index(index_former_airs_instructions_on_company_instruction) VARCHAR(255)"`
	CompanyInstructionId      int64     `xorm:"index(index_former_airs_instructions_on_company_instruction) BIGINT(20)"`
	ContactInstructionType    string    `xorm:"index(index_former_airs_instructions_on_contact_instruction) VARCHAR(255)"`
	ContactInstructionId      int64     `xorm:"index(index_former_airs_instructions_on_contact_instruction) BIGINT(20)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	ContactInstructionContent string    `xorm:"TEXT"`
	ShipperType               string    `xorm:"index(index_former_airs_instructions_on_shipper_type_and_shipper_id) VARCHAR(255)"`
	ShipperId                 int64     `xorm:"index(index_former_airs_instructions_on_shipper_type_and_shipper_id) BIGINT(20)"`
	ConsigneeType             string    `xorm:"index(index_former_airs_instructions_on_contact_consignee) VARCHAR(255)"`
	ConsigneeId               int64     `xorm:"index(index_former_airs_instructions_on_contact_consignee) BIGINT(20)"`
	NotifyPartyType           string    `xorm:"index(index_former_airs_instructions_on_notify_party) VARCHAR(255)"`
	NotifyPartyId             int64     `xorm:"index(index_former_airs_instructions_on_notify_party) BIGINT(20)"`
	PodAgentType              string    `xorm:"index(index_former_airs_instructions_on_pod_agent) VARCHAR(255)"`
	PodAgentId                int64     `xorm:"index(index_former_airs_instructions_on_pod_agent) BIGINT(20)"`
	PodAgentContent           string    `xorm:"TEXT"`
	DestinationAirportId      int       `xorm:"index INT(11)"`
	TransshipmentAirportId    int       `xorm:"index INT(11)"`
	Status                    string    `xorm:"VARCHAR(255)"`
	InvoiceNo                 string    `xorm:"VARCHAR(255)"`
	RandomFile                int       `xorm:"default 0 TINYINT(1)"`
	FlightDate                time.Time `xorm:"DATETIME"`
	Remarks                   string    `xorm:"TEXT"`
	WaysOfDeclarationId       int       `xorm:"index INT(11)"`
	InstructionTypeId         int       `xorm:"index INT(11)"`
	ServiceContractNo         string    `xorm:"VARCHAR(255)"`
	BaseDataItemId            int       `xorm:"index INT(11)"`
	BaseDataTradeTermId       int       `xorm:"index INT(11)"`
	TransshipmentId           int       `xorm:"index INT(11)"`
	DescriptionOfGood         string    `xorm:"TEXT"`
	BaseDataPackageType       int       `xorm:"INT(11)"`
	Flight                    string    `xorm:"VARCHAR(255)"`
	DeliveryAddress           string    `xorm:"TEXT"`
	ArriveAddress             string    `xorm:"TEXT"`
	EstimatedTimeOfDeparture  time.Time `xorm:"DATETIME"`
	BaseDataBillProduceId     int       `xorm:"INT(11)"`
	Dimension                 string    `xorm:"comment('尺寸') TEXT"`
	PayPodId                  int       `xorm:"INT(11)"`
	HawbRemarks               string    `xorm:"TEXT"`
	IsGoodsAttachment         int       `xorm:"default 0 TINYINT(1)"`
	SubBlNo                   string    `xorm:"VARCHAR(255)"`
	CurrencyId                int       `xorm:"INT(11)"`
	ExecuteDate               time.Time `xorm:"DATETIME"`
	HblRemarks                string    `xorm:"TEXT"`
	Ratio                     float32   `xorm:"comment('材积换算系数') FLOAT"`
	RatioWeight               float32   `xorm:"comment('材积重') FLOAT"`
	Bubble                    float32   `xorm:"comment('分泡%') FLOAT"`
	ChargableWeight           float32   `xorm:"comment('计费重') FLOAT"`
	Type                      string    `xorm:"default 'Former::Airs::Instruction' VARCHAR(255)"`
}
