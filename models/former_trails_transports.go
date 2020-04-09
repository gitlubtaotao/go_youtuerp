package models

import (
	"time"
)

type FormerTrailsTransports struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	Type                      string    `xorm:"index VARCHAR(255)"`
	SerialNo                  string    `xorm:"comment('流水号') VARCHAR(255)"`
	CompanyInstructionType    string    `xorm:"index(index_former_trails_transports_company_instruction) VARCHAR(255)"`
	CompanyInstructionId      int64     `xorm:"index(index_former_trails_transports_company_instruction) BIGINT(20)"`
	CompanyInstructionContent string    `xorm:"comment('委托单位详情') TEXT"`
	OfWay                     int       `xorm:"default 0 comment('运输方式') index INT(11)"`
	UserCompanyId             int       `xorm:"comment('运输公司') index INT(11)"`
	UserId                    int       `xorm:"comment('联系人') index INT(11)"`
	InvoiceNo                 string    `xorm:"comment('发票') VARCHAR(255)"`
	InvoiceStatus             string    `xorm:"VARCHAR(255)"`
	AssignStatus              string    `xorm:"VARCHAR(255)"`
	Status                    string    `xorm:"comment('状态') VARCHAR(255)"`
	OrderMasterId             int       `xorm:"index INT(11)"`
	LockVersion               int       `xorm:"index INT(11)"`
	PlaceOfDeparture          string    `xorm:"comment('起运地') TEXT"`
	Destination               string    `xorm:"comment('目的地') TEXT"`
	AssociatedFormers         string    `xorm:"comment('关联表单') TEXT"`
	TruckerRepTel             string    `xorm:"comment('联系电话') VARCHAR(255)"`
	LicensePlateNumber        string    `xorm:"comment('车牌号') VARCHAR(255)"`
	HkLicensePlateNumber      string    `xorm:"comment('香港车牌号') VARCHAR(255)"`
	DriverMobi                string    `xorm:"comment('司机联系电话') VARCHAR(255)"`
	HkDriverMobi              string    `xorm:"comment('香港司机联系电话') VARCHAR(255)"`
	CustomsOrderInfo          string    `xorm:"comment('报关信息') TEXT"`
	Remarks                   string    `xorm:"comment('备注') TEXT"`
	DescriptionOfGood         string    `xorm:"comment('品名') TEXT"`
	Number                    string    `xorm:"comment('件数') TEXT"`
	GrossWeight               string    `xorm:"comment('重量') TEXT"`
	Size                      string    `xorm:"comment('体积') TEXT"`
	LoadingDate               string    `xorm:"VARCHAR(255)"`
	DeletedAt                 time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	Receiver                  string    `xorm:"VARCHAR(50)"`
	UserOperatorId            int       `xorm:"comment('操作') INT(11)"`
	UserSalesmanId            int       `xorm:"comment('业务') INT(11)"`
	UserMarketId              int       `xorm:"comment('市场') INT(11)"`
	UserCustomerId            int       `xorm:"comment('客服') INT(11)"`
	UserFileId                int       `xorm:"comment('文件') INT(11)"`
	BaseDataInstructionTypeId int       `xorm:"comment('委托类型') INT(11)"`
	WaysOfDeclarationId       int       `xorm:"comment('报关方式') INT(11)"`
	TransshipmentId           int       `xorm:"comment('转运') INT(11)"`
	BaseDataItemId            int       `xorm:"comment('装运条款') INT(11)"`
	ServiceContractNo         string    `xorm:"comment('合同编号') VARCHAR(255)"`
	Marks                     string    `xorm:"TEXT"`
	IsWeighing                int       `xorm:"comment('过磅') TINYINT(1)"`
	IsDrivingLicense          int       `xorm:"comment('转关司机带本') TINYINT(1)"`
	IsDeclare                 int       `xorm:"default 0 TINYINT(1)"`
	Dimension                 string    `xorm:"comment('尺寸') TEXT"`
	IsGoodsAttachment         int       `xorm:"default 0 comment('导出货描列表') TINYINT(1)"`
	BaseDataPackageType       int       `xorm:"INT(11)"`
}
