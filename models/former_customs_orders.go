package models

import (
	"time"
)

type FormerCustomsOrders struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	CustomNo                  string    `xorm:"comment('报关单号') VARCHAR(255)"`
	LicenceNo                 string    `xorm:"comment('许可证号') VARCHAR(255)"`
	InspectNo                 string    `xorm:"comment('商检编号') VARCHAR(255)"`
	HsCode                    string    `xorm:"comment('商品编码') VARCHAR(255)"`
	CustomBrokerType          string    `xorm:"index(custom_broker) VARCHAR(255)"`
	CustomBrokerId            int64     `xorm:"index(custom_broker) BIGINT(20)"`
	ApplicationDate           time.Time `xorm:"comment('报关日期') DATETIME"`
	BaseDataCurrencyId        int       `xorm:"comment('申报货币') INT(11)"`
	BaseDataCiqtypeId         int       `xorm:"comment('报关类型') INT(11)"`
	BaseDataSeaPortId         int       `xorm:"comment('港口') INT(11)"`
	ConfirmDate               time.Time `xorm:"comment('确认日期') DATETIME"`
	SubmitDate                time.Time `xorm:"DATETIME"`
	BaseDataDockId            int       `xorm:"comment('码头') INT(11)"`
	ContractNo                string    `xorm:"comment('合同编号') VARCHAR(255)"`
	PreRecordNo               string    `xorm:"comment('预录编号') VARCHAR(255)"`
	CancelNo                  string    `xorm:"comment('核销单号') VARCHAR(255)"`
	HasDrawback               int       `xorm:"comment('是否退税') TINYINT(1)"`
	DrawbackAddress           string    `xorm:"comment('退税地址') TEXT"`
	FileDeliverAddress        string    `xorm:"comment('报关资料送至地址') TEXT"`
	Remark                    string    `xorm:"comment('备注') TEXT"`
	OriginCountryId           int       `xorm:"comment('原产地（国家）') INT(11)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	SerialNo                  string    `xorm:"comment('序列号') VARCHAR(255)"`
	OrderMasterId             int       `xorm:"index INT(11)"`
	Status                    string    `xorm:"VARCHAR(255)"`
	InvoiceStatus             string    `xorm:"VARCHAR(255)"`
	InvoiceNo                 string    `xorm:"VARCHAR(255)"`
	DeletedAt                 time.Time `xorm:"index DATETIME"`
	CompanyInstructionId      int       `xorm:"INT(11)"`
	CompanyInstructionType    string    `xorm:"VARCHAR(255)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	CustomBrokerContent       string    `xorm:"TEXT"`
	DepartureCountryId        int       `xorm:"INT(11)"`
	DestinationAirportId      int       `xorm:"INT(11)"`
	WaysOfTransportationId    int       `xorm:"INT(11)"`
	DestinationCountryId      int       `xorm:"INT(11)"`
	IsGoodsAttachment         int       `xorm:"default 0 comment('导出货描列表') TINYINT(1)"`
}
