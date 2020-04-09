package models

import (
	"time"
)

type FormerSeasBookings struct {
	Id                          int64     `xorm:"pk autoincr BIGINT(20)"`
	SerialNo                    string    `xorm:"comment('序列号') VARCHAR(255)"`
	SoNo                        string    `xorm:"comment('SO编号') TEXT"`
	BookingDate                 time.Time `xorm:"comment('订舱日期') DATETIME"`
	ConfirmDate                 time.Time `xorm:"comment('确认日期') DATETIME"`
	SupplierCompanyAgentId      int       `xorm:"comment('供应商选择') index INT(11)"`
	SupplierAgentId             int       `xorm:"comment('该供应商联系人') index INT(11)"`
	CarrierContact              string    `xorm:"comment('船公司联系人') VARCHAR(255)"`
	CarrierTel                  string    `xorm:"comment('船公司联系电话') VARCHAR(255)"`
	EstimatedTimeOfDeparture    time.Time `xorm:"comment('开船日期') DATETIME"`
	EstimatedTimeOfArrival      time.Time `xorm:"comment('到港日期') DATETIME"`
	CutOffDate                  time.Time `xorm:"comment('截关日期') DATETIME"`
	CutDocDate                  time.Time `xorm:"comment('截文件日期') DATETIME"`
	CutCargoDate                time.Time `xorm:"comment('截货日期') DATETIME"`
	CargoReceivedDate           time.Time `xorm:"comment('收货人提货时间') DATETIME"`
	Vessel                      string    `xorm:"comment('船名') VARCHAR(255)"`
	Voyage                      string    `xorm:"comment('航次') VARCHAR(255)"`
	BaseDataSeaPortId           int       `xorm:"index INT(11)"`
	ChargeDescription           string    `xorm:"comment('费用描述') TEXT"`
	BookingNote                 string    `xorm:"comment('订舱备注') TEXT"`
	CreatedAt                   time.Time `xorm:"not null DATETIME"`
	UpdatedAt                   time.Time `xorm:"not null DATETIME"`
	Status                      string    `xorm:"VARCHAR(255)"`
	OrderMasterId               int       `xorm:"INT(11)"`
	InvoiceStatus               string    `xorm:"VARCHAR(255)"`
	AssignStatus                string    `xorm:"VARCHAR(255)"`
	BoatCompanyId               int       `xorm:"index INT(11)"`
	BoatCompanyName             string    `xorm:"comment('船名') VARCHAR(255)"`
	SupplierPhone               string    `xorm:"VARCHAR(255)"`
	Remarks                     string    `xorm:"TEXT"`
	InvoiceNo                   string    `xorm:"VARCHAR(255)"`
	ShipperType                 string    `xorm:"index(index_former_seas_bookings_on_shipper_type_and_shipper_id) VARCHAR(255)"`
	ShipperId                   int64     `xorm:"index(index_former_seas_bookings_on_shipper_type_and_shipper_id) BIGINT(20)"`
	ShipperContent              string    `xorm:"comment('发货人') TEXT"`
	ConsigneeType               string    `xorm:"index(index_former_seas_bookings_on_consignee_type_and_consignee_id) VARCHAR(255)"`
	ConsigneeId                 int64     `xorm:"index(index_former_seas_bookings_on_consignee_type_and_consignee_id) BIGINT(20)"`
	ConsigneeContent            string    `xorm:"comment('收货人') TEXT"`
	NotifyPartyType             string    `xorm:"index(index_notify_party) VARCHAR(255)"`
	NotifyPartyId               int64     `xorm:"index(index_notify_party) BIGINT(20)"`
	NotifyPartyContent          string    `xorm:"comment('通知人') TEXT"`
	CargoEn                     string    `xorm:"comment('货物英文名') TEXT"`
	CargoCn                     string    `xorm:"comment('货物中文名') TEXT"`
	Marks                       string    `xorm:"comment('标记') TEXT"`
	Number                      string    `xorm:"comment('包装件数') TEXT"`
	DescriptionOfGood           string    `xorm:"comment('品名') TEXT"`
	GrossWeight                 string    `xorm:"comment('毛重') TEXT"`
	Size                        string    `xorm:"comment('体积') TEXT"`
	SeaPortPolId                int       `xorm:"index INT(11)"`
	SeaPortViaId                int       `xorm:"index INT(11)"`
	SeaPortPodId                int       `xorm:"index INT(11)"`
	PlaceOfDelivery             string    `xorm:"comment('目的地') TEXT"`
	PlaceOfReceipt              string    `xorm:"comment('接货地') TEXT"`
	PodAgentType                string    `xorm:"index(index_former_seas_bookings_and pod_pagent) VARCHAR(255)"`
	PodAgentId                  int64     `xorm:"index(index_former_seas_bookings_and pod_pagent) BIGINT(20)"`
	PodAgentContent             string    `xorm:"TEXT"`
	SupplierCompanyAgentContent string    `xorm:"TEXT"`
	BaseDataPackageType         int       `xorm:"INT(11)"`
	SupplierAgentContent        string    `xorm:"TEXT"`
	BoxSizeCount                string    `xorm:"VARCHAR(255)"`
	OceanChangesPaytypeId       int       `xorm:"comment('运费支付方式') INT(11)"`
	OtherChangesPaytypeId       int       `xorm:"comment('其他运费支付方式') INT(11)"`
	DeletedAt                   time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	Dimension                   string    `xorm:"comment('尺寸') TEXT"`
}
