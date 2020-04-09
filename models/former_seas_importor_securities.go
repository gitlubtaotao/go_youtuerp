package models

import (
	"time"
)

type FormerSeasImportorSecurities struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId             int64     `xorm:"index BIGINT(20)"`
	SerialNo                  string    `xorm:"comment('表单序号') VARCHAR(255)"`
	SellerId                  int       `xorm:"comment('卖方信息，从cooperator中选择,这个是需要关联的') INT(11)"`
	BuyerId                   int       `xorm:"comment('买方信息，从cooperator中选择,这个需要关联') INT(11)"`
	PartyId                   int       `xorm:"comment('party信息，从cooperator中选择，与以上保持一致格式') INT(11)"`
	ManufacturerId            int       `xorm:"comment('制造商信息，从cooperator中选择') INT(11)"`
	OriginCountryId           int       `xorm:"comment('原产地国家') INT(11)"`
	ImportorIrsNo             string    `xorm:"comment('进口税号') VARCHAR(255)"`
	ConsigneeIrsNo            string    `xorm:"comment('收货人编号(联邦税务局数字)') VARCHAR(255)"`
	ResponsiblePersonName     string    `xorm:"comment('责任人名字') VARCHAR(255)"`
	ResponsiblePersonPhone    string    `xorm:"comment('责任人联系方式') VARCHAR(255)"`
	HbLNo                     string    `xorm:"comment('货代提单') VARCHAR(255)"`
	MbLNo                     string    `xorm:"comment('船东提单') VARCHAR(255)"`
	CtrnSeal                  string    `xorm:"comment('柜封号') VARCHAR(255)"`
	ScacCode                  string    `xorm:"comment('自申报号') VARCHAR(255)"`
	AmsBlNo                   string    `xorm:"comment('ams船东提单号') VARCHAR(255)"`
	ContainerStuffingLocation string    `xorm:"comment('货物装箱位置') TEXT"`
	Consolidator              int       `xorm:"comment('拼箱公司') INT(11)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	Status                    string    `xorm:"comment('状态') VARCHAR(255)"`
	InvoiceNo                 string    `xorm:"comment('发票号') VARCHAR(255)"`
	SellerContent             string    `xorm:"TEXT"`
	BuyerContent              string    `xorm:"TEXT"`
	PartyContent              string    `xorm:"TEXT"`
	ManufacturerContent       string    `xorm:"TEXT"`
	CompanyInstructionId      int       `xorm:"INT(11)"`
	CompanyInstructionType    string    `xorm:"VARCHAR(255)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	IsGoodsAttachment         int       `xorm:"default 0 comment('导出货描列表') TINYINT(1)"`
	ConsolidatorContent       string    `xorm:"comment('拼箱公司名称和地址') TEXT"`
	TariffSchedule            string    `xorm:"comment('HS商品编号（前六位）') VARCHAR(255)"`
	SascCode                  string    `xorm:"comment('SASC Code') VARCHAR(255)"`
	DeletedAt                 time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
