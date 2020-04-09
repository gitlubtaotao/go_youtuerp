package models

import (
	"time"
)

type FormerSeaDeliveries struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId             int64     `xorm:"index BIGINT(20)"`
	PolId                     int       `xorm:"comment('起运港') INT(11)"`
	PodId                     int       `xorm:"comment('目的港') INT(11)"`
	EstimatedTimeOfDeparture  time.Time `xorm:"comment('离港日期') DATETIME"`
	EstimatedTimeOfArrival    time.Time `xorm:"comment('到港日期') DATETIME"`
	CourierCompanyId          int       `xorm:"comment('快递公司') INT(11)"`
	CourierNumber             string    `xorm:"comment('快递单号') VARCHAR(64)"`
	BoatCompanyId             int       `xorm:"comment('船公司') INT(11)"`
	Vessel                    string    `xorm:"comment('船名') VARCHAR(128)"`
	Voyage                    string    `xorm:"comment('航次') VARCHAR(64)"`
	WarehousingDate           time.Time `xorm:"comment('入仓时间') DATETIME"`
	DeliverDate               time.Time `xorm:"comment('派送日期') DATETIME"`
	CompanyInstructionId      int       `xorm:"comment('委托单位') INT(11)"`
	CompanyInstructionType    string    `xorm:"comment('委托单位类型') VARCHAR(64)"`
	CompanyInstructionContent string    `xorm:"comment('委托单位详情') TEXT"`
	ShipperId                 int       `xorm:"comment('发货人') INT(11)"`
	ShipperType               string    `xorm:"VARCHAR(64)"`
	ShipperContent            string    `xorm:"comment('发货人详情') TEXT"`
	ConsigneeId               int       `xorm:"comment('收货人') INT(11)"`
	ConsigneeType             string    `xorm:"VARCHAR(64)"`
	ConsigneeContent          string    `xorm:"comment('收货人详情') TEXT"`
	NotifyPartyId             int       `xorm:"comment('通知人') INT(11)"`
	NotifyPartyType           string    `xorm:"VARCHAR(64)"`
	NotifyPartyContent        string    `xorm:"comment('通知人详情') TEXT"`
	OceanChangesPaytypeId     int       `xorm:"comment('运费付款方式') INT(11)"`
	OtherChangesPaytypeId     int       `xorm:"comment('其他付款方式') INT(11)"`
	InstructionTypeId         int       `xorm:"comment('委托类型') INT(11)"`
	WaysOfDeclarationId       int       `xorm:"comment('报关方式') INT(11)"`
	BaseDataTradeTermsId      int       `xorm:"comment('贸易条款') INT(11)"`
	TransshipmentId           int       `xorm:"comment('转运') INT(11)"`
	BaseDataItemId            int       `xorm:"comment('装运条款') INT(11)"`
	HsCode                    string    `xorm:"comment('商品编码') VARCHAR(128)"`
	ServiceContractNo         string    `xorm:"comment('合同编码') VARCHAR(64)"`
	InvoiceNo                 string    `xorm:"comment('发票号') VARCHAR(64)"`
	Insurance                 int       `xorm:"comment('是否需要发票') TINYINT(1)"`
	InsuranceAmount           string    `xorm:"comment('保险金额') DECIMAL(19,4)"`
	ViaId                     int       `xorm:"comment('中转港') INT(11)"`
	ReceiverAddress           string    `xorm:"comment('收货地址') TEXT"`
	ConsignorAddress          string    `xorm:"comment('发货地址') TEXT"`
	ReceiverCityId            int       `xorm:"comment('收货城市') INT(11)"`
	ConsignorCityId           int       `xorm:"comment('发货城市') INT(11)"`
	BargeDate                 time.Time `xorm:"comment('驳船日期') DATETIME"`
	BargeId                   int       `xorm:"comment('驳船名') INT(11)"`
	Enabled                   int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	Status                    string    `xorm:"comment('状态') VARCHAR(32)"`
	Note                      string    `xorm:"comment('备注') TEXT"`
	ClientNote                string    `xorm:"comment('客户备注') TEXT"`
	SerialNumber              string    `xorm:"comment('流水号') VARCHAR(64)"`
	Marks                     string    `xorm:"comment('唛头') TEXT"`
	Number                    string    `xorm:"comment('数量') TEXT"`
	DescriptionOfGood         string    `xorm:"comment('货物描述') TEXT"`
	GrossWeight               string    `xorm:"comment('毛重') TEXT"`
	Size                      string    `xorm:"comment('体积') TEXT"`
	SerialNo                  string    `xorm:"comment('序列号') VARCHAR(64)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	Ratio                     float32   `xorm:"comment('材积换算系数') FLOAT"`
	RatioWeight               float32   `xorm:"comment('材积重') FLOAT"`
	Bubble                    float32   `xorm:"comment('分泡%') FLOAT"`
	Dimension                 string    `xorm:"comment('体积') TEXT"`
	ChargedWeight             float32   `xorm:"comment('计费重') FLOAT"`
	BaseDataPackageType       int       `xorm:"comment('包装类型') INT(11)"`
	DeletedAt                 time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
