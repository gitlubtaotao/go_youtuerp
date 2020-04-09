package models

import (
	"time"
)

type FormerAirsBookings struct {
	Id                          int64     `xorm:"pk autoincr BIGINT(20)"`
	ShipperContent              string    `xorm:"comment('发货人信息') TEXT"`
	ConsigneeContent            string    `xorm:"comment('收货人信息') TEXT"`
	NotifyPartyContent          string    `xorm:"TEXT"`
	MawbNo                      string    `xorm:"comment('航空提单号码') VARCHAR(255)"`
	CarrierName                 string    `xorm:"comment('航空公司名称') VARCHAR(255)"`
	HawbNo                      string    `xorm:"comment('公司提单号') VARCHAR(255)"`
	DepartureAirportId          int       `xorm:"index INT(11)"`
	CodeOfTwoId                 int       `xorm:"index INT(11)"`
	Destination                 string    `xorm:"comment('目的地地址') TEXT"`
	OceanChangesPaytypeId       string    `xorm:"VARCHAR(255)"`
	OtherChangesPaytypeId       string    `xorm:"VARCHAR(255)"`
	Marks                       string    `xorm:"comment('唛头') TEXT"`
	PackageNo                   string    `xorm:"comment('包装数量') TEXT"`
	DescriptionOfGood           string    `xorm:"TEXT"`
	GrossWeight                 string    `xorm:"comment('毛重') TEXT"`
	Measurement                 string    `xorm:"comment('尺码') TEXT"`
	ShipperDeclaration          string    `xorm:"comment('发货人声明') TEXT"`
	CreatedAt                   time.Time `xorm:"not null DATETIME"`
	UpdatedAt                   time.Time `xorm:"not null DATETIME"`
	OrderMasterId               int64     `xorm:"index BIGINT(20)"`
	SerialNo                    string    `xorm:"comment('序列号') VARCHAR(255)"`
	BookingDate                 time.Time `xorm:"comment('订舱日期') DATETIME"`
	ConfirmDate                 time.Time `xorm:"comment('确认日期') DATETIME"`
	ShipperType                 string    `xorm:"index(index_former_airs_instructions_on_shipper) VARCHAR(255)"`
	ShipperId                   int64     `xorm:"index(index_former_airs_instructions_on_shipper) BIGINT(20)"`
	ConsigneeType               string    `xorm:"index(index_former_airs_instructions_on_consignee) VARCHAR(255)"`
	ConsigneeId                 int64     `xorm:"index(index_former_airs_instructions_on_consignee) BIGINT(20)"`
	NotifyPartyType             string    `xorm:"index(index_former_airs_instructions_on_notify_party) VARCHAR(255)"`
	NotifyPartyId               int64     `xorm:"index(index_former_airs_instructions_on_notify_party) BIGINT(20)"`
	PodAgentType                string    `xorm:"index(index_former_airs_instructions_on_pod_agent) VARCHAR(255)"`
	PodAgentId                  int64     `xorm:"index(index_former_airs_instructions_on_pod_agent) BIGINT(20)"`
	PodAgentContent             string    `xorm:"TEXT"`
	DestinationAirportId        int       `xorm:"index INT(11)"`
	TransshipmentAirportId      int       `xorm:"index INT(11)"`
	Status                      string    `xorm:"VARCHAR(255)"`
	InvoiceStatus               string    `xorm:"VARCHAR(255)"`
	InvoiceNo                   string    `xorm:"VARCHAR(255)"`
	AssignStatus                string    `xorm:"VARCHAR(255)"`
	SupplierAgentType           string    `xorm:"index(index_former_airs_instructions_on_supplier_agent) VARCHAR(255)"`
	SupplierAgentId             int64     `xorm:"index(index_former_airs_instructions_on_supplier_agent) BIGINT(20)"`
	SupplierAgentContent        string    `xorm:"TEXT"`
	FlightDate                  string    `xorm:"VARCHAR(255)"`
	CostDescription             string    `xorm:"TEXT"`
	Remarks                     string    `xorm:"TEXT"`
	UserMarketId                int       `xorm:"index INT(11)"`
	RandomFile                  int       `xorm:"default 0 TINYINT(1)"`
	AssociatedFormers           string    `xorm:"TEXT"`
	SupplierCompanyAgentId      int       `xorm:"index INT(11)"`
	SupplierCompanyAgentContent string    `xorm:"TEXT"`
	BaseDataPackageType         int       `xorm:"INT(11)"`
	Flight                      string    `xorm:"VARCHAR(255)"`
	DeliveryAddress             string    `xorm:"TEXT"`
	ArriveAddress               string    `xorm:"TEXT"`
	EstimatedTimeOfDeparture    time.Time `xorm:"DATETIME"`
	Dimension                   string    `xorm:"comment('尺寸') TEXT"`
}
