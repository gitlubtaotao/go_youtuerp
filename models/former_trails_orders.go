package models

import (
	"time"
)

type FormerTrailsOrders struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	CustomerNo                string    `xorm:"comment('工厂客户号') VARCHAR(255)"`
	UserCompanyId             int       `xorm:"index INT(11)"`
	Remarks                   string    `xorm:"comment('备注') TEXT"`
	InvoiceNo                 string    `xorm:"comment('发票号') VARCHAR(255)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	OrderMasterId             int       `xorm:"index INT(11)"`
	Status                    string    `xorm:"VARCHAR(255)"`
	SerialNo                  string    `xorm:"VARCHAR(255)"`
	InvoiceStatus             string    `xorm:"VARCHAR(255)"`
	AssignStatus              string    `xorm:"VARCHAR(255)"`
	DeletedAt                 time.Time `xorm:"index DATETIME"`
	AssociatedFormers         string    `xorm:"TEXT"`
	SoNo                      string    `xorm:"index VARCHAR(255)"`
	CompanyInstructionType    string    `xorm:"index(index_former_trails_orders_on_company_instruction) VARCHAR(255)"`
	CompanyInstructionId      int64     `xorm:"index(index_former_trails_orders_on_company_instruction) BIGINT(20)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	DriverMobi                string    `xorm:"comment('司机联系电话') VARCHAR(255)"`
	LoadingDate               string    `xorm:"VARCHAR(255)"`
	LicensePlateNumber        string    `xorm:"VARCHAR(255)"`
	Departure                 string    `xorm:"VARCHAR(100)"`
	Destination               string    `xorm:"VARCHAR(100)"`
	BaseDataInstructionTypeId int       `xorm:"comment('委托类型') INT(11)"`
	WaysOfDeclarationId       int       `xorm:"comment('报关方式') INT(11)"`
	BaseDataTradeTermsId      int       `xorm:"comment('贸易条款') INT(11)"`
	TransshipmentId           int       `xorm:"comment('转运') INT(11)"`
	BaseDataItemId            int       `xorm:"comment('装运条款') INT(11)"`
	ServiceContractNo         string    `xorm:"comment('合同编号') VARCHAR(255)"`
	OfWay                     int       `xorm:"default 0 comment('运输方式') SMALLINT(6)"`
	Number                    string    `xorm:"comment('拖车货物总件数') TEXT"`
	GrossWeight               string    `xorm:"comment('拖车总重量') TEXT"`
	Size                      string    `xorm:"comment('货物体积') TEXT"`
	DescriptionOfGood         string    `xorm:"comment('货物描述') TEXT"`
	Marks                     string    `xorm:"TEXT"`
	IsWeighing                int       `xorm:"comment('过磅') TINYINT(1)"`
	IsLockers                 int       `xorm:"comment('小柜摆尾') TINYINT(1)"`
	IsDeclare                 int       `xorm:"comment('报关单随车') TINYINT(1)"`
	IsDrivingLicense          int       `xorm:"TINYINT(1)"`
	ContainerNo               string    `xorm:"TEXT"`
	SealNo                    string    `xorm:"TEXT"`
	BoxSizeCount              string    `xorm:"VARCHAR(255)"`
	Dimension                 string    `xorm:"comment('尺寸') TEXT"`
	IsGoodsAttachment         int       `xorm:"default 0 comment('导出货描列表') TINYINT(1)"`
	BaseDataPackageType       int       `xorm:"INT(11)"`
}
