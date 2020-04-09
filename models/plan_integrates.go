package models

import (
	"time"
)

type PlanIntegrates struct {
	Id                     int64     `xorm:"pk autoincr BIGINT(20)"`
	PlanMainId             int64     `xorm:"index BIGINT(20)"`
	CreatedAt              time.Time `xorm:"not null DATETIME"`
	UpdatedAt              time.Time `xorm:"not null DATETIME"`
	Enabled                int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	Category               string    `xorm:"comment('类型[拖车，中港，内陆]') VARCHAR(64)"`
	TransportType          string    `xorm:"comment('运输方式') index VARCHAR(64)"`
	TrailerCompanyId       int       `xorm:"comment('拖车行') INT(11)"`
	TrailerContact         string    `xorm:"comment('拖车联系人') VARCHAR(64)"`
	TrailerPhone           string    `xorm:"comment('拖车联系人电话') VARCHAR(32)"`
	LoadingDate            time.Time `xorm:"comment('装货时间') DATETIME"`
	Departure              string    `xorm:"comment('装卸货地') VARCHAR(64)"`
	CarNumber              string    `xorm:"comment('车牌') VARCHAR(64)"`
	IsFumigation           int       `xorm:"default 0 comment('是否需要熏蒸') TINYINT(1)"`
	FumigationId           int       `xorm:"comment('熏蒸公司') INT(11)"`
	IsCertificate          int       `xorm:"default 0 comment('是否需要产地证') TINYINT(1)"`
	TraderId               int       `xorm:"comment('贸易商') INT(11)"`
	IsInsurance            int       `xorm:"default 0 comment('是否需要保险') TINYINT(1)"`
	InsuranceId            int       `xorm:"comment('保险公司') INT(11)"`
	Beneficiary            string    `xorm:"comment('受益人') VARCHAR(64)"`
	InsuranceAmount        string    `xorm:"comment('保险金额') DECIMAL(19,4)"`
	IsMagneticTest         int       `xorm:"default 0 comment('是否需要磁检') TINYINT(1)"`
	MagneticTestId         int       `xorm:"comment('磁检机构') INT(11)"`
	CompanyInstructionId   int       `xorm:"comment('委托单位') index INT(11)"`
	CompanyInstructionType string    `xorm:"comment('委托单位类型') VARCHAR(255)"`
	IsIdentification       int       `xorm:"default 0 comment('是否需要鉴定') TINYINT(1)"`
	IdentificationId       int       `xorm:"comment('鉴定机构') INT(11)"`
	IsDrivingLicense       int       `xorm:"default 0 comment('转关带司机本') TINYINT(1)"`
	IsDeclarationLicence   int       `xorm:"default 0 comment('报关证单随车') TINYINT(1)"`
	IsWeighing             int       `xorm:"default 0 comment('是否需要过磅') TINYINT(1)"`
	IsLockers              int       `xorm:"default 0 comment('是否需要小柜摆尾') TINYINT(1)"`
	Note                   string    `xorm:"comment('备注') TEXT"`
	IsWarehouse            int       `xorm:"default 0 comment('是否需要仓库/场装') TINYINT(1)"`
	WarehouseId            int       `xorm:"comment('仓库/场装') INT(11)"`
	Pol                    string    `xorm:"comment('起运港') VARCHAR(255)"`
	PlaceOfDeparture       string    `xorm:"comment('出发地') TEXT"`
	Destination            string    `xorm:"comment('目的地') TEXT"`
	CommodityInspectionId  int       `xorm:"INT(11)"`
	IsCommodityInspection  int       `xorm:"default 0 TINYINT(1)"`
	DeletedAt              time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion            int       `xorm:"default 0 INT(11)"`
}
