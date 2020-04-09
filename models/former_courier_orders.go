package models

import (
	"time"
)

type FormerCourierOrders struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId             int64     `xorm:"index BIGINT(20)"`
	SerialNo                  string    `xorm:"comment('序列号') VARCHAR(64)"`
	BaseDataInstructionTypeId int       `xorm:"comment('委托公司类型') INT(11)"`
	WaysOfDeclarationId       int       `xorm:"comment('报关方式') INT(11)"`
	BaseDataTradeTermsId      int       `xorm:"comment('贸易条款') INT(11)"`
	ServiceContractNo         string    `xorm:"comment('合同编号') VARCHAR(64)"`
	InvoiceNo                 string    `xorm:"comment('发票') VARCHAR(64)"`
	CourierNumber             string    `xorm:"comment('快递单号') VARCHAR(64)"`
	ReceiverAddress           string    `xorm:"comment('收货地址') TEXT"`
	Address                   string    `xorm:"comment('提货地址') TEXT"`
	Note                      string    `xorm:"comment('备注') TEXT"`
	Insurance                 int       `xorm:"comment('是否需要保险') TINYINT(1)"`
	InsuranceAmount           string    `xorm:"comment('保险') DECIMAL(15,4)"`
	Status                    string    `xorm:"comment('转态') VARCHAR(64)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	DeletedAt                 time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	CourierCodeId             int       `xorm:"index INT(11)"`
	ConsigneeId               int       `xorm:"INT(11)"`
	ConsigneeType             string    `xorm:"VARCHAR(255)"`
	ConsigneeContent          string    `xorm:"TEXT"`
	ShipperId                 int       `xorm:"INT(11)"`
	ShipperType               string    `xorm:"VARCHAR(255)"`
	ShipperContent            string    `xorm:"TEXT"`
	CompanyInstructionId      int       `xorm:"INT(11)"`
	CompanyInstructionType    string    `xorm:"VARCHAR(255)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	Marks                     string    `xorm:"TEXT"`
	Number                    string    `xorm:"TEXT"`
	DescriptionOfGood         string    `xorm:"TEXT"`
	GrossWeight               string    `xorm:"TEXT"`
	Size                      string    `xorm:"TEXT"`
	EstimatedTimeOfDeparture  time.Time `xorm:"DATETIME"`
	EstimatedTimeOfArrival    time.Time `xorm:"DATETIME"`
	CodeOfTwoId               int       `xorm:"INT(11)"`
	OceanChangesPaytypeId     int       `xorm:"INT(11)"`
	OtherChangesPaytypeId     int       `xorm:"INT(11)"`
	DestinationAirportId      int       `xorm:"INT(11)"`
	TransshipmentAirportId    int       `xorm:"INT(11)"`
	DepartureAirportId        int       `xorm:"INT(11)"`
	NotifyPartyId             int       `xorm:"INT(11)"`
	NotifyPartyType           string    `xorm:"VARCHAR(255)"`
	NotifyPartyContent        string    `xorm:"TEXT"`
	TransshipmentCode         string    `xorm:"VARCHAR(255)"`
	TransshipmentDate         time.Time `xorm:"DATETIME"`
	BaseDataPackageType       int       `xorm:"INT(11)"`
	CodeList                  string    `xorm:"comment('快递可能存在多个小包，多个快递的集合') TEXT"`
	Ratio                     float32   `xorm:"comment('材积换算系数') FLOAT"`
	RatioWeight               float32   `xorm:"comment('材积重') FLOAT"`
	Bubble                    float32   `xorm:"comment('分泡%') FLOAT"`
	Dimension                 string    `xorm:"comment('体积') TEXT"`
	ChargedWeight             float32   `xorm:"comment('计费重') FLOAT"`
	IsGoodsAttachment         int       `xorm:"default 0 comment('导出货描列表') TINYINT(1)"`
}
